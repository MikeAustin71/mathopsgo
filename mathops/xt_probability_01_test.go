package mathops

import (
	"math/big"
	"testing"
)

func TestProbability_CombinationsNoRepsBigInt_01(t *testing.T) {
	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "560"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_02(t *testing.T) {
	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "1820"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_03(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "2598960"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_04(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "495918532948104"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_05(t *testing.T) {
	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "31824"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_06(t *testing.T) {
	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "26334"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_07(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "3819816"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsNoRepsBigInt_08(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "12"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "1"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "25"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}
}

func TestProbability_CombinationsNoRepsBigInt_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItemsChosen=0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems=0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItemsChosen < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_01(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "35"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_02(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "705432"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_03(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "351"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_04(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "63205303218876"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_05(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "817190"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_06(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "7726160"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_07(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "84"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_08(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "36"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_09(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "8936928"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_10(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "79208745"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_11(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "77558760"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_12(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "12"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt("+
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_14(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItemsChosen == 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_15(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_16(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItemsChosen < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsBigIntNum_01(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 3
	expectedResultStr := "560"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_02(t *testing.T) {

	numOfItemsInt := 16
	numOfItemsChosenInt := 12
	expectedResultStr := "1820"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_03(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 5
	expectedResultStr := "2598960"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_04(t *testing.T) {

	numOfItemsInt := 52
	numOfItemsChosenInt := 26
	expectedResultStr := "495918532948104"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_05(t *testing.T) {

	numOfItemsInt := 18
	numOfItemsChosenInt := 7
	expectedResultStr := "31824"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_06(t *testing.T) {

	numOfItemsInt := 22
	numOfItemsChosenInt := 5
	expectedResultStr := "26334"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_07(t *testing.T) {

	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_08(t *testing.T) {
	numOfItemsInt := 56
	numOfItemsChosenInt := 5
	expectedResultStr := "3819816"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_09(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	expectedResultStr := "1"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_10(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 1
	expectedResultStr := "25"
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_11(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_12(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_13(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_14(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_15(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	allowRepetitions := false

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' " +
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_16(t *testing.T) {
	numOfItemsInt := 5
	numOfItemsChosenInt := 3
	expectedResultStr := "35"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_17(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 11
	expectedResultStr := "705432"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_18(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 2
	expectedResultStr := "351"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_19(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 24
	expectedResultStr := "63205303218876"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_20(t *testing.T) {
	numOfItemsInt := 10
	numOfItemsChosenInt := 14
	expectedResultStr := "817190"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_21(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 15
	expectedResultStr := "7726160"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_22(t *testing.T) {
	numOfItemsInt := 7
	numOfItemsChosenInt := 3
	expectedResultStr := "84"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_23(t *testing.T) {
	numOfItemsInt := 3
	numOfItemsChosenInt := 7
	expectedResultStr := "36"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_24(t *testing.T) {
	numOfItemsInt := 62
	numOfItemsChosenInt := 5
	expectedResultStr := "8936928"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_25(t *testing.T) {
	numOfItemsInt := 97
	numOfItemsChosenInt := 5
	expectedResultStr := "79208745"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_26(t *testing.T) {
	numOfItemsInt := 15
	numOfItemsChosenInt := 15
	expectedResultStr := "77558760"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_27(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	expectedResultStr := "12"
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	result, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsBigIntNum("+
			"numOfItems, numOfItemsChosen, allowRepetitions). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}
}

func TestProbability_CombinationsBigIntNum_28(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_29(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_30(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}

func TestProbability_CombinationsBigIntNum_31(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	allowRepetitions := true

	numOfItems := BigIntNum{}.NewInt(numOfItemsInt, 0)
	numOfItemsChosen := BigIntNum{}.NewInt(numOfItemsChosenInt, 0)

	_, err := Probability{}.CombinationsBigIntNum(numOfItems, numOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}
}
