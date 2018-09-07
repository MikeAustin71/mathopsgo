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
