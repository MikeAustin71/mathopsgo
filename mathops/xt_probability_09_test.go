package mathops

import "testing"

func TestProbability_CombinationsUint_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint_14(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_15(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_16(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_17(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_18(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_19(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_20(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_21(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_22(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_23(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_24(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_25(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint_26(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := uint(numOfItemsInt)
	numOfItemsChosen := uint(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint32_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint32_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint32_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint32_14(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_15(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_16(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_17(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_18(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_19(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_20(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_21(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_22(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_23(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_24(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_25(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint32("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint32_26(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint32_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := uint32(numOfItemsInt)
	numOfItemsChosen := uint32(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint64_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint64_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint64_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint64_14(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_15(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_16(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_17(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_18(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_19(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_20(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_21(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_22(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_23(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_24(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_25(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsUint64("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsUint64_26(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsUint64_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := uint64(numOfItemsInt)
	numOfItemsChosen := uint64(numOfItemsChosenInt)

	_, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}
