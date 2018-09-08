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
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_05(t *testing.T) {
	numOfItemsInt := 25
	numOfItemsChosenInt := 25
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedResultStr := "1"

	result, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsNoRepsBigInt(" +
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsNoRepsBigInt_06(t *testing.T) {
	numOfItemsInt := 26
	numOfItemsChosenInt := 52
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))


	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil " +
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_07(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := 0
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))


	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil " +
			"numOfItemsChosen=0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_08(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))


	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil " +
			"numOfItems=0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_09(t *testing.T) {
	numOfItemsInt := -52
	numOfItemsChosenInt := 26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))


	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil " +
			"numOfItems < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsNoRepsBigInt_10(t *testing.T) {
	numOfItemsInt := 52
	numOfItemsChosenInt := -26
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))


	_, err := Probability{}.CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil " +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
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
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}


func TestProbability_CombinationsWithRepsBigInt_07(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 1
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))
	expectedResultStr := "12"

	result, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("Error returned by Probability{}.CombinationsWithRepsBigInt(" +
			"numOfItems, numOfItemsChosen). Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_CombinationsWithRepsBigInt_08(t *testing.T) {
	numOfItemsInt := 0
	numOfItemsChosenInt := 15
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil " +
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_09(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := 0
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil " +
			"numOfItemsChosen == 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_10(t *testing.T) {
	numOfItemsInt := -12
	numOfItemsChosenInt := 6
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil " +
			"numOfItems < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}

func TestProbability_CombinationsWithRepsBigInt_11(t *testing.T) {
	numOfItemsInt := 12
	numOfItemsChosenInt := -6
	numOfItems := big.NewInt(int64(numOfItemsInt))
	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := Probability{}.CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil " +
			"numOfItemsChosen < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

}