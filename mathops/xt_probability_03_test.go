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
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_06(t *testing.T) {

	n := big.NewInt(6)
	r := big.NewInt(6)

	expectedResultStr := "720"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_07(t *testing.T) {

	n := big.NewInt(63)
	r := big.NewInt(4)

	expectedResultStr := "14295960"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_08(t *testing.T) {

	n := big.NewInt(63)
	r := big.NewInt(1)

	expectedResultStr := "63"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_09(t *testing.T) {

	n := big.NewInt(26)
	r := big.NewInt(3)

	expectedResultStr := "15600"

	result, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsNoRepsBigInt_10(t *testing.T) {

	nInt := 4
	rInt := 63
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_11(t *testing.T) {

	nInt := 63
	rInt := -1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_12(t *testing.T) {

	nInt := -63
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_13(t *testing.T) {

	nInt := 0
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"However no error was generated. n == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsNoRepsBigInt_14(t *testing.T) {

	nInt := 15
	rInt := 0
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsNoRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsNoRepsBigInt(n, r) "+
			"However no error was generated. r == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_01(t *testing.T) {

	n := big.NewInt(5)
	r := big.NewInt(3)

	expectedResultStr := "125"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
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
	rInt := 10
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	expectedResultStr := "1048576"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_08(t *testing.T) {

	nInt := 9
	rInt := 9
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_09(t *testing.T) {

	nInt := 10
	rInt := 1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	expectedResultStr := "10"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_10(t *testing.T) {

	nInt := 5
	rInt := 5
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	expectedResultStr := "3125"

	result, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResultStr, actualResultStr)
	}

}

func TestProbability_PermutationsWithRepsBigInt_11(t *testing.T) {

	nInt := 63
	rInt := -1
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_12(t *testing.T) {

	nInt := -63
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_13(t *testing.T) {

	nInt := 0
	rInt := 4
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsWithRepsBigInt_14(t *testing.T) {

	nInt := 15
	rInt := 0
	n := big.NewInt(int64(nInt))
	r := big.NewInt(int64(rInt))

	_, err := Probability{}.PermutationsWithRepsBigInt(n, r)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r == 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_01(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(3, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_02(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(3, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsBigIntNum_03(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(10, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_04(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(20, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_05(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(52, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_06(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(5, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_07(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(20, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_08(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(5, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(11, 0)

	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_09(t *testing.T) {
	nBigIntNum := 0
	rBigIntNum := 4
	numOfItems := BigIntNum{}.NewIntExponent(nBigIntNum, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rBigIntNum, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nBigIntNum, rBigIntNum)
	}

}

func TestProbability_PermutationsBigIntNum_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_11(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_12(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_13(t *testing.T) {
	nInt := 5
	rInt := 11
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := false

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}


func TestProbability_PermutationsBigIntNum_14(t *testing.T) {
	nBigIntNum := 0
	rBigIntNum := 4
	numOfItems := BigIntNum{}.NewIntExponent(nBigIntNum, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rBigIntNum, 0)
	allowRepetitions := false

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nBigIntNum, rBigIntNum)
	}

}

func TestProbability_PermutationsBigIntNum_15(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := false

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_16(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := false

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_17(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := BigIntNum{}.NewIntExponent(nInt, 0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := false

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_01(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3, 0)

	numOfItemsPicked := Decimal{}.NewInt(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_02(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3, 0)

	numOfItemsPicked := Decimal{}.NewInt(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsDecimal_03(t *testing.T) {

	numOfItems := Decimal{}.NewInt(10, 0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_04(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20, 0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_05(t *testing.T) {

	numOfItems := Decimal{}.NewInt(52, 0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_06(t *testing.T) {

	numOfItems := Decimal{}.NewInt(5, 0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_07(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20, 0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_08(t *testing.T) {

	numOfItems := Decimal{}.NewInt(5, 0)

	numOfItemsPicked := Decimal{}.NewInt(11, 0)

	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_09(t *testing.T) {

	numOfItems := Decimal{}.NewInt(56, 0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_10(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_11(t *testing.T) {

	numOfItems := Decimal{}.NewInt(12, 0)

	numOfItemsPicked := Decimal{}.NewInt(7, 0)

	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_12(t *testing.T) {

	numOfItems := Decimal{}.NewInt(18, 0)

	numOfItemsPicked := Decimal{}.NewInt(8, 0)

	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_13(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := Decimal{}.NewInt(9, 0)

	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_14(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := Decimal{}.NewInt(9, 0)

	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}


func TestProbability_PermutationsDecimal_15(t *testing.T) {
	nDecimal := 0
	rDecimal := 4

	numOfItems := Decimal{}.NewInt(nDecimal, 0)

	numOfItemsPicked := Decimal{}.NewInt(rDecimal, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nDecimal, rDecimal)
	}

}

func TestProbability_PermutationsDecimal_16(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_17(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_18(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_19(t *testing.T) {
	nInt := 5
	rInt := 11

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(" +
			"numOfItems, numOfItemsPicked, allowRepetitions) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_20(t *testing.T) {
	nDecimal := 0
	rDecimal := 4

	numOfItems := Decimal{}.NewInt(nDecimal, 0)

	numOfItemsPicked := Decimal{}.NewInt(rDecimal, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nDecimal, rDecimal)
	}

}

func TestProbability_PermutationsDecimal_21(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_22(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_23(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_01(t *testing.T) {

	numOfItems := IntAry{}.NewInt(3, 0)

	numOfItemsPicked := IntAry{}.NewInt(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_02(t *testing.T) {

	numOfItems := IntAry{}.NewInt(3, 0)

	numOfItemsPicked := IntAry{}.NewInt(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsIntAry_03(t *testing.T) {

	numOfItems := IntAry{}.NewInt(10, 0)

	numOfItemsPicked := IntAry{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_04(t *testing.T) {

	numOfItems := IntAry{}.NewInt(20, 0)

	numOfItemsPicked := IntAry{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_05(t *testing.T) {

	numOfItems := IntAry{}.NewInt(52, 0)

	numOfItemsPicked := IntAry{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_06(t *testing.T) {

	numOfItems := IntAry{}.NewInt(5, 0)

	numOfItemsPicked := IntAry{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_07(t *testing.T) {

	numOfItems := IntAry{}.NewInt(20, 0)

	numOfItemsPicked := IntAry{}.NewInt(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_08(t *testing.T) {

	numOfItems := IntAry{}.NewInt(5, 0)

	numOfItemsPicked := IntAry{}.NewInt(11, 0)

	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_09(t *testing.T) {

	numOfItems := IntAry{}.NewInt(56, 0)

	numOfItemsPicked := IntAry{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_10(t *testing.T) {

	numOfItems := IntAry{}.NewInt(9, 0)

	numOfItemsPicked := IntAry{}.NewInt(3, 0)

	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_11(t *testing.T) {

	numOfItems := IntAry{}.NewInt(12, 0)

	numOfItemsPicked := IntAry{}.NewInt(7, 0)

	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_12(t *testing.T) {

	numOfItems := IntAry{}.NewInt(18, 0)

	numOfItemsPicked := IntAry{}.NewInt(8, 0)

	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_13(t *testing.T) {

	numOfItems := IntAry{}.NewInt(9, 0)

	numOfItemsPicked := IntAry{}.NewInt(9, 0)

	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_14(t *testing.T) {

	numOfItems := IntAry{}.NewInt(9, 0)

	numOfItemsPicked := IntAry{}.NewInt(9, 0)

	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_15(t *testing.T) {
	nIntAry := 0
	rIntAry := 4

	numOfItems := IntAry{}.NewInt(nIntAry, 0)

	numOfItemsPicked := IntAry{}.NewInt(rIntAry, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nIntAry, rIntAry)
	}

}

func TestProbability_PermutationsIntAry_16(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_17(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_18(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_19(t *testing.T) {
	nInt := 5
	rInt := 11

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(" +
			"numOfItems, numOfItemsPicked, allowRepetitions) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_20(t *testing.T) {
	nIntAry := 0
	rIntAry := 4

	numOfItems := IntAry{}.NewInt(nIntAry, 0)

	numOfItemsPicked := IntAry{}.NewInt(rIntAry, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nIntAry, rIntAry)
	}

}

func TestProbability_PermutationsIntAry_21(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_22(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_23(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := false

	_, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}
}

func TestProbability_PermutationsINumMgr_01(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3, 0)

	numOfItemsPicked := IntAry{}.NewInt(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_02(t *testing.T) {

	numOfItems := IntAry{}.NewInt(3, 0)

	numOfItemsPicked := NumStrDto{}.NewInt(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsINumMgr_03(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(10, 0)

	numOfItemsPicked := IntAry{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_04(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20, 0)

	numOfItemsPicked := NumStrDto{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsINumMgr_05(t *testing.T) {

	numOfItems := IntAry{}.NewInt(52, 0)

	numOfItemsPicked := IntAry{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsINumMgr_06(t *testing.T) {

	numOfItems := BigIntNum{}.NewInt(5, 0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_07(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_08(t *testing.T) {

	numOfItems := Decimal{}.NewInt(5, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(11, 0)

	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_09(t *testing.T) {

	numOfItems := Decimal{}.NewInt(56, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_10(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(3, 0)

	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_11(t *testing.T) {

	numOfItems := Decimal{}.NewInt(12, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(7, 0)

	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_12(t *testing.T) {

	numOfItems := Decimal{}.NewInt(18, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(8, 0)

	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_13(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(9, 0)

	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_14(t *testing.T) {

	numOfItems := Decimal{}.NewInt(9, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(9, 0)

	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_15(t *testing.T) {
	nINumMgr := 0
	rINumMgr := 4

	numOfItems := IntAry{}.NewInt(nINumMgr, 0)

	numOfItemsPicked := NumStrDto{}.NewInt(rINumMgr, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nINumMgr, rINumMgr)
	}

}

func TestProbability_PermutationsINumMgr_16(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_17(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_18(t *testing.T) {
	nInt := 15
	rInt := -2
	allowRepetitions := true

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)


	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_19(t *testing.T) {
	nInt := 5
	rInt := 11
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)


	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(" +
			"&numOfItems, &numOfItemsPicked, allowRepetitions) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_20(t *testing.T) {
	nINumMgr := 0
	rINumMgr := 4
	allowRepetitions := false

	numOfItems := IntAry{}.NewInt(nINumMgr, 0)

	numOfItemsPicked := NumStrDto{}.NewInt(rINumMgr, 0)

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nINumMgr, rINumMgr)
	}

}

func TestProbability_PermutationsINumMgr_21(t *testing.T) {
	nInt := 15
	rInt := 0
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := BigIntNum{}.NewInt(rInt, 0)

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_22(t *testing.T) {
	nInt := -15
	rInt := 2
	allowRepetitions := false

	numOfItems := IntAry{}.NewInt(nInt, 0)

	numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_23(t *testing.T) {
	nInt := 15
	rInt := -2
	allowRepetitions := false

	numOfItems := Decimal{}.NewInt(nInt, 0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_01(t *testing.T) {

	numOfItems := int(3)
	numOfItemsPicked := int(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_02(t *testing.T) {

	numOfItems := int(3)
	numOfItemsPicked := int(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_03(t *testing.T) {

	numOfItems := int(10)
	numOfItemsPicked := int(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_04(t *testing.T) {

	numOfItems := int(20)
	numOfItemsPicked := int(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_05(t *testing.T) {

	numOfItems := int(52)
	numOfItemsPicked := int(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_06(t *testing.T) {

	numOfItems := int(5)
	numOfItemsPicked := int(3)
	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_07(t *testing.T) {

	numOfItems := int(20)
	numOfItemsPicked := int(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_08(t *testing.T) {

	numOfItems := int(5)
	numOfItemsPicked := int(11)
	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_09(t *testing.T) {

	numOfItems := int(56)
	numOfItemsPicked := int(5)
	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_10(t *testing.T) {

	numOfItems := int(9)
	numOfItemsPicked := int(3)
	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_11(t *testing.T) {

	numOfItems := int(12)
	numOfItemsPicked := int(7)
	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_12(t *testing.T) {

	numOfItems := int(18)
	numOfItemsPicked := int(8)
	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_13(t *testing.T) {

	numOfItems := int(9)
	numOfItemsPicked := int(9)
	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_14(t *testing.T) {

	numOfItems := int(9)
	numOfItemsPicked := int(9)
	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsInt_15(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_16(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_17(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_18(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_19(t *testing.T) {
	nInt := 5
	rInt := 11
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt(" +
			"numOfItems, numOfItemsPicked, allowRepetitions) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_20(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_21(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_22(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_23(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}
