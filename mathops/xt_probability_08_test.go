package mathops

import (
  "math/big"
  "testing"
)

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

  numOfItems := BigIntNum{}.NewIntExponent(56, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(5, 0)

  allowRepetitions := false
  expectedResultStr := "458377920"

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

func TestProbability_PermutationsBigIntNum_10(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(9, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(3, 0)

  allowRepetitions := false
  expectedResultStr := "504"

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

func TestProbability_PermutationsBigIntNum_11(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(12, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(7, 0)

  allowRepetitions := false
  expectedResultStr := "3991680"

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

func TestProbability_PermutationsBigIntNum_12(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(18, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(8, 0)

  allowRepetitions := false
  expectedResultStr := "1764322560"

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

func TestProbability_PermutationsBigIntNum_13(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(9, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(9, 0)

  allowRepetitions := false
  expectedResultStr := "362880"

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

func TestProbability_PermutationsBigIntNum_14(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(9, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(9, 0)

  allowRepetitions := true
  expectedResultStr := "387420489"

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

func TestProbability_PermutationsBigIntNum_15(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(9, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(1, 0)

  allowRepetitions := false
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

func TestProbability_PermutationsBigIntNum_16(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(9, 0)
  numOfItemsPicked := BigIntNum{}.NewIntExponent(1, 0)

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

func TestProbability_PermutationsBigIntNum_17(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_18(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_19(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_20(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_21(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_22(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_23(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_24(t *testing.T) {
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

func TestProbability_PermutationsBigIntNum_25(t *testing.T) {
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

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := Decimal{}.NewInt(1, 0)

  allowRepetitions := false
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

func TestProbability_PermutationsDecimal_16(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := Decimal{}.NewInt(1, 0)

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

func TestProbability_PermutationsDecimal_17(t *testing.T) {
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

func TestProbability_PermutationsDecimal_18(t *testing.T) {
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

func TestProbability_PermutationsDecimal_19(t *testing.T) {
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

func TestProbability_PermutationsDecimal_20(t *testing.T) {
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

func TestProbability_PermutationsDecimal_21(t *testing.T) {
  nInt := 5
  rInt := 11

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsDecimal(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsDecimal("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsDecimal_22(t *testing.T) {
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

func TestProbability_PermutationsDecimal_23(t *testing.T) {
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

func TestProbability_PermutationsDecimal_24(t *testing.T) {
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

func TestProbability_PermutationsDecimal_25(t *testing.T) {
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
