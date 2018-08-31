package mathops

import (
	"math/big"
	"testing"
)



func TestProbability_PermutationsUint64_01(t *testing.T) {

	numOfItems := uint64(3)
	numOfItemsPicked := uint64(2)
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked)

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

func TestProbability_PermutationsBigInt_01(t *testing.T) {

	n := big.NewInt(5)
	r := big.NewInt(3)

	expectedResultStr := "60"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsBigInt_02(t *testing.T) {

	n := big.NewInt(12)
	r := big.NewInt(2)

	expectedResultStr := "132"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}



func TestProbability_PermutationsBigInt_03(t *testing.T) {

	n := big.NewInt(20)
	r := big.NewInt(5)

	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsBigInt_04(t *testing.T) {

	n := big.NewInt(52)
	r := big.NewInt(5)

	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsBigInt_05(t *testing.T) {

	n := big.NewInt(23)
	r := big.NewInt(2)

	expectedResultStr := "506"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsBigInt_06(t *testing.T) {

	n := big.NewInt(63)
	r := big.NewInt(4)

	expectedResultStr := "14295960"

	result, err := Probability{}.PermutationsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsBigInt_07(t *testing.T) {

	nInt := 4
	rInt := 63
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigInt_08(t *testing.T) {

	nInt := 63
	rInt := -1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigInt(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigInt_09(t *testing.T) {

	nInt := -63
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigInt(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigInt_10(t *testing.T) {

	nInt := 0
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigInt(n, r) " +
			"However no error was generated. n == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigInt_11(t *testing.T) {

	nInt := 15
	rInt := 0
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))


	_, err := Probability{}.PermutationsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigInt(n, r) " +
			"However no error was generated. r == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}