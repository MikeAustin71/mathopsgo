package mathops

import "testing"

func TestProbability_CombinationsNumStrDto_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNumStrDto_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsNumStrDto_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

