package mathops

import "testing"



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