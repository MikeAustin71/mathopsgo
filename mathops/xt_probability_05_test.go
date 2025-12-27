package mathops

import (
  "fmt"
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_CombinationsNumStrDto_01(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_01"

  numOfItemsInt := 16

  numOfItemsChosenInt := 3

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "560"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_02(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_02"

  numOfItemsInt := 16

  numOfItemsChosenInt := 12

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1820"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_03(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_03"

  numOfItemsInt := 52

  numOfItemsChosenInt := 5

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "2598960"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_04(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_04"

  numOfItemsInt := 52

  numOfItemsChosenInt := 26

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "495918532948104"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_05(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_05"

  numOfItemsInt := 18

  numOfItemsChosenInt := 7

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "31824"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_06(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_06"

  numOfItemsInt := 22

  numOfItemsChosenInt := 5

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "26334"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_07(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_07"

  numOfItemsInt := 56

  numOfItemsChosenInt := 5

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "3819816"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_08(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_08"

  numOfItemsInt := 56

  numOfItemsChosenInt := 5

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "3819816"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_09(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_02"

  numOfItemsInt := 25

  numOfItemsChosenInt := 25

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_10(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_02"

  numOfItemsInt := 25

  numOfItemsChosenInt := 1

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "25"

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

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  numStrDtoResult, err := new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err := new(Probability).CombinationsNumStrDto(\n"+
      "numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
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

func TestProbability_CombinationsNumStrDto_11(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_11"

  numOfItemsInt := 26

  numOfItemsChosenInt := 52

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  _, err = new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numStrDtoNumOfItems < numStrDtoNumOfItemsChosen\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsNumStrDto_12(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_12"

  numOfItemsInt := 52

  numOfItemsChosenInt := 0

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  _, err = new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsNumStrDto_13(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_13"

  numOfItemsInt := 0

  numOfItemsChosenInt := 26

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  _, err = new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsNumStrDto_14(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_14"

  numOfItemsInt := -52

  numOfItemsChosenInt := 26

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  _, err = new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numStrDtoNumOfItems <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsNumStrDto_15(t *testing.T) {

  ePrefix := "TestProbability_CombinationsNumStrDto_15"

  numOfItemsInt := 52

  numOfItemsChosenInt := -26

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItems, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsInt, 0)\n"+
      "numOfItemsInt= '%v'\n"+
      "precision= '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsIntStr, err.Error())
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

  numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "precision = '0'\n"+
      "Error= '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, err.Error())
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

  _, err = new(Probability).CombinationsNumStrDto(numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsNumStrDto(\n"+
      "  numStrDtoNumOfItems, numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numStrDtoNumOfItemsChosen <= 0\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsNumStrDto_16(t *testing.T) {
  numOfItemsInt := 5
  numOfItemsChosenInt := 3
  expectedResultStr := "35"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_17(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 11
  expectedResultStr := "705432"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_18(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 2
  expectedResultStr := "351"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_19(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 24
  expectedResultStr := "63205303218876"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_20(t *testing.T) {
  numOfItemsInt := 10
  numOfItemsChosenInt := 14
  expectedResultStr := "817190"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_21(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 15
  expectedResultStr := "7726160"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_22(t *testing.T) {
  numOfItemsInt := 7
  numOfItemsChosenInt := 3
  expectedResultStr := "84"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_23(t *testing.T) {
  numOfItemsInt := 3
  numOfItemsChosenInt := 7
  expectedResultStr := "36"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_24(t *testing.T) {
  numOfItemsInt := 62
  numOfItemsChosenInt := 5
  expectedResultStr := "8936928"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_25(t *testing.T) {
  numOfItemsInt := 97
  numOfItemsChosenInt := 5
  expectedResultStr := "79208745"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_26(t *testing.T) {
  numOfItemsInt := 15
  numOfItemsChosenInt := 15
  expectedResultStr := "77558760"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_27(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 1
  expectedResultStr := "12"
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  result, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumStrDto_28(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 15
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  _, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumStrDto_29(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 0
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  _, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumStrDto_30(t *testing.T) {
  numOfItemsInt := -12
  numOfItemsChosenInt := 6
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  _, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumStrDto_31(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := -6
  allowRepetitions := true

  numOfItems := NumStrDto{}.NewInt(numOfItemsInt, 0)
  numOfItemsChosen := NumStrDto{}.NewInt(numOfItemsChosenInt, 0)

  _, err := Probability{}.CombinationsNumStrDto(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_01(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 3
  expectedResultStr := "560"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumStrDto("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_02(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 12
  expectedResultStr := "1820"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_03(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 5
  expectedResultStr := "2598960"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_04(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 26
  expectedResultStr := "495918532948104"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_05(t *testing.T) {

  numOfItemsInt := 18
  numOfItemsChosenInt := 7
  expectedResultStr := "31824"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_06(t *testing.T) {

  numOfItemsInt := 22
  numOfItemsChosenInt := 5
  expectedResultStr := "26334"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_07(t *testing.T) {

  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_08(t *testing.T) {
  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_09(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 25
  expectedResultStr := "1"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_10(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 1
  expectedResultStr := "25"
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_11(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 52
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_12(t *testing.T) {
  numOfItemsInt := 52
  numOfItemsChosenInt := 0
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_13(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 26
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_14(t *testing.T) {
  numOfItemsInt := -52
  numOfItemsChosenInt := 26
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_15(t *testing.T) {
  numOfItemsInt := 52
  numOfItemsChosenInt := -26
  allowRepetitions := false

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_16(t *testing.T) {
  numOfItemsInt := 5
  numOfItemsChosenInt := 3
  expectedResultStr := "35"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_17(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 11
  expectedResultStr := "705432"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_18(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 2
  expectedResultStr := "351"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_19(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 24
  expectedResultStr := "63205303218876"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_20(t *testing.T) {
  numOfItemsInt := 10
  numOfItemsChosenInt := 14
  expectedResultStr := "817190"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_21(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 15
  expectedResultStr := "7726160"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_22(t *testing.T) {
  numOfItemsInt := 7
  numOfItemsChosenInt := 3
  expectedResultStr := "84"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_23(t *testing.T) {
  numOfItemsInt := 3
  numOfItemsChosenInt := 7
  expectedResultStr := "36"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_24(t *testing.T) {
  numOfItemsInt := 62
  numOfItemsChosenInt := 5
  expectedResultStr := "8936928"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_25(t *testing.T) {
  numOfItemsInt := 97
  numOfItemsChosenInt := 5
  expectedResultStr := "79208745"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_26(t *testing.T) {
  numOfItemsInt := 15
  numOfItemsChosenInt := 15
  expectedResultStr := "77558760"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_27(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 1
  expectedResultStr := "12"
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  result, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsNumberStr("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsNumberStr_28(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 15
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_29(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 0
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_30(t *testing.T) {
  numOfItemsInt := -12
  numOfItemsChosenInt := 6
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsNumberStr_31(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := -6
  allowRepetitions := true

  numOfItems := fmt.Sprintf("%v", numOfItemsInt)
  numOfItemsChosen := fmt.Sprintf("%v", numOfItemsChosenInt)

  _, err := Probability{}.CombinationsNumberStr(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}
