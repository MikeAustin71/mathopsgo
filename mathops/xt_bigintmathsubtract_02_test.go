package mathops

import (
  "fmt"
  "testing"
)

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_01"

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39.45",
    "200.1",
    "105.6",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(
      decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_02"

  var err error

  // minuend =   5051
  minuendStr := "5051"

  subtrahendStrs := []string{
    "8000",
    "6051.123456",
    "-30871.25",
    "604.55",
    "9100.123",
    "-115.76",
  }

  expectedStrs := []string{
    "-2949",
    "-1000.123456",
    "35922.25",
    "4446.45",
    "-4049.123",
    "5166.76",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(
      decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_03"

  var err error

  // minuend =   -20051.974578
  minuendStr := "-20051.974578"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-20528.518376",
    "-26103.098034",
    "250819.275422",
    "-35656.5335101",
    "-1007152.097578",
    "94503.785422",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(
      decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_04"

  var err error

  // minuend =   0
  minuendStr := "0"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-476.543798",
    "-6051.123456",
    "270871.25",
    "-15604.5589321",
    "-987100.123",
    "114555.76",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_05"

  var err error

  // minuend =   0
  minuendStr := "98.2"

  subtrahendStrs := []string{
    "0",
    "0.000",
    "0",
    "0",
    "0.00000",
    "0.0",
  }

  expectedStrs := []string{
    "98.2",
    "98.200",
    "98.2",
    "98.2",
    "98.20000",
    "98.2",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalOutputToArray_06"

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39,45",
    "200,1",
    "105,6",
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).\n"+
      "  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]Decimal, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]Decimal, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(Decimal).NewNumStrWithNumSeps(subtrahendStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStrWithNumSeps(subtrahendStrs[%d], usaNumSeps)\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(Validating subtrahendAry[%d])\n"+
        "Validation Error='%v'\n\n",
        ePrefix, i, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(Decimal).NewNumStrWithNumSeps(expectedStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err = new(Decimal).\n"+
        "  NewNumStrWithNumSeps(expectedStrs[%d], &expectedNumSeps)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractDecimalOutputToArray(decMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalOutputToArray(\n"+
      "  decMinuend, subtrahendAry[...])\n"+
      "decMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decMinuend,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].SetNumericSeparatorsDto(expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].SetNumericSeparatorsDto(expectedNumSeps)\n"+
        "result[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, result[k], expectedNumSeps.String(), err.Error())
      return
    }

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "result[%d]= '%v'\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, k, result[k], err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractDecimalSeries_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalSeries_01"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahends := 6

  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
    decMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalSeries(\n"+
      " decMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "decMinuend= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalSeries_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalSeries_02"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahends := 6

  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
    decMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalSeries(\n"+
      " decMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "decMinuend= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalSeries_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalSeries_03"

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahends := 6

  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
    decMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalSeries(\n"+
      " decMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "decMinuend= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalSeries_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalSeries_04"

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahends := 6

  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(Decimal).NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(Decimal).NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(Decimal).NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(Decimal).NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(Decimal).NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(Decimal).NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
    decMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalSeries(\n"+
      " decMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "decMinuend= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractDecimalSeries_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractDecimalSeries_05"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  subtrahend0 := "123.894000"
  subtrahend1 := "67.1"
  subtrahend2 := "93.0"
  subtrahend3 := "-124498.67158"
  subtrahend4 := "647129.57"
  subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decMinuend, err := new(Decimal).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuend, err := new(Decimal).\n"+
      " NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = decMinuend.IsValid("Validating decMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decMinuend.IsValid('Validating decMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decMinuendNumStr, err := decMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decMinuendNumStr, err := decMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != decMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected decMinuendNumStr = '%v'\n"+
      "  Actual decMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, decMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahends := 6

  subtrahendAry := make([]Decimal, lenSubtrahends)

  subtrahendAry[0], err = new(Decimal).NewNumStrWithNumSeps(subtrahend0, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend0, usaNumSeps)\n"+
      "subtrahend0= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(Decimal).NewNumStrWithNumSeps(subtrahend1, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend1, usaNumSeps)\n"+
      "subtrahend1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(Decimal).NewNumStrWithNumSeps(subtrahend2, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend2, usaNumSeps)\n"+
      "subtrahend2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(Decimal).NewNumStrWithNumSeps(subtrahend3, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend3, usaNumSeps)\n"+
      "subtrahend3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(Decimal).NewNumStrWithNumSeps(subtrahend4, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend4, usaNumSeps)\n"+
      "subtrahend4= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(Decimal).NewNumStrWithNumSeps(subtrahend5, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(Decimal).\n"+
      "  NewNumStrWithNumSeps(subtrahend5, usaNumSeps)\n"+
      "subtrahend5= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahend5, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractDecimalSeries(
    decMinuend,
    subtrahendAry[0],
    subtrahendAry[1],
    subtrahendAry[2],
    subtrahendAry[3],
    subtrahendAry[4],
    subtrahendAry[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractDecimalSeries(\n"+
      " decMinuend, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],\n"+
      "  subtrahendAry[3], subtrahendAry[4], subtrahendAry[5])\n"+
      "decMinuend= '%v'\n"+
      "subtrahendAry[0]= '%v'\n"+
      "subtrahendAry[1]= '%v'\n"+
      "subtrahendAry[2]= '%v'\n"+
      "subtrahendAry[3]= '%v'\n"+
      "subtrahendAry[4]= '%v'\n"+
      " subtrahendAry[5]= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decMinuendNumStr, subtrahendAry[0], subtrahendAry[1], subtrahendAry[2],
      subtrahendAry[3], subtrahendAry[4], subtrahendAry[5],
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAry_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAry_01"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99.999"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != iaMinuendNumStr\n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = iaSubtrahend.IsValid("Validating iaSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaSubtrahend.IsValid('Validating iaSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != iaSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != iaSubtrahendNumStr \n"+
      "Expected iaSubtrahendNumStr = '%v'\n"+
      "  Actual iaSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, iaSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractIntAry(
    iaMinuend, iaSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAry(\n"+
      "  iaMinuend, iaSubtrahend)\n"+
      "iaMinuend= '%v'\n"+
      "iaSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAry_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAry_02"

  // minuend = 949321.6712
  minuendStr := "949321.6712"

  // subtrahend = 45678.21
  subtrahendStr := "45678.21"

  // result = 903643.4612
  expectedBigINumStr := "903643.4612"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != iaMinuendNumStr\n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = iaSubtrahend.IsValid("Validating iaSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaSubtrahend.IsValid('Validating iaSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != iaSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != iaSubtrahendNumStr \n"+
      "Expected iaSubtrahendNumStr = '%v'\n"+
      "  Actual iaSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, iaSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAry(\n"+
      "  iaMinuend, iaSubtrahend)\n"+
      "iaMinuend= '%v'\n"+
      "iaSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAry_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAry_03"

  // minuend = -5876458.56789012
  minuendStr := "-5876458.56789012"

  // subtrahend = 847129.876
  subtrahendStr := "847129.876"

  // result = -6723588.44389012
  expectedBigINumStr := "-6723588.44389012"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != iaMinuendNumStr\n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = iaSubtrahend.IsValid("Validating iaSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaSubtrahend.IsValid('Validating iaSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != iaSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != iaSubtrahendNumStr \n"+
      "Expected iaSubtrahendNumStr = '%v'\n"+
      "  Actual iaSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, iaSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAry(\n"+
      "  iaMinuend, iaSubtrahend)\n"+
      "iaMinuend= '%v'\n"+
      "iaSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAry_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAry_04"

  // minuend = -289.673849
  minuendStr := "-289.673849"

  // subtrahend = -14579.012
  subtrahendStr := "-14579.012"

  // result = 14289.338151
  expectedBigINumStr := "14289.338151"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != iaMinuendNumStr\n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahend, err := new(IntAry).NewNumStr(subtrahendStr)\n"+
      "subtrahendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, err.Error())
    return
  }

  err = iaSubtrahend.IsValid("Validating iaSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaSubtrahend.IsValid('Validating iaSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != iaSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != iaSubtrahendNumStr \n"+
      "Expected iaSubtrahendNumStr = '%v'\n"+
      "  Actual iaSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, iaSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAry(\n"+
      "  iaMinuend, iaSubtrahend)\n"+
      "iaMinuend= '%v'\n"+
      "iaSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAry_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAry_05"

  // minuend = 123.32
  minuendStr := "123.32"

  // subtrahend = 23.321
  subtrahendStr := "23.321"

  // result = 99.999
  expectedBigINumStr := "99,999"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != iaMinuendNumStr\n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  iaSubtrahend, err := new(IntAry).NewNumStrWithNumSeps(subtrahendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "subtrahendBiNum, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(subtrahendStr, usaNumSeps)\n"+
      "subtrahendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, subtrahendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaSubtrahend.IsValid("Validating iaSubtrahend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaSubtrahend.IsValid('Validating iaSubtrahend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaSubtrahendNumStr, err := iaSubtrahend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if subtrahendStr != iaSubtrahendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because subtrahendStr != iaSubtrahendNumStr \n"+
      "Expected iaSubtrahendNumStr = '%v'\n"+
      "  Actual iaSubtrahendNumStr = '%v'\n\n",
      ePrefix, subtrahendStr, iaSubtrahendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathSubtract).SubtractIntAry(iaMinuend, iaSubtrahend)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAry(\n"+
      "  iaMinuend, iaSubtrahend)\n"+
      "iaMinuend= '%v'\n"+
      "iaSubtrahend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSubtrahendNumStr,
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryArray_01"

  var err error

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //subtrahend0 := "123.894000"
  //subtrahend1 := "67.1"
  //subtrahend2 := "93.0"
  //subtrahend3 := "-124498.67158"
  //subtrahend4 := "647129.57"
  //subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998.231036"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(IntAry).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractIntAryArray(
    iaMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryArray(\n"+
      "  iaMinuend, subtrahendAry[...])\n"+
      "iaMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryArray_02"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "737.21",
    "9637591.879546",
    "28",
    "5284.9765",
    "-189291837.12",
    "7638932.12398765",
  }

  //subtrahend0 := "737.21"
  //subtrahend1 := "9637591.879546"
  //subtrahend2 := "28"
  //subtrahend3 := "5284.9765"
  //subtrahend4 := "-189291837.12"
  //subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(IntAry).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryArray(\n"+
      "  iaMinuend, subtrahendAry[...])\n"+
      "iaMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryArray_03"

  var err error

  // minuend =   1,718,973,642.1234567
  minuendStr := "1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "737.21",
    "9637591.879546",
    "28",
    "5284.9765",
    "-189291837.12",
    "7638932.12398765",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:  2,757,547,756.287792379
  expectedBigINumStr := "2757547756.287792379"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(IntAry).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryArray(\n"+
      "  iaMinuend, subtrahendAry[...])\n"+
      "iaMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryArray_04"

  var err error

  // minuend =   -1,718,973,642.1234567
  minuendStr := "-1718973642.1234567"

  // subtrahendStrs
  subtrahendStrs := []string{
    "-28934682.721",
    "424.987654321",
    "-987",
    "62.94",
    "-999999999.99999",
    "-9638932.371",
  }

  //subtrahend0 := "-28934682.721"
  //subtrahend1 := "424.987654321"
  //subtrahend2 := "-987"
  //subtrahend3 := "62.94"
  //subtrahend4 := "-999999999.99999"
  //subtrahend5 := "-9638932.371"

  // Result:   -680,399,527.959121021
  expectedBigINumStr := "-680399527.959121021"

  expectedBigINumSign := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  iaMinuend, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(IntAry).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryArray(\n"+
      "  iaMinuend, subtrahendAry[...])\n"+
      "iaMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryArray_05"

  var err error

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  // minuend = 7328941.123456
  minuendStr := "7328941.123456"

  // subtrahendStrs
  subtrahendStrs := []string{
    "123.894000",
    "67.1",
    "93.0",
    "-124498.67158",
    "647129.57",
    "28",
  }

  //subtrahend0 := "123.894000"
  //subtrahend1 := "67.1"
  //subtrahend2 := "93.0"
  //subtrahend3 := "-124498.67158"
  //subtrahend4 := "647129.57"
  //subtrahend5 := "28"

  // result = 6805998.231036
  expectedBigINumStr := "6805998,231036"

  expectedBigINumSign := 1

  iaMinuend, err := new(IntAry).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaMinuend.IsValid("Validating iaMinuend")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.IsValid('Validating iaMinuend')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != iaMinuendNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != decMinuendNumStr \n"+
      "Expected iaMinuendNumStr = '%v'\n"+
      "  Actual iaMinuendNumStr = '%v'\n\n",
      ePrefix, minuendStr, iaMinuendNumStr)

    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid("Validating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid('Validating expectedBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, expectedBigINumberStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSeps, err := expectedBigINum.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  lenSubtrahendsArray := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahendsArray)

  for i := 0; i < lenSubtrahendsArray; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStrWithNumSeps(subtrahendStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(IntAry).\n"+
        "  NewNumStrWithNumSeps(subtrahendStrs[%d], usaNumSeps)\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid(ePrefix)\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // end of loop

  result, err := new(BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryArray(\n"+
      "  iaMinuend, subtrahendAry[...])\n"+
      "iaMinuend= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigINumStr, resultNumStr, err.Error())
    return
  }

  if !expectedEqualsResult {
    t.Errorf("%v\n"+
      "Error: Expected and 'result' values NOT Equal\n"+
      "Because expectedEqualsResult = 'false' \n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Big Int Values NOT Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigInt) != 0 \n"+
      "Expected resultBigInt = '%v'\n"+
      "  Actual resultBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))

    return
  }

  if expectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values NOT Equal\n"+
      "Because expectedBigINumStr != resultNumStr \n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Number Sign Values NOT Equal\n"+
      "Because expectedBigINumSign != resultSignValue \n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != resultNumSeps \n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  if !expectedNumSeps.Equal(expectedBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separators Values NOT Equal\n"+
      "Because expectedNumSeps != expectedBigINumSeps \n"+
      "Expected expectedBigINumSeps = '%v'\n"+
      "  Actual expectedBigINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), expectedBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_01"

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39.45",
    "200.1",
    "105.6",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = subtrahendAry[%d].IsValid()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_02"

  var err error

  // minuend =   5051
  minuendStr := "5051"

  subtrahendStrs := []string{
    "8000",
    "6051.123456",
    "-30871.25",
    "604.55",
    "9100.123",
    "-115.76",
  }

  expectedStrs := []string{
    "-2949",
    "-1000.123456",
    "35922.25",
    "4446.45",
    "-4049.123",
    "5166.76",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_03"

  var err error

  // minuend =   -20051.974578
  minuendStr := "-20051.974578"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-20528.518376",
    "-26103.098034",
    "250819.275422",
    "-35656.5335101",
    "-1007152.097578",
    "94503.785422",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_04"

  var err error

  // minuend =   0
  minuendStr := "0"

  subtrahendStrs := []string{
    "476.543798",
    "6051.123456",
    "-270871.25",
    "15604.5589321",
    "987100.123",
    "-114555.76",
  }

  expectedStrs := []string{
    "-476.543798",
    "-6051.123456",
    "270871.25",
    "-15604.5589321",
    "-987100.123",
    "114555.76",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_05"

  var err error

  // minuend =   0
  minuendStr := "98.2"

  subtrahendStrs := []string{
    "0",
    "0.000",
    "0",
    "0",
    "0.00000",
    "0.0",
  }

  expectedStrs := []string{
    "98.2",
    "98.200",
    "98.2",
    "98.2",
    "98.20000",
    "98.2",
  }

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStr(subtrahendStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStr(subtrahendStrs[%d])\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStr(expectedStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(Decimal).\n"+
        "  NewNumStr(expectedStrs[%d])\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}

func TestBigIntMathSubtract_SubtractIntAryOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathSubtract_SubtractIntAryOutputToArray_06"

  var err error

  // minuend =   100
  minuendStr := "100"

  subtrahendStrs := []string{
    "5",
    "10",
    "30",
    "60.55",
    "-100.1",
    "-5.6",
  }

  expectedStrs := []string{
    "95",
    "90",
    "70",
    "39,45",
    "200,1",
    "105,6",
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  minuendIntAry, err := new(IntAry).NewNumStrWithNumSeps(minuendStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAry, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(minuendStr, usaNumSeps)\n"+
      "minuendStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, minuendStr, usaNumSeps.String(), err.Error())
    return
  }

  err = minuendIntAry.IsValid("Validating minuendIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = minuendIntAry.IsValid('Validating minuendIntAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  minuendIntAryNumStr, err := minuendIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "minuendIntAryNumStr, err := minuendIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if minuendStr != minuendIntAryNumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because minuendStr != minuendIntAryNumStr \n"+
      "Expected minuendIntAryNumStr = '%v'\n"+
      "  Actual minuendIntAryNumStr = '%v'\n\n",
      ePrefix, minuendStr, minuendIntAryNumStr)

    return
  }

  lenSubtrahends := len(subtrahendStrs)

  subtrahendAry := make([]IntAry, lenSubtrahends)

  lenExpectedNumStrsAry := len(expectedStrs)

  if lenExpectedNumStrsAry != lenSubtrahends {
    t.Errorf("%v\n"+
      "Error: Test Data is corrupted!\n"+
      "Because lenExpectedNumStrsAry := len(subtrahendStrs)\n"+
      "Expected Length of subtrahendStrs Array = '%v'\n"+
      "  Actual Length of subtrahendStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenSubtrahends)

    return
  }

  expectedResultsAry := make([]IntAry, lenSubtrahends)

  var expectedResultsNumStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAry[i], err = new(IntAry).NewNumStrWithNumSeps(subtrahendStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAry[%d], err = new(Decimal).\n"+
        "  NewNumStrWithNumSeps(subtrahendStrs[%d], usaNumSeps)\n"+
        "subtrahendStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, subtrahendStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = subtrahendAry[i].IsValid(fmt.Sprintf("Validating subtrahendAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        fmt.Sprintf("err = subtrahendAry[%d].IsValid()\n", i)+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedResultsAry[i], err = new(IntAry).NewNumStrWithNumSeps(expectedStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsAry[%d], err := new(IntAry).\n"+
        "  NewNumStrWithNumSeps(expectedStrs[%d], expectedNumSeps)\n"+
        "expectedStrs[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, expectedStrs[i], expectedNumSeps.String(), err.Error())

      return
    }

    err = expectedResultsAry[i].IsValid(fmt.Sprintf("Validating expectedResultsAry[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err := expectedResultsAry[%d].IsValid()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, expectedStrs[i], err.Error())

      return
    }

    expectedResultsNumStr, err = expectedResultsAry[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultsNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "expectedStrs[%d]= '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, i, expectedStrs[i], err.Error())
      return
    }

    if expectedStrs[i] != expectedResultsNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedStrs[%d] != expectedResultsNumStr \n"+
        "Expected expectedResultsNumStr = '%v'\n"+
        "  Actual expectedResultsNumStr = '%v'\n\n",
        ePrefix, i, expectedStrs[i], expectedResultsNumStr)

      return
    }

  } // End of Loop

  result, err :=
    new(BigIntMathSubtract).SubtractIntAryOutputToArray(minuendIntAry, subtrahendAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathSubtract).SubtractIntAryOutputToArray(\n"+
      "  minuendIntAry, subtrahendAry[...])\n"+
      "minuendIntAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      minuendIntAry,
      err.Error())

    return
  }

  var expectedResultEqualsResult bool

  var expectedResultNumStr, resultNumStr string

  var resultNumSeps NumericSeparatorDto

  for k := 0; k < lenSubtrahends; k++ {

    err = result[k].SetNumericSeparatorsDto(expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].SetNumericSeparatorsDto(\n"+
        "  expectedNumSeps)\n"+
        "expectedNumSeps= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), err.Error())
      return
    }

    err = result[k].IsValid(fmt.Sprintf("Validating result[%d]", k))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = result[%d].IsValid('Validating result[%d]')\n"+
        "Validation Error= '%v'\n\n",
        ePrefix, k, k, err.Error())
      return
    }

    resultNumStr, err = result[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    resultNumSeps, err = result[k].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultNumStr, err = expectedResultsAry[k].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultNumStr, err = expectedResultsAry[%d].GetNumStr()\n"+
        "Error= '%v'\n\n", ePrefix, k, err.Error())
      return
    }

    expectedResultEqualsResult, err = expectedResultsAry[k].Equal(&result[k])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedResultEqualsResult, err =\n"+
        "  expectedResultsAry[%d].Equal(&result[%d])\n"+
        "expectedResultsAry[%d]= '%v'\n"+
        "result[%d]= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, k, k, k, expectedResultNumStr, k, resultNumStr, err.Error())
      return
    }

    if !expectedResultEqualsResult {
      t.Errorf("%v\n"+
        "Error: Expected and 'result' values NOT Equal\n"+
        "Because expectedResultEqualsResult = 'false' \n"+
        "Expected result[%d] = '%v'\n"+
        "  Actual result[%d] = '%v'\n\n",
        ePrefix, k, expectedResultNumStr, k, resultNumStr)

      return
    }

    if expectedResultNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Number String Values NOT Equal\n"+
        "Because expectedResultNumStr != resultNumStr \n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedResultNumStr, resultNumStr)

      return
    }

    if !expectedNumSeps.Equal(resultNumSeps) {
      t.Errorf("%v\n"+
        "Error: Number Sign Values NOT Equal\n"+
        "Because expectedNumSeps != resultNumSeps \n"+
        "Expected resultNumSeps[%d] = '%v'\n"+
        "  Actual resultNumSeps[%d] = '%v'\n\n",
        ePrefix, k, expectedNumSeps.String(), k, resultNumSeps.String())

      return
    }

  } // End of loop

  return
}
