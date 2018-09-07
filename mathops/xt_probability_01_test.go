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
