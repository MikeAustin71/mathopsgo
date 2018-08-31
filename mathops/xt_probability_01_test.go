package mathops

import (
	"math/big"
	"testing"
)




func TestProbability_PermutationsUint64_01(t *testing.T) {

	numOfItems := uint64(3)
	numOfItemsPicked := uint64(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}


}

func TestProbability_PermutationsUint64_02(t *testing.T) {

	numOfItems := uint64(3)
	numOfItemsPicked := uint64(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}


}

func TestProbability_PermutationsUint64_03(t *testing.T) {

	numOfItems := uint64(10)
	numOfItemsPicked := uint64(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}


func TestProbability_PermutationsNoRepsBigInt_01(t *testing.T) {

	n := big.NewInt(5)
	r := big.NewInt(3)

	expectedResultStr := "60"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_02(t *testing.T) {

	n := big.NewInt(12)
	r := big.NewInt(2)

	expectedResultStr := "132"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_03(t *testing.T) {

	n := big.NewInt(20)
	r := big.NewInt(5)

	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_04(t *testing.T) {

	n := big.NewInt(52)
	r := big.NewInt(5)

	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_05(t *testing.T) {

	n := big.NewInt(23)
	r := big.NewInt(2)

	expectedResultStr := "506"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_06(t *testing.T) {

	n := big.NewInt(63)
	r := big.NewInt(4)

	expectedResultStr := "14295960"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_07(t *testing.T) {

	nInt := 4
	rInt := 63
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_08(t *testing.T) {

	nInt := 63
	rInt := -1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_09(t *testing.T) {

	nInt := -63
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_10(t *testing.T) {

	nInt := 0
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"However no error was generated. n == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_11(t *testing.T) {

	nInt := 15
	rInt := 0
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) " +
			"However no error was generated. r == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}