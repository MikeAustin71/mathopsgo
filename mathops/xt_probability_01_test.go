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

func TestProbability_PermutationsBigIntNum_01(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(3,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_02(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(3,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsBigIntNum_03(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(10,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_04(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(20,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_05(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(52,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_06(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(5,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_07(t *testing.T) {

	numOfItems  := BigIntNum{}.NewIntExponent(20,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsBigIntNum_08(t *testing.T) {

	nBigIntNum := 20
	rBigIntNum := 58

	numOfItems  := BigIntNum{}.NewIntExponent(nBigIntNum,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rBigIntNum, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nBigIntNum, rBigIntNum)
	}

}

func TestProbability_PermutationsBigIntNum_09(t *testing.T) {
	nBigIntNum := 0
	rBigIntNum := 4
	numOfItems  := BigIntNum{}.NewIntExponent(nBigIntNum,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rBigIntNum, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nBigIntNum, rBigIntNum)
	}

}

func TestProbability_PermutationsBigIntNum_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems  := BigIntNum{}.NewIntExponent(nInt,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_11(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems  := BigIntNum{}.NewIntExponent(nInt,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsBigIntNum_12(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems  := BigIntNum{}.NewIntExponent(nInt,0)
	numOfItemsPicked := BigIntNum{}.NewIntExponent(rInt, 0)
	allowRepetitions := true

	_, err := Probability{}.PermutationsBigIntNum(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_01(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3,0)

	numOfItemsPicked := Decimal{}.NewInt(2, 0)

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_02(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3,0)

	numOfItemsPicked := Decimal{}.NewInt(2, 0)

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsDecimal_03(t *testing.T) {

	numOfItems := Decimal{}.NewInt(10,0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_04(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20,0)


	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_05(t *testing.T) {

	numOfItems := Decimal{}.NewInt(52,0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_06(t *testing.T) {

	numOfItems := Decimal{}.NewInt(5,0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_07(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20,0)

	numOfItemsPicked := Decimal{}.NewInt(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsDecimal_08(t *testing.T) {

	nDecimal := 20
	rDecimal := 58

	numOfItems := Decimal{}.NewInt(nDecimal,0)

	numOfItemsPicked := Decimal{}.NewInt(rDecimal, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nDecimal, rDecimal)
	}

}

func TestProbability_PermutationsDecimal_09(t *testing.T) {
	nDecimal := 0
	rDecimal := 4

	numOfItems := Decimal{}.NewInt(nDecimal,0)

	numOfItemsPicked := Decimal{}.NewInt(rDecimal, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nDecimal, rDecimal)
	}

}

func TestProbability_PermutationsDecimal_10(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := Decimal{}.NewInt(nInt,0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_11(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems := Decimal{}.NewInt(nInt,0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsDecimal_12(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := Decimal{}.NewInt(nInt,0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_01(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(2, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_02(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(2, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsIntAry_03(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(10,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(3, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_04(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(20,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(20,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(5, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_05(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(52,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(52,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(5, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_06(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(3, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_07(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(20,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(20,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(5, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsIntAry_08(t *testing.T) {

	nIntAry := 20
	rIntAry := 58

	numOfItems, err  := IntAry{}.NewInt(nIntAry,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(20,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rIntAry, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(58,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nIntAry, rIntAry)
	}

}

func TestProbability_PermutationsIntAry_09(t *testing.T) {
	nIntAry := 0
	rIntAry := 4

	numOfItems, err  := IntAry{}.NewInt(nIntAry,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(0,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rIntAry, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(4,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nIntAry, rIntAry)
	}

}

func TestProbability_PermutationsIntAry_10(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems, err  := IntAry{}.NewInt(nInt,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(0,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rInt, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(4,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_11(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems, err  := IntAry{}.NewInt(nInt,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(-15,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rInt, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsIntAry_12(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems, err  := IntAry{}.NewInt(nInt,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(15,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rInt, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(-2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_01(t *testing.T) {

	numOfItems := Decimal{}.NewInt(3,0)

	numOfItemsPicked, err := IntAry{}.NewInt(2, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false

	expectedResultStr := "6"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_02(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := NumStrDto{}.NewInt(2, 0)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsINumMgr_03(t *testing.T) {

	numOfItems := BigIntNum{}.NewIntExponent(10,0)


	numOfItemsPicked, err := IntAry{}.NewInt(3, 0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(3,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_04(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20,0)

	numOfItemsPicked, err := NumStrDto{}.NewInt(5, 0)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}
}

func TestProbability_PermutationsINumMgr_05(t *testing.T) {

	numOfItems, err  := IntAry{}.NewInt(52,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(52,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(5, 0)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewInt(5,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_06(t *testing.T) {

	numOfItems  := BigIntNum{}.NewInt(5,0)

	numOfItemsPicked := Decimal{}.NewInt(3, 0)

	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_07(t *testing.T) {

	numOfItems := Decimal{}.NewInt(20,0)

	numOfItemsPicked := BigIntNum{}.NewInt(5, 0)

	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsINumMgr_08(t *testing.T) {

	nINumMgr := 20
	rINumMgr := 58

	numOfItems, err  := IntAry{}.NewInt(nINumMgr,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(20,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := NumStrDto{}.NewInt(rINumMgr, 0)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(58,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nINumMgr, rINumMgr)
	}

}

func TestProbability_PermutationsINumMgr_09(t *testing.T) {
	nINumMgr := 0
	rINumMgr := 4

	numOfItems, err  := IntAry{}.NewInt(nINumMgr,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(0,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := NumStrDto{}.NewInt(rINumMgr, 0)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(4,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nINumMgr, rINumMgr)
	}

}

func TestProbability_PermutationsINumMgr_10(t *testing.T) {
	nInt := 15
	rInt := 0

	numOfItems := Decimal{}.NewInt(nInt,0)

	numOfItemsPicked := BigIntNum{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_11(t *testing.T) {
	nInt := -15
	rInt := 2

	numOfItems, err  := IntAry{}.NewInt(nInt,0)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewInt(-15,0). " +
			"Error='%v' ", err.Error())
	}

	numOfItemsPicked, err := IntAry{}.NewInt(rInt, 0)

	if err != nil {
		t.Errorf("Error returned by INumMgr{}.NewInt(2,0). " +
			"Error='%v' ", err.Error())
	}

	allowRepetitions := true

	_, err = Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsINumMgr_12(t *testing.T) {
	nInt := 15
	rInt := -2

	numOfItems := Decimal{}.NewInt(nInt,0)

	numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

	allowRepetitions := true

	_, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) " +
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
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
		t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). " +
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt_08(t *testing.T) {
	nInt := 20
	rInt := 58
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_09(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_10(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_11(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt_12(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int(nInt)
	numOfItemsPicked := int(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) " +
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}
