package mathops

import "testing"

func TestProbability_CombinationsINumMgr_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := IntAry{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := IntAry{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := IntAry{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := IntAry{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsINumMgr("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsINumMgr_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := Decimal{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := IntAry{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsINumMgr_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := Decimal{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsINumMgr(&numOfItems, &numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := int(numOfItemsInt)
	numOfItemsChosen := int(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt32_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt32_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := int32(numOfItemsInt)
	numOfItemsChosen := int32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsInt64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsInt64_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsInt64_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := int64(numOfItemsInt)
	numOfItemsChosen := int64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsInt64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

