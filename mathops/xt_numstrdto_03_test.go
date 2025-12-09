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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
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
      ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

    return
  }

  if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Decimal Separator Symbols Don't MATCH!\n"+
      "Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
      "Expected numStrDtoResultDecimalSeparator = '%v'\n"+
      "  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
      ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

    return
  }

  if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Thousands Separator Symbols Don't MATCH!\n"+
      "Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
      "Expected numStrDtoResultThousandsSeparator = '%v'\n"+
      "  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
      ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

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

  nStr := "123.45"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "$123.45"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_03(t *testing.T) {

  nStr := "12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "$12,345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_04(t *testing.T) {

  nStr := "12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "$12,345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_05(t *testing.T) {

  nStr := "-12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "($12,345.29)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_06(t *testing.T) {

  nStr := "-12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "($12,345)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_07(t *testing.T) {

  nStr := "-12345"
  // '\U000020ac', // Euro €
  currencySymbol := '€'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(€12,345)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetCurrencySymbol('\U000020ac')

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetCurrencyParen_08(t *testing.T) {

  nStr := "12345.12"
  // '\U000020ac', // Euro €
  currencySymbol := '€'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "€12,345.12"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetCurrencySymbol('\U000020ac')

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetCurrencyParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetDecimal_01(t *testing.T) {

  numStr := "198649257.12345678"

  controlDecimal, err := Decimal{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by controlDecimal = Decimal{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v'",
      numStr, err.Error())
  }

  controlNDto, err := NumStrDto{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by controlNDto = NumStrDto{}.NewNumStr(numStr). "+
      "numStr='%v' Error='%v'",
      numStr, err.Error())
  }

  decActual, err := controlNDto.GetDecimal()

  if err != nil {
    t.Errorf("Error returned by decActual = controlNDto.GetDecimal(). "+
      "Error='%v'", err.Error())
  }

  if numStr != decActual.GetNumStr() {
    t.Errorf("Error: Expected decActual.GetNumStr()='%v'. Instead, "+
      "decActual.GetNumStr()='%v'.",
      numStr, decActual.GetNumStr())
  }

  if !controlDecimal.Equal(decActual) {
    t.Errorf("Error: controlDecimal NOT EQUAL to decActual! "+
      "controlDecimal='%v' decActual='%v'",
      controlDecimal.GetNumStr(), decActual.GetNumStr())
  }
}

func TestNumStrDto_GetIntAry(t *testing.T) {

  numStr := "589627.123456"

  controlNDto, err := NumStrDto{}.NewNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(numStr). "+
      "numStr='%v'  Error='%v'", numStr, err.Error())
  }

  controlIa, err := IntAry{}.NewNumStr(numStr)

  actualIa, err := controlNDto.GetIntAry()

  if err != nil {
    t.Errorf("Error returned by controlNDto.GetIntAryElements(). "+
      "numStr='%v'  Error='%v'", numStr, err.Error())
  }

  if numStr != actualIa.GetNumStr() {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      numStr, actualIa.GetNumStr())
  }

  if !controlIa.Equals(&actualIa) {
    t.Errorf("Error: controlIa NOT EQUAL to actual actualIa! "+
      "controlNDto='%v' nDto='%v'",
      controlNDto.GetNumStr(), actualIa.GetNumStr())
  }

}

func TestNumStrDto_GetNumStr_01(t *testing.T) {

  nStr := "123456.97"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123456.97"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_02(t *testing.T) {

  nStr := "123.45"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123.45"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_03(t *testing.T) {

  nStr := "12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_04(t *testing.T) {

  nStr := "12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_05(t *testing.T) {

  nStr := "12345.1234"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345.1234"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_06(t *testing.T) {

  nStr := "1234567890.25"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "1234567890.25"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_07(t *testing.T) {

  nStr := "-12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-12345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_08(t *testing.T) {

  nStr := "-12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-12345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_09(t *testing.T) {

  nStr := "-123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_10(t *testing.T) {

  nStr := "-0.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-0.123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumStr_11(t *testing.T) {

  nStr := "-1234567890.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-1234567890.123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_01(t *testing.T) {

  nStr := "123456.97"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123456.97"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_02(t *testing.T) {

  nStr := "123.45"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123.45"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_03(t *testing.T) {

  nStr := "12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_04(t *testing.T) {

  nStr := "12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_05(t *testing.T) {

  nStr := "12345.1234"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12345.1234"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_06(t *testing.T) {

  nStr := "1234567890.25"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "1234567890.25"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_07(t *testing.T) {

  nStr := "-12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(12345.29)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_08(t *testing.T) {

  nStr := "-12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(12345)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_09(t *testing.T) {

  nStr := "-123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_10(t *testing.T) {

  nStr := "-0.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(0.123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetNumParen_11(t *testing.T) {

  nStr := "-1234567890.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(1234567890.123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetNumParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}
