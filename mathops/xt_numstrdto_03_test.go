package mathops

import (
  "math/big"
  "testing"
)

func TestNumStrDto_GetAbsoluteBigInt_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsoluteBigInt_01"

  inputNumberStr := "-123.456"

  expectedInputNumberStr := "123456"

  expectedNumberBigInt, isOk := big.NewInt(0).SetString(expectedInputNumberStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumberBigInt, isOk :=\n"+
      "  big.NewInt(0).SetString(expectedInputNumberStr, 10)\n"+
      "Error - isOk == 'false'\n"+
      "expectedInputNumberStr= '%v'\n\n",
      ePrefix, expectedInputNumberStr)
    return
  }

  numStrDtoResult, err := new(NumStrDto).ParseNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).ParseNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultNumberBigInt, _, err := numStrDtoResult.GetAbsoluteBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "absBigInt, _, err := n1.GetAbsoluteBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedNumberBigInt.Cmp(numStrDtoResultNumberBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Big Int Number Values DON'T MATCH!\n"+
      "Because expectedNumberBigInt.Cmp(numStrDtoResultNumberBigInt) != 0\n"+
      "Expected numStrDtoResultNumberBigInt = '%v'\n"+
      "  Actual numStrDtoResultNumberBigInt = '%v'\n\n",
      ePrefix, expectedNumberBigInt.Text(10), numStrDtoResultNumberBigInt.Text(10))

    return
  }

  return
}

func TestNumStrDto_GetAbsFracRunes_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsFracRunes_01"

  inputNumberStr := "-123.456"

  expectedFractionalRunes := []rune{'4', '5', '6'}

  lenExpectedFractionalRunes := len(expectedFractionalRunes)

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating initial numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating initial numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  lenNumStrDtoResultFracRunes := len(numStrDtoResultFracRunes)

  if lenExpectedFractionalRunes != lenNumStrDtoResultFracRunes {
    t.Errorf("%v\n"+
      "Error: Length of Expected and Actual Fractional Runes DON'T MATCH!\n"+
      "Because lenExpectedFractionalRunes != lenNumStrDtoResultFracRunes\n"+
      "Expected lenNumStrDtoResultFracRunes = '%v'\n"+
      "  Actual lenNumStrDtoResultFracRunes = '%v'\n\n",
      ePrefix, lenExpectedFractionalRunes, lenNumStrDtoResultFracRunes)

    return
  }

  var expectedFractionalRune, numStrDtoResultFracRune rune

  for i := 0; i < lenNumStrDtoResultFracRunes; i++ {

    expectedFractionalRune = expectedFractionalRunes[i]

    numStrDtoResultFracRune = numStrDtoResultFracRunes[i]

    if expectedFractionalRune != numStrDtoResultFracRune {

      t.Errorf("%v\n"+
        "Error: Expected and Actual Fractional Runes DON'T MATCH!\n"+
        "Because expectedFractionalRune != numStrDtoResultFracRune\n"+
        "Expected numStrDtoResultFracRune = '%v'\n"+
        "  Actual numStrDtoResultFracRune = '%v'\n"+
        "Iteration Cycle Number: '%v' \n\n",
        ePrefix,
        string(expectedFractionalRune),
        string(numStrDtoResultFracRune),
        i)

      return
    }

  } // End of 'for' loop

  return
}

func TestNumStrDto_GetAbsFracRunes_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsFracRunes_02"

  inputNumberStr := "123"

  lenExpectedFractionalRunes := 0

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating initial numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating initial numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultFracRunes, err := numStrDtoResult.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultFracRunes, err :=\n"+
      "  numStrDtoResult.GetAbsFracRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  lenNumStrDtoResultFracRunes := len(numStrDtoResultFracRunes)

  if lenExpectedFractionalRunes != lenNumStrDtoResultFracRunes {
    t.Errorf("%v\n"+
      "Error: Length of Expected and Actual Fractional Runes DON'T MATCH!\n"+
      "Because lenExpectedFractionalRunes != lenNumStrDtoResultFracRunes\n"+
      "Expected lenNumStrDtoResultFracRunes = '%v'\n"+
      "  Actual lenNumStrDtoResultFracRunes = '%v'\n\n",
      ePrefix, lenExpectedFractionalRunes, lenNumStrDtoResultFracRunes)

    return
  }

  return
}

func TestNumStrDto_GetAbsIntRunes_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsIntRunes_01"

  inputNumberStr := "-123.456"

  expectedIntRunes := []rune{'1', '2', '3'}

  lenExpectedIntRunes := len(expectedIntRunes)

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating initial numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating initial numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultIntRunes, err :=\n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  lenNumStrDtoResultIntRunes := len(numStrDtoResultIntRunes)

  if lenExpectedIntRunes != lenNumStrDtoResultIntRunes {
    t.Errorf("%v\n"+
      "Error: Length of Expected and Actual Integer Runes DON'T MATCH!\n"+
      "Because lenExpectedIntRunes != lenNumStrDtoResultIntRunes\n"+
      "Expected lenNumStrDtoResultIntRunes = '%v'\n"+
      "  Actual lenNumStrDtoResultIntRunes = '%v'\n\n",
      ePrefix, lenExpectedIntRunes, lenNumStrDtoResultIntRunes)

    return
  }

  var expectedIntRune, numStrDtoResultIntRune rune

  for i := 0; i < lenNumStrDtoResultIntRunes; i++ {

    expectedIntRune = expectedIntRunes[i]

    numStrDtoResultIntRune = numStrDtoResultIntRunes[i]

    if expectedIntRune != numStrDtoResultIntRune {

      t.Errorf("%v\n"+
        "Error: Expected and Actual Integer Runes DON'T MATCH!\n"+
        "Because expectedIntRune != numStrDtoResultIntRune\n"+
        "Expected numStrDtoResultIntRune = '%v'\n"+
        "  Actual numStrDtoResultIntRune = '%v'\n"+
        "Iteration Cycle Number: '%v' \n\n",
        ePrefix,
        string(expectedIntRune),
        string(numStrDtoResultIntRune),
        i)

      return
    }

  }

  return
}

func TestNumStrDto_GetAbsIntRunes_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsIntRunes_02"

  inputNumberStr := "-123"

  expectedIntRunes := []rune{'1', '2', '3'}

  lenExpectedIntRunes := len(expectedIntRunes)

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating initial numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating initial numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultIntRunes, err :=\n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  lenNumStrDtoResultIntRunes := len(numStrDtoResultIntRunes)

  if lenExpectedIntRunes != lenNumStrDtoResultIntRunes {
    t.Errorf("%v\n"+
      "Error: Length of Expected and Actual Integer Runes DON'T MATCH!\n"+
      "Because lenExpectedIntRunes != lenNumStrDtoResultIntRunes\n"+
      "Expected lenNumStrDtoResultIntRunes = '%v'\n"+
      "  Actual lenNumStrDtoResultIntRunes = '%v'\n\n",
      ePrefix, lenExpectedIntRunes, lenNumStrDtoResultIntRunes)

    return
  }

  var expectedIntRune, numStrDtoResultIntRune rune

  for i := 0; i < lenNumStrDtoResultIntRunes; i++ {

    expectedIntRune = expectedIntRunes[i]

    numStrDtoResultIntRune = numStrDtoResultIntRunes[i]

    if expectedIntRune != numStrDtoResultIntRune {

      t.Errorf("%v\n"+
        "Error: Expected and Actual Integer Runes DON'T MATCH!\n"+
        "Because expectedIntRune != numStrDtoResultIntRune\n"+
        "Expected numStrDtoResultIntRune = '%v'\n"+
        "  Actual numStrDtoResultIntRune = '%v'\n"+
        "Iteration Cycle Number: '%v' \n\n",
        ePrefix,
        string(expectedIntRune),
        string(numStrDtoResultIntRune),
        i)

      return
    }

  }

  return
}

func TestNumStrDto_GetAbsIntRunes_03(t *testing.T) {

  ePrefix := "TestNumStrDto_GetAbsIntRunes_03"

  inputNumberStr := "123456789.7865"

  expectedIntRunes := []rune{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

  lenExpectedIntRunes := len(expectedIntRunes)

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating initial numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating initial numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Input and Actual Number Strings ARE NOT EQUAL!\n"+
      "numStrDtoResult intitialization FAILED!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultIntRunes, err := numStrDtoResult.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultIntRunes, err :=\n"+
      "  numStrDtoResult.GetAbsIntRunes()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  lenNumStrDtoResultIntRunes := len(numStrDtoResultIntRunes)

  if lenExpectedIntRunes != lenNumStrDtoResultIntRunes {
    t.Errorf("%v\n"+
      "Error: Length of Expected and Actual Integer Runes DON'T MATCH!\n"+
      "Because lenExpectedIntRunes != lenNumStrDtoResultIntRunes\n"+
      "Expected lenNumStrDtoResultIntRunes = '%v'\n"+
      "  Actual lenNumStrDtoResultIntRunes = '%v'\n\n",
      ePrefix, lenExpectedIntRunes, lenNumStrDtoResultIntRunes)

    return
  }

  var expectedIntRune, numStrDtoResultIntRune rune

  for i := 0; i < lenNumStrDtoResultIntRunes; i++ {

    expectedIntRune = expectedIntRunes[i]

    numStrDtoResultIntRune = numStrDtoResultIntRunes[i]

    if expectedIntRune != numStrDtoResultIntRune {

      t.Errorf("%v\n"+
        "Error: Expected and Actual Integer Runes DON'T MATCH!\n"+
        "Because expectedIntRune != numStrDtoResultIntRune\n"+
        "Expected numStrDtoResultIntRune = '%v'\n"+
        "  Actual numStrDtoResultIntRune = '%v'\n"+
        "Iteration Cycle Number: '%v' \n\n",
        ePrefix,
        string(expectedIntRune),
        string(numStrDtoResultIntRune),
        i)

      return
    }

  }

  return
}

func TestNumStrDto_GetBigIntNum_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetBigIntNum_01"

  inputNumberInt64 := int64(123456123456)

  expectedNumberStr := "123456.123456"

  inputNumPrecisionUint := uint(6)

  bigI := big.NewInt(inputNumberInt64)

  expectedBigINum, err := new(BigIntNum).NewBigInt(bigI, inputNumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewBigInt(bigI, inputNumPrecisionUint)\n"+
      "bigI= '%v'\n"+
      "inputNumPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigI.Text(10),
      inputNumPrecisionUint,
      err.Error())

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
      "expectedBigINumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != expectedBigINumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual BigI Number Strings ARE NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumberStr \n"+
      "Expected expectedBigINumberStr = '%v'\n"+
      "  Actual expectedBigINumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, expectedBigINumberStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewBigInt(bigI, inputNumPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewBigInt(bigI, inputNumPrecisionUint)\n"+
      "bigI= '%v'\n"+
      "inputNumPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigI.Text(10),
      inputNumPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResult Number String Values NOT Equal\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr \n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINum, err := numStrDtoResult.GetBigIntNum()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  err = numStrDtoResultBigINum.IsValid("Validating numStrDtoResultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " err = numStrDtoResultBigINum.IsValid('Validating numStrDtoResultBigINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigINumNumberStr, err := numStrDtoResultBigINum.GetBigIntNum()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigINumNumberStr, err :=\n"+
      "  numStrDtoResultBigINum.GetBigIntNum()\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultBigINumNumberStr, err.Error())
    return
  }

  expectedAndnumStrDtoResultBigIntAreEqual, err := expectedBigINum.Equal(numStrDtoResultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndnumStrDtoResultBigIntAreEqual, err :=\n"+
      "  expectedBigINum.Equal(numStrDtoResultBigINum)\n"+
      "expectedBigINum= '%v'\n"+
      "numStrDtoResultBigINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedBigINumberStr,
      numStrDtoResultBigINumNumberStr,
      err.Error())

    return
  }

  if !expectedAndnumStrDtoResultBigIntAreEqual {
    t.Errorf("%v\n"+
      "Error: Expected and numStrDtoResultBigInt objects ARE NOT EQUAL!\n"+
      "Because expectedAndnumStrDtoResultBigIntAreEqual == false\n"+
      "Expected expectedAndnumStrDtoResultBigIntAreEqual = '%v'\n"+
      "  Actual expectedAndnumStrDtoResultBigIntAreEqual = '%v'\n\n",
      ePrefix, true, false)

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_01"

  inputNumberStr := "123456.97"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123456.97"

  expectedCurrencyStr := "$123,456.97"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_02"

  inputNumberStr := "123.45"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedCurrencyStr := "$123.45"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_03(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_03"

  inputNumberStr := "12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.29"

  expectedCurrencyStr := "$12,345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_04(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_04"

  inputNumberStr := "12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345"

  expectedCurrencyStr := "$12,345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_05(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_05"

  inputNumberStr := "12345.1234"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.1234"

  expectedCurrencyStr := "$12,345.1234"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_06(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_06"

  inputNumberStr := "-12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.1234"

  expectedCurrencyStr := "-$12,345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_07(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_07"

  inputNumberStr := "-12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedCurrencyStr := "-$12,345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_08(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_08"

  inputNumberStr := "-12345"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedCurrencyStr := "-€12,345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_09(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_09"

  inputNumberStr := "12345.12"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedCurrencyStr := "€12,345.12"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_10(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_10"

  inputNumberStr := "-12345"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "-12345"

  expectedCurrencyStr := "-€12 345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_11(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_11"

  inputNumberStr := "-12345.35"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "-12345,35"

  expectedCurrencyStr := "-€12 345,35"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_12(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_12"

  inputNumberStr := "12345.35"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "12345,35"

  expectedCurrencyStr := "€12 345,35"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_13(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_13"

  inputNumberStr := "12345"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "12345"

  expectedCurrencyStr := "€12 345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_14(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_14"

  inputNumberStr := "-12345.12"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "-12345,12"

  expectedCurrencyStr := "€12 345,12"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyStr_15(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyStr_15"

  inputNumberStr := "12345"

  // '\U000020ac', // Euro €
  // currencySymbol := '€'
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := ','

  expectedThousandsSeparator := ' '

  expectedNumberStr := "12345"

  expectedCurrencyStr := "€12 345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyStr, err := numStrDtoResult.GetCurrencyStr()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyStr != numStrDtoCurrencyStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyStr, numStrDtoCurrencyStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_01"

  inputNumberStr := "123456.97"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123456.97"

  expectedCurrencyParenStr := "$123,456.97"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_02"

  inputNumberStr := "123.45"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedCurrencyParenStr := "$123.45"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_03(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_03"

  inputNumberStr := "12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedCurrencyParenStr := "$12,345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_04(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_04"

  inputNumberStr := "12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedCurrencyParenStr := "$12,345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_05(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_05"

  inputNumberStr := "-12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345.29"

  expectedCurrencyParenStr := "($12,345.29)"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_06(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_06"

  inputNumberStr := "-12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedCurrencyParenStr := "($12,345)"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_07(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_07"

  inputNumberStr := "-12345"

  // '\U000020ac', // Euro €
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedCurrencyParenStr := "(€12,345)"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetCurrencyParen_08(t *testing.T) {

  ePrefix := "TestNumStrDto_GetCurrencyParen_08"

  inputNumberStr := "12345.12"

  // '\U000020ac', // Euro €
  expectedCurrencySymbol := '\U000020ac'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.12"

  expectedCurrencyParenStr := "€12,345.12"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoCurrencyParenStr, err := numStrDtoResult.GetCurrencyParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedCurrencyParenStr != numStrDtoCurrencyParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Strings Don't MATCH!\n"+
      "Because!!!!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "  Actual fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedCurrencyParenStr, numStrDtoCurrencyParenStr)

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetDecimal_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetDecimal_01"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  inputNumberStr := "198649257.12345678"

  expectedNumberStr := inputNumberStr

  expectedPrecisionInt := 8

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decimalExpected, err := new(Decimal).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalExpected, err := new(Decimal).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = decimalExpected.IsValid("Validating decimalExpected-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decimalExpected.IsValid('Validating decimalExpected-originalNumberStr')\n"+
      "decimalExpected set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalExpectedNumberStr, err := decimalExpected.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalExpectedNumberStr, err := decimalExpected.GetNumStr()\n"+
      "decimalExpected set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != decimalExpectedNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because inputNumberStr != decimalExpectedNumberStr \n"+
      "Expected decimalExpectedNumberStr = '%v'\n"+
      "  Actual decimalExpectedNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, decimalExpectedNumberStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and actual number strings DO NOT MATCH!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  decimalResult, err := numStrDtoResult.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResult, err := numStrDtoResult.GetDecimal()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  err = decimalResult.IsValid("Validating decimalResult-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decimalResult.IsValid('Validating decimalResult-originalNumberStr')\n"+
      "decimalResult set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalResultNumberStr, err := decimalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResultNumberStr, err := decimalResult.GetNumStr()\n"+
      "decimalResult set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalResultPrecisionInt, err := decimalResult.GetPrecision()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResultPrecisionInt, err :=\n"+
      "  decimalResult.GetPrecision()\n"+
      "decimalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decimalResultNumberStr, err.Error())
    return
  }

  decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResultPrecisionUint, err :=\n"+
      "  decimalResult.GetPrecisionUint()\n"+
      "decimalResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decimalResultNumberStr, err.Error())
    return
  }

  decimalResultSignValue, err := decimalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResultSignValue, err := decimalResult.GetSign()\n"+
      "decimalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decimalResultNumberStr, err.Error())
    return
  }

  decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
      "decimalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decimalResultNumberStr, err.Error())
    return
  }

  decimalExpectedEqualsDecimalResult, err := decimalExpected.Equal(decimalResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalExpectedEqualsDecimalResult, err :=\n"+
      "  decimalExpected.Equal(decimalResult)\n"+
      "decimalExpected= '%v'\n"+
      "decimalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalExpectedNumberStr,
      decimalResultNumberStr,
      err.Error())

    return
  }

  if !decimalExpectedEqualsDecimalResult {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal objects ARE NOT EQUAL!\n"+
      "Because decimalExpectedEqualsDecimalResult == false\n"+
      "Expected decimalResult = '%v'\n"+
      "  Actual decimalResult = '%v'\n\n",
      ePrefix, decimalExpectedNumberStr, decimalResultNumberStr)

    return
  }

  if expectedNumberStr != decimalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and decimalResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != decimalResultNumberStr\n"+
      "Expected decimalResultNumberStr = '%v'\n"+
      "  Actual decimalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decimalResultNumberStr)

    return
  }

  if expectedPrecisionInt != decimalResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & decimalResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
      "Expected decimalResultPrecisionInt = '%v'\n"+
      "  Actual decimalResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != decimalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & decimalResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
      "Expected decimalResultPrecisionUint = '%v'\n"+
      "  Actual decimalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

    return
  }

  if expectedSignValue != decimalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & decimalResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != decimalResultSignValue\n"+
      "Expected decimalResultSignValue = '%v'\n"+
      "  Actual decimalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, decimalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decimalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decimalResultNumSeps \n"+
      "Expected decimalResultNumSeps = '%v'\n"+
      "  Actual decimalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetIntAry(t *testing.T) {

  ePrefix := "TestNumStrDto_GetDecimal_01"

  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  inputNumberStr := "589627.123456"

  expectedNumberStr := inputNumberStr

  expectedPrecisionInt := 6

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryExpected, err := new(IntAry).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExpected, err := new(IntAry).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = intAryExpected.IsValid("Validating intAryExpected-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryExpected.IsValid('Validating intAryExpected-originalNumberStr')\n"+
      "intAryExpected set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryExpectedNumberStr, err := intAryExpected.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExpectedNumberStr, err := intAryExpected.GetNumStr()\n"+
      "intAryExpected set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != intAryExpectedNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because inputNumberStr != intAryExpectedNumberStr \n"+
      "Expected intAryExpectedNumberStr = '%v'\n"+
      "  Actual intAryExpectedNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, intAryExpectedNumberStr)

    return
  }

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if inputNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and actual number strings DO NOT MATCH!\n"+
      "Because inputNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, inputNumberStr, numStrDtoResultNumberStr)

    return
  }

  intAryResult, err := numStrDtoResult.GetIntAry()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := numStrDtoResult.GetIntAry()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  err = intAryResult.IsValid("Validating intAryResult-originalNumberStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryResult.IsValid('Validating intAryResult-originalNumberStr')\n"+
      "intAryResult set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumberStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultPrecisionInt := intAryResult.GetPrecision()

  intAryResultPrecisionUint, err := intAryResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultPrecisionUint, err :=\n"+
      "  intAryResult.GetPrecisionUint()\n"+
      "intAryResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
    return
  }

  intAryExpectedEqualsDecimalResult, err := intAryExpected.Equal(&intAryResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryExpectedEqualsDecimalResult, err :=\n"+
      "  intAryExpected.Equal(intAryResult)\n"+
      "intAryExpected= '%v'\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryExpectedNumberStr,
      intAryResultNumberStr,
      err.Error())

    return
  }

  if !intAryExpectedEqualsDecimalResult {
    t.Errorf("%v\n"+
      "Error: Expected and Actual IntAry objects ARE NOT EQUAL!\n"+
      "Because intAryExpectedEqualsDecimalResult == false\n"+
      "Expected intAryResult = '%v'\n"+
      "  Actual intAryResult = '%v'\n\n",
      ePrefix, intAryExpectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedNumberStr != intAryResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and intAryResult Number Strings ARE NOT EQUAL!\n"+
      "Because expectedNumberStr != intAryResultNumberStr\n"+
      "Expected intAryResultNumberStr = '%v'\n"+
      "  Actual intAryResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryResultNumberStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_01"

  inputNumberStr := "123456.97"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123456.97"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_02"

  inputNumberStr := "123.45"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_03(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_03"

  inputNumberStr := "12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_04(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_04"

  inputNumberStr := "12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_05(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_05"

  inputNumberStr := "12345.1234"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.1234"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_06(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_06"

  inputNumberStr := "1234567890.25"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "1234567890.25"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_07(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_07"

  inputNumberStr := "-12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_08(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_08"

  inputNumberStr := "-12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_09(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_09"

  inputNumberStr := "-123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-123"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_10(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_10"

  inputNumberStr := "-0.123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-0.123"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumStr_11(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumStr_11"

  inputNumberStr := "-1234567890.123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-1234567890.123"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  return
}

func TestNumStrDto_GetNumParen_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_01"

  inputNumberStr := "123456.97"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123456.97"

  expectedNumberParenStr := "123456.97"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_02(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_02"

  inputNumberStr := "123.45"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "123.45"

  expectedNumberParenStr := "123.45"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_03(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_03"

  inputNumberStr := "12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.29"

  expectedNumberParenStr := "12345.29"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_04(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_04"

  inputNumberStr := "12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345"

  expectedNumberParenStr := "12345"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_05(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_05"

  inputNumberStr := "12345.1234"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "12345.1234"

  expectedNumberParenStr := "12345.1234"

  expectedPrecisionInt := 4

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_06(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_06"

  inputNumberStr := "1234567890.25"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "1234567890.25"

  expectedNumberParenStr := "1234567890.25"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_07(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_07"

  inputNumberStr := "-12345.29"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-12345.29"

  expectedNumberParenStr := "(12345.29)"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_08(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_08"

  // Note: This test uses
  //  new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &inputNumSeps)

  inputNumberStr := "-12345"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  inputNumSeps := NumericSeparatorDto{}

  inputNumSeps.DecimalSeparator = expectedDecimalSeparator

  inputNumSeps.ThousandsSeparator = expectedThousandsSeparator

  inputNumSeps.CurrencySymbol = expectedCurrencySymbol

  expectedNumberStr := "-12345"

  expectedNumberParenStr := "(12345)"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &inputNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStrWithNumSeps(inputNumberStr, &inputNumSeps)\n"+
      "inputNumberStr= '%v'\n"+
      "inputNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, inputNumSeps.String(), err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_09(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_09"

  inputNumberStr := "-123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-123"

  expectedNumberParenStr := "(123)"

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_10(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_10"

  inputNumberStr := "-0.123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-0.123"

  expectedNumberParenStr := "(0.123)"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_11(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_11"

  inputNumberStr := "-1234567890.123"

  expectedCurrencySymbol := '$'

  expectedDecimalSeparator := '.'

  expectedThousandsSeparator := ','

  expectedNumberStr := "-1234567890.123"

  expectedNumberParenStr := "(1234567890.123)"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

  expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

  expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

  numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparators(\n"+
      "  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
      "decimalSeparator= '%v'\n"+
      "thousandsSeparator= '%v'\n"+
      "currencySymbol= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      string(expectedDecimalSeparator),
      string(expectedThousandsSeparator),
      string(expectedCurrencySymbol),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because currencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedCurrencySymbol), string(numStrDtoResultCurrencySymbol))

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedDecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}

func TestNumStrDto_GetNumParen_12(t *testing.T) {

  ePrefix := "TestNumStrDto_GetNumParen_12"

  inputNumberStr := "-1234567890.123"

  inputNumSeps := NumericSeparatorDto{}

  inputNumSeps.DecimalSeparator = '.'

  inputNumSeps.ThousandsSeparator = ','

  inputNumSeps.CurrencySymbol = '$'

  expectedNumberStr := "-1234567890,123"

  expectedNumberParenStr := "(1234567890,123)"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  //frenchDecSeparator := ','
  //frenchThousandsSeparator := ' '
  // '\U000020ac'
  //frenchCurrencySymbol := '€'

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = ','

  expectedNumSeps.ThousandsSeparator = ' '

  // '\U000020ac', // Euro €
  expectedNumSeps.CurrencySymbol = '\U000020ac'

  numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &inputNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(NumStrDto).\n"+
      "  NewNumStrWithNumSeps(inputNumberStr, &inputNumSeps)\n"+
      "inputNumberStr= '%v'\n"+
      "inputNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, inputNumSeps.String(), err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  err = numStrDtoResult.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

  numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultPrecisionUint, err :=\n"+
      "  numStrDtoResult.GetPrecisionUint()\n"+
      "numStrDtoResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

  numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

  numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

  numStrDtoResultNumParenStr, err := numStrDtoResult.GetNumParen()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumParenStr, err :=\n"+
      "  numStrDtoResult.GetNumParen()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedNumSeps.CurrencySymbol != numStrDtoResultCurrencySymbol {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Currency Symbols Don't MATCH!\n"+
      "Because expectedNumSeps.CurrencySymbol != numStrDtoResultCurrencySymbol\n"+
      "Expected numStrDtoResultCurrencySymbol = '%v'\n"+
      "  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
      ePrefix, string(expectedNumSeps.CurrencySymbol), numStrDtoResultCurrencySymbol)

    return
  }

  if expectedNumSeps.DecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedNumSeps.DecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, string(expectedNumSeps.DecimalSeparator), string(numStrDtoResultDecimalSeparator))

    return
  }

  if expectedNumSeps.ThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedNumSeps.ThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, string(expectedNumSeps.ThousandsSeparator), string(numStrDtoResultThousandsSeparator))

    return
  }

  if expectedNumberStr != numStrDtoResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual number strings DO NOT MATCH!\n"+
      "Because expectedNumberStr != numStrDtoResultNumberStr\n"+
      "Expected numStrDtoResultNumberStr = '%v'\n"+
      "  Actual numStrDtoResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedNumberStr, numStrDtoResultNumParenStr)
  }

  if expectedNumberParenStr != numStrDtoResultNumParenStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Parenthesis String DON'T MATCH!\n"+
      "Because expectedNumberParenStr != numStrDtoResultNumParenStr\n"+
      "Expected numStrDtoResultNumParenStr = '%v'\n"+
      "  Actual numStrDtoResultNumParenStr = '%v'\n\n",
      ePrefix, expectedNumberParenStr, numStrDtoResultNumParenStr)

    return
  }

  return
}
