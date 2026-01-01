package mathops

import "testing"

func TestProbability_PermutationsNumStrDto_01(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(3, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(2, 0)
  allowRepetitions := false
  expectedResultStr := "6"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_02(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(3, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(2, 0)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_03(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(10, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(3, 0)
  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_04(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(20, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_05(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(52, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_06(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(5, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(3, 0)
  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_07(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(20, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_08(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(5, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(11, 0)
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_09(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(11, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := false
  expectedResultStr := "55440"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_10(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(11, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := true
  expectedResultStr := "161051"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_11(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(56, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_12(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(9, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(3, 0)
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_13(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(12, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(7, 0)
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_14(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(18, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(8, 0)
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_15(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(9, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(9, 0)
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_16(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(9, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(9, 0)
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_17(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(9, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(1, 0)
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_18(t *testing.T) {

  numOfItems := NumStrDto{}.NewInt(9, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(1, 0)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumStrDto_19(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_20(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_21(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_22(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_23(t *testing.T) {
  nInt := 5
  rInt := 11
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_24(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_25(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_26(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumStrDto_27(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := NumStrDto{}.NewInt(nInt, 0)
  numOfItemsPicked := NumStrDto{}.NewInt(rInt, 0)
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumStrDto(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsNumberStr_01(t *testing.T) {

  numOfItems := "3"
  numOfItemsPicked := "2"
  allowRepetitions := false
  expectedResultStr := "6"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_02(t *testing.T) {

  numOfItems := "3"
  numOfItemsPicked := "2"
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_03(t *testing.T) {

  numOfItems := "10"
  numOfItemsPicked := "3"
  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_04(t *testing.T) {

  numOfItems := "20"
  numOfItemsPicked := "5"
  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_05(t *testing.T) {

  numOfItems := "52"
  numOfItemsPicked := "5"
  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_06(t *testing.T) {

  numOfItems := "5"
  numOfItemsPicked := "3"
  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_07(t *testing.T) {

  numOfItems := "20"
  numOfItemsPicked := "5"
  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsNumberStr_08(t *testing.T) {

  numOfItems := "5"
  numOfItemsPicked := "11"
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_09(t *testing.T) {

  numOfItems := "11"
  numOfItemsPicked := "5"
  allowRepetitions := false
  expectedResultStr := "55440"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_10(t *testing.T) {

  numOfItems := "11"
  numOfItemsPicked := "5"
  allowRepetitions := true
  expectedResultStr := "161051"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_11(t *testing.T) {

  numOfItems := "56"
  numOfItemsPicked := "5"
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_12(t *testing.T) {

  numOfItems := "9"
  numOfItemsPicked := "3"
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_13(t *testing.T) {

  numOfItems := "12"
  numOfItemsPicked := "7"
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_14(t *testing.T) {

  numOfItems := "18"
  numOfItemsPicked := "8"
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_15(t *testing.T) {

  numOfItems := "9"
  numOfItemsPicked := "9"
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_16(t *testing.T) {

  numOfItems := "9"
  numOfItemsPicked := "9"
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_17(t *testing.T) {

  numOfItems := "9"
  numOfItemsPicked := "1"
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_18(t *testing.T) {

  numOfItems := "9"
  numOfItemsPicked := "1"
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsNumberStr_19(t *testing.T) {

  numOfItems := "0"
  numOfItemsPicked := "4"
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_20(t *testing.T) {

  numOfItems := "15"
  numOfItemsPicked := "0"
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_21(t *testing.T) {

  numOfItems := "-15"
  numOfItemsPicked := "2"
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_22(t *testing.T) {

  numOfItems := "15"
  numOfItemsPicked := "-2"
  allowRepetitions := true

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_23(t *testing.T) {

  numOfItems := "5"
  numOfItemsPicked := "11"
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_24(t *testing.T) {

  numOfItems := "0"
  numOfItemsPicked := "4"
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}..PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_25(t *testing.T) {

  numOfItems := "15"
  numOfItemsPicked := "0"
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_26(t *testing.T) {

  numOfItems := "-15"
  numOfItemsPicked := "2"
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}

func TestProbability_PermutationsNumberStr_27(t *testing.T) {

  numOfItems := "15"
  numOfItemsPicked := "-2"
  allowRepetitions := false

  _, err := Probability{}.PermutationsNumberStr(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsNumberStr("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", numOfItems, numOfItemsPicked)
  }

}
