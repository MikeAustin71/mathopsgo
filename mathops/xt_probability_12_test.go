package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_PermutationsNumStrDto_01(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_01"

  numOfItemsInt := 3

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "6"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_02(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_02"

  numOfItemsInt := 3

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "9"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_03(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_03"

  numOfItemsInt := 10

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "1000"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_04(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_04"

  numOfItemsInt := 20

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1860480"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_05(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_05"

  numOfItemsInt := 52

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "311875200"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_06(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_06"

  numOfItemsInt := 5

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "125"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_07(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_07"

  numOfItemsInt := 20

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "3200000"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_08(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_08"

  numOfItemsInt := 5

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 11

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "48828125"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_09(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_09"

  numOfItemsInt := 11

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "55440"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_10(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_10"

  numOfItemsInt := 11

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "161051"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_11(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_11"

  numOfItemsInt := 56

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "458377920"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_12(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_12"

  numOfItemsInt := 9

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "504"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_13(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_13"

  numOfItemsInt := 12

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 7

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "3991680"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_14(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_14"

  numOfItemsInt := 18

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 8

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1764322560"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_15(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_15"

  numOfItemsInt := 9

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 9

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "362880"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_16(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_16"

  numOfItemsInt := 9

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 9

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "387420489"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_17(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_17"

  numOfItemsInt := 9

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 1

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "9"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_18(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_18"

  numOfItemsInt := 9

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 1

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "9"

  expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
      "expectedNumStr= '%v'\n"+
      "Error: isOk == false\n\n",
      ePrefix,
      expectedNumStr)

    return
  }

  expectedPrecisionInt := 0

  expectedPrecisionUint := uint(expectedPrecisionInt)

  big10 := big.NewInt(10)

  baseExp := big.NewInt(int64(expectedPrecisionInt))

  expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedNumStrDtoResult, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedNumStrDtoResult.IsValid("Validating expectedNumStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDtoResult.IsValid('Validating expectedNumStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoResultStr, err := expectedNumStrDtoResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedNumStrDtoResultStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedNumStrDtoResultStr \n"+
      "Expected expectedNumStrDtoResultStr = '%v'\n"+
      "  Actual expectedNumStrDtoResultStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedNumStrDtoResultStr)

    return
  }

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  numStrDtoResult, err := new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).PermutationsUint(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
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
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
      "numStrDtoResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultScaleFactorBigInt, err :=\n"+
      "  numStrDtoResult.GetScaleFactor()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  numStrDtoResultBigInt, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigInt, err :=\n"+
      "  numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedNumStr != numStrDtoResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != numStrDtoResultNumStr\n"+
      "Expected numStrDtoResultNumStr = '%v'\n"+
      "  Actual numStrDtoResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, numStrDtoResultNumStr)

    return
  }

  expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDtoResult.EqualTo(numStrDtoResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultNumStrDtosAreEqual, err :=\n"+
      "  expectedNumStrDtoResult.Equal(numStrDtoResult)\n"+
      "expectedNumStrDtoResult= '%v'\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr, err.Error())
    return
  }

  if expectedAndResultNumStrDtosAreEqual == false {
    t.Errorf("%v\n"+
      "Error: numStrDtoResult is INVALID!\n"+
      "Because expectedAndResultNumStrDtosAreEqual == false\n"+
      "Expected numStrDtoResult = '%v'\n"+
      "  Actual numStrDtoResult = '%v'\n\n",
      ePrefix, expectedNumStrDtoResultStr, numStrDtoResultNumStr)

    return
  }

  if expectedPrecisionInt != numStrDtoResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
      "Expected numStrDtoResultPrecisionInt = '%v'\n"+
      "  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != numStrDtoResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
      "Expected numStrDtoResultPrecisionUint = '%v'\n"+
      "  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

    return
  }

  if expectedSignValue != numStrDtoResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != numStrDtoResultSignValue\n"+
      "Expected numStrDtoResultSignValue = '%v'\n"+
      "  Actual numStrDtoResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, numStrDtoResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != numStrDtoResultNumSeps \n"+
      "Expected numStrDtoResultNumSeps = '%v'\n"+
      "  Actual numStrDtoResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
      "Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
      "  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      numStrDtoResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(numStrDtoResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual numStrDtoResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(numStrDtoResultBigInt) != 0\n"+
      "Expected numStrDtoResultBigInt = '%v'\n"+
      "  Actual numStrDtoResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), numStrDtoResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_19(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_19"

  numOfItemsInt := 0

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 4

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_20(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_20"

  numOfItemsInt := 15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 0

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_21(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_2X"

  numOfItemsInt := -15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_22(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_22"

  numOfItemsInt := 15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := -2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_23(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_23"

  numOfItemsInt := 5

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 11

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItems < numStrDtoNumOfItemsChosen\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_24(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_24"

  numOfItemsInt := 0

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 4

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_25(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_25"

  numOfItemsInt := 15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 0

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_26(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_26"

  numOfItemsInt := -15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsNumStrDto_27(t *testing.T) {

  ePrefix := "TestProbability_PermutationsNumStrDto_27"

  numOfItemsInt := 15

  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := -2

  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsPrecisionUint,
      err.Error())
    return
  }

  err = numStrDtoNumOfItems.IsValid("Validating numStrDtoNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItems.IsValid('Validating numStrDtoNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsNumStr, err := numStrDtoNumOfItems.GetNumStr()\n"+
      "numStrDtoNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != numStrDtoNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != numStrDtoNumOfItemsNumStr\n"+
      "Expected numStrDtoNumOfItemsNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, numStrDtoNumOfItemsNumStr)

    return
  }

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = numStrDtoNumOfItemsChosen.IsValid("Validating numStrDtoNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := numStrDtoNumOfItemsChosen.IsValid('Validating numStrDtoNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosenNumStr, err := numStrDtoNumOfItemsChosen.GetNumStr()\n"+
      "numStrDtoNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: numStrDtoNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != numStrDtoNumOfItemsChosenNumStr\n"+
      "Expected numStrDtoNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual numStrDtoNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, numStrDtoNumOfItemsChosenNumStr)

    return
  }

  _, err = new(Probability).PermutationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
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
