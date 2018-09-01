package mathops

import (
	"math/big"
	"testing"
)

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

func TestProbability_PermutationsWithRepsBigInt_01(t *testing.T) {

	n := big.NewInt(5)
	r := big.NewInt(3)

	expectedResultStr := "125"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_02(t *testing.T) {

	n := big.NewInt(10)
	r := big.NewInt(3)

	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_03(t *testing.T) {

	n := big.NewInt(20)
	r := big.NewInt(5)

	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_04(t *testing.T) {

	n := big.NewInt(52)
	r := big.NewInt(5)

	expectedResultStr := "380204032"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_05(t *testing.T) {

	n := big.NewInt(23)
	r := big.NewInt(2)

	expectedResultStr := "529"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_06(t *testing.T) {

	n := big.NewInt(63)
	r := big.NewInt(4)

	expectedResultStr := "15752961"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_07(t *testing.T) {

	nInt := 4
	rInt := 63
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_08(t *testing.T) {

	nInt := 63
	rInt := -1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_09(t *testing.T) {

	nInt := -63
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_10(t *testing.T) {

	nInt := 0
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_11(t *testing.T) {

	nInt := 15
	rInt := 0
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_01(t *testing.T) {

	numOfItems := int64(3)
	numOfItemsPicked := int64(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_02(t *testing.T) {

	numOfItems := int64(3)
	numOfItemsPicked := int64(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}


}

func TestProbability_PermutationsInt64_03(t *testing.T) {

	numOfItems := int64(10)
	numOfItemsPicked := int64(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_04(t *testing.T) {

	numOfItems := int64(20)
	numOfItemsPicked := int64(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_05(t *testing.T) {

	numOfItems := int64(52)
	numOfItemsPicked := int64(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_06(t *testing.T) {

	numOfItems := int64(5)
	numOfItemsPicked := int64(3)
	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_07(t *testing.T) {

	numOfItems := int64(20)
	numOfItemsPicked := int64(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_08(t *testing.T) {
	nInt := 20
	rInt := 58
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_09(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_11(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_12(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_01(t *testing.T) {

	numOfItems := uint(3)
	numOfItemsPicked := uint(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_02(t *testing.T) {

	numOfItems := uint(3)
	numOfItemsPicked := uint(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}


}

func TestProbability_PermutationsUint_03(t *testing.T) {

	numOfItems := uint(10)
	numOfItemsPicked := uint(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_04(t *testing.T) {

	numOfItems := uint(20)
	numOfItemsPicked := uint(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_05(t *testing.T) {

	numOfItems := uint(52)
	numOfItemsPicked := uint(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_06(t *testing.T) {

	numOfItems := uint(5)
	numOfItemsPicked := uint(3)
	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_07(t *testing.T) {

	numOfItems := uint(20)
	numOfItemsPicked := uint(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_08(t *testing.T) {
	nInt := 20
	rInt := 58
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_09(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint32_01(t *testing.T) {

	numOfItems := uint32(3)
	numOfItemsPicked := uint32(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_02(t *testing.T) {

	numOfItems := uint32(3)
	numOfItemsPicked := uint32(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}


}

func TestProbability_PermutationsUint32_03(t *testing.T) {

	numOfItems := uint32(10)
	numOfItemsPicked := uint32(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_04(t *testing.T) {

	numOfItems := uint32(20)
	numOfItemsPicked := uint32(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_05(t *testing.T) {

	numOfItems := uint32(52)
	numOfItemsPicked := uint32(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_06(t *testing.T) {

	numOfItems := uint32(5)
	numOfItemsPicked := uint32(3)
	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_07(t *testing.T) {

	numOfItems := uint32(20)
	numOfItemsPicked := uint32(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint32_08(t *testing.T) {
	nInt := 20
	rInt := 58
	numOfItems := uint32(nInt)
	numOfItemsPicked := uint32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint32_09(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := uint32(nInt)
	numOfItemsPicked := uint32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint32_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := uint32(nInt)
	numOfItemsPicked := uint32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

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

func TestProbability_PermutationsUint64_04(t *testing.T) {

	numOfItems := uint64(20)
	numOfItemsPicked := uint64(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

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

func TestProbability_PermutationsUint64_05(t *testing.T) {

	numOfItems := uint64(52)
	numOfItemsPicked := uint64(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

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

func TestProbability_PermutationsUint64_06(t *testing.T) {

	numOfItems := uint64(5)
	numOfItemsPicked := uint64(3)
	allowRepetitions := true
	expectedResultStr := "125"

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

func TestProbability_PermutationsUint64_07(t *testing.T) {

	numOfItems := uint64(20)
	numOfItemsPicked := uint64(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

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

func TestProbability_PermutationsUint64_08(t *testing.T) {
	nInt := 20
	rInt := 58
	numOfItems := uint64(nInt)
	numOfItemsPicked := uint64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint64_09(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := uint64(nInt)
	numOfItemsPicked := uint64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint64_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := uint64(nInt)
	numOfItemsPicked := uint64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}
