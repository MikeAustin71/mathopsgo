package mathops

import "testing"

func TestProbability_PermutationsInt32_01(t *testing.T) {

	numOfItems := int32(3)
	numOfItemsPicked := int32(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_02(t *testing.T) {

	numOfItems := int32(3)
	numOfItemsPicked := int32(2)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_03(t *testing.T) {

	numOfItems := int32(10)
	numOfItemsPicked := int32(3)
	allowRepetitions := true
	expectedResultStr := "1000"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_04(t *testing.T) {

	numOfItems := int32(20)
	numOfItemsPicked := int32(5)
	allowRepetitions := false
	expectedResultStr := "1860480"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_05(t *testing.T) {

	numOfItems := int32(52)
	numOfItemsPicked := int32(5)
	allowRepetitions := false
	expectedResultStr := "311875200"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_06(t *testing.T) {

	numOfItems := int32(5)
	numOfItemsPicked := int32(3)
	allowRepetitions := true
	expectedResultStr := "125"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_07(t *testing.T) {

	numOfItems := int32(20)
	numOfItemsPicked := int32(5)
	allowRepetitions := true
	expectedResultStr := "3200000"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_08(t *testing.T) {

	numOfItems := int32(5)
	numOfItemsPicked := int32(11)
	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_09(t *testing.T) {

	numOfItems := int32(56)
	numOfItemsPicked := int32(5)
	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_10(t *testing.T) {

	numOfItems := int32(9)
	numOfItemsPicked := int32(3)
	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_11(t *testing.T) {

	numOfItems := int32(12)
	numOfItemsPicked := int32(7)
	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_12(t *testing.T) {

	numOfItems := int32(18)
	numOfItemsPicked := int32(8)
	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_13(t *testing.T) {

	numOfItems := int32(9)
	numOfItemsPicked := int32(9)
	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_14(t *testing.T) {

	numOfItems := int32(9)
	numOfItemsPicked := int32(9)
	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_15(t *testing.T) {

	numOfItems := int32(9)
	numOfItemsPicked := int32(1)
	allowRepetitions := false
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_16(t *testing.T) {

	numOfItems := int32(9)
	numOfItemsPicked := int32(1)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt32_17(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_18(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_19(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_20(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_21(t *testing.T) {
	nInt := 5
	rInt := 11
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_22(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_23(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_24(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt32_25(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int32(nInt)
	numOfItemsPicked := int32(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_01(t *testing.T) {

	numOfItems := int64(3)
	numOfItemsPicked := int64(2)
	allowRepetitions := false
	expectedResultStr := "6"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
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

	numOfItems := int64(5)
	numOfItemsPicked := int64(11)
	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_09(t *testing.T) {

	numOfItems := int64(56)
	numOfItemsPicked := int64(5)
	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_10(t *testing.T) {

	numOfItems := int64(9)
	numOfItemsPicked := int64(3)
	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_11(t *testing.T) {

	numOfItems := int64(12)
	numOfItemsPicked := int64(7)
	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_12(t *testing.T) {

	numOfItems := int64(18)
	numOfItemsPicked := int64(8)
	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_13(t *testing.T) {

	numOfItems := int64(9)
	numOfItemsPicked := int64(9)
	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_14(t *testing.T) {

	numOfItems := int64(9)
	numOfItemsPicked := int64(9)
	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_15(t *testing.T) {

	numOfItems := int64(9)
	numOfItemsPicked := int64(1)
	allowRepetitions := false
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_16(t *testing.T) {

	numOfItems := int64(9)
	numOfItemsPicked := int64(1)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsInt64_17(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_18(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_19(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_20(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_21(t *testing.T) {
	nInt := 5
	rInt := 11
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}
}

func TestProbability_PermutationsInt64_22(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}
}

func TestProbability_PermutationsInt64_23(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_24(t *testing.T) {
	nInt := -15
	rInt := 2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsInt64_25(t *testing.T) {
	nInt := 15
	rInt := -2
	numOfItems := int64(nInt)
	numOfItemsPicked := int64(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
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

	numOfItems := uint(5)
	numOfItemsPicked := uint(11)
	allowRepetitions := true
	expectedResultStr := "48828125"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_09(t *testing.T) {

	numOfItems := uint(11)
	numOfItemsPicked := uint(5)
	allowRepetitions := false
	expectedResultStr := "55440"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_10(t *testing.T) {

	numOfItems := uint(11)
	numOfItemsPicked := uint(5)
	allowRepetitions := true
	expectedResultStr := "161051"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_11(t *testing.T) {

	numOfItems := uint(56)
	numOfItemsPicked := uint(5)
	allowRepetitions := false
	expectedResultStr := "458377920"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_12(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(3)
	allowRepetitions := false
	expectedResultStr := "504"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_13(t *testing.T) {

	numOfItems := uint(12)
	numOfItemsPicked := uint(7)
	allowRepetitions := false
	expectedResultStr := "3991680"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_14(t *testing.T) {

	numOfItems := uint(18)
	numOfItemsPicked := uint(8)
	allowRepetitions := false
	expectedResultStr := "1764322560"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_15(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(9)
	allowRepetitions := false
	expectedResultStr := "362880"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_16(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(9)
	allowRepetitions := true
	expectedResultStr := "387420489"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_17(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(1)
	allowRepetitions := false
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_18(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(1)
	allowRepetitions := true
	expectedResultStr := "9"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_19(t *testing.T) {

	numOfItems := uint(9)
	numOfItemsPicked := uint(4)
	allowRepetitions := true
	expectedResultStr := "6561"

	result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err != nil {
		t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
			"numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
			numOfItems, numOfItemsPicked, err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedResultStr != actualNumStr {
		t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
			expectedResultStr, actualNumStr)
	}

}

func TestProbability_PermutationsUint_20(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_21(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := true

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_22(t *testing.T) {
	nInt := 5
	rInt := 11
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
			"However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_23(t *testing.T) {
	nInt := 0
	rInt := 4
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_24(t *testing.T) {
	nInt := 15
	rInt := 0
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from PermutationsUint("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
	}

}

func TestProbability_PermutationsUint_25(t *testing.T) {
	nInt := 0
	rInt := 0
	numOfItems := uint(nInt)
	numOfItemsPicked := uint(rInt)
	allowRepetitions := false

	_, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
			"numOfItems, numOfItemsPicked, allowRepetitions) "+
			"However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
	}

}
