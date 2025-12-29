package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_PermutationsIntAry_01(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_01"

  numOfItemsInt := 3
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "6"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_02(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_02"

  numOfItemsInt := 3
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "9"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_03(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_03"

  numOfItemsInt := 10
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "1000"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_04(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_04"

  numOfItemsInt := 20
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1860480"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_05(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_05"

  numOfItemsInt := 52
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "311875200"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_06(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_06"

  numOfItemsInt := 5
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "125"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_07(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_07"

  numOfItemsInt := 20
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "3200000"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_08(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_08"

  numOfItemsInt := 5
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 11
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "48828125"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_09(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_09"

  numOfItemsInt := 56
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 5
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "458377920"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_10(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_10"

  numOfItemsInt := 9
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 3
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "504"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_11(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_11"

  numOfItemsInt := 12
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 7
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "3991680"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_12(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_12"

  numOfItemsInt := 18
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 8
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "1764322560"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_13(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_13"

  numOfItemsInt := 9
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 9
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "362880"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_14(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_14"

  numOfItemsInt := 9
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 9
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "387420489"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_15(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_15"

  numOfItemsInt := 9
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 1
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "9"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_16(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_16"

  numOfItemsInt := 9
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 1
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  expectedNumStr := "9"

  expectedMagnitudeInt := len(expectedNumStr) - 1

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

  intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItems, err := new(IntAry).\n"+
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

  err = intAryNumOfItems.IsValid("Validating intAryNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItems.IsValid('Validating intAryNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsNumStr, err := intAryNumOfItems.GetNumStr()\n"+
      "intAryNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != intAryNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != intAryNumOfItemsNumStr\n"+
      "Expected intAryNumOfItemsNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, intAryNumOfItemsNumStr)

    return
  }

  intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosen, err := new(IntAry).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "numOfItemsChosenPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenInt,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = intAryNumOfItemsChosen.IsValid("Validating intAryNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryNumOfItemsChosen.IsValid('Validating intAryNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumOfItemsChosenNumStr, err := intAryNumOfItemsChosen.GetNumStr()\n"+
      "intAryNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: intAryNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != intAryNumOfItemsChosenNumStr\n"+
      "Expected intAryNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual intAryNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, intAryNumOfItemsChosenNumStr)

    return
  }

  intAryResult, err := new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResult, err := new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = intAryResult.IsValid("Validating intAryResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := intAryResult.IsValid('Validating intAryResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryResultNumStr, err := intAryResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumStr, err := intAryResult.GetNumStr()\n"+
      "intAryResult set to final value\n"+
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
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultScaleFactorBigInt, err := intAryResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultScaleFactorBigInt, err :=\n"+
      "  intAryResult.GetScaleVal()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultMagnitudeInt, err := intAryResult.GetMagnitude()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultMagnitudeInt, err :=\n"+
      "  intAryResult.GetMagnitude()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultSignValue, err := intAryResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultSignValue, err := intAryResult.GetSign()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
      "intAryResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryResultNumStr, err.Error())
    return
  }

  intAryResultBigInt, err := intAryResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryResultBigInt, err :=\n"+
      "  intAryResult.GetBigInt()\n"+
      "intAryResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryResultNumStr, err.Error())
    return
  }

  if expectedNumStr != intAryResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != intAryResultNumStr\n"+
      "Expected intAryResultNumStr = '%v'\n"+
      "  Actual intAryResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, intAryResultNumStr)

    return
  }

  if expectedPrecisionInt != intAryResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
      "Expected intAryResultPrecisionInt = '%v'\n"+
      "  Actual intAryResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != intAryResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
      "Expected intAryResultPrecisionUint = '%v'\n"+
      "  Actual intAryResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

    return
  }

  if expectedSignValue != intAryResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != intAryResultSignValue\n"+
      "Expected intAryResultSignValue = '%v'\n"+
      "  Actual intAryResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != intAryResultNumSeps \n"+
      "Expected intAryResultNumSeps = '%v'\n"+
      "  Actual intAryResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(intAryResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=intAryResultScaleFactorBigInt\n"+
      "Expected intAryResultScaleFactorBigInt = '%v'\n"+
      "  Actual intAryResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      intAryResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedMagnitudeInt != intAryResultMagnitudeInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
      "Because expectedMagnitudeInt != intAryResultMagnitudeInt\n"+
      "Expected intAryResultMagnitudeInt = '%v'\n"+
      "  Actual intAryResultMagnitudeInt = '%v'\n\n",
      ePrefix, expectedMagnitudeInt, intAryResultMagnitudeInt)

    return
  }

  if expectedBigInt.Cmp(intAryResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual intAryResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(intAryResultBigInt) != 0\n"+
      "Expected intAryResultBigInt = '%v'\n"+
      "  Actual intAryResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), intAryResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_17(t *testing.T) {
  nIntAry := 0
  rIntAry := 4

  numOfItems := IntAry{}.NewInt(nIntAry, 0)

  numOfItemsPicked := IntAry{}.NewInt(rIntAry, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nIntAry, rIntAry)
  }

}

func TestProbability_PermutationsIntAry_18(t *testing.T) {
  nInt := 15
  rInt := 0

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_19(t *testing.T) {
  nInt := -15
  rInt := 2

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_20(t *testing.T) {
  nInt := 15
  rInt := -2

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_21(t *testing.T) {
  nInt := 5
  rInt := 11

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_22(t *testing.T) {
  nIntAry := 0
  rIntAry := 4

  numOfItems := IntAry{}.NewInt(nIntAry, 0)

  numOfItemsPicked := IntAry{}.NewInt(rIntAry, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nIntAry, rIntAry)
  }

}

func TestProbability_PermutationsIntAry_23(t *testing.T) {
  nInt := 15
  rInt := 0

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_24(t *testing.T) {
  nInt := -15
  rInt := 2

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsIntAry_25(t *testing.T) {
  nInt := 15
  rInt := -2

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := false

  _, err := Probability{}.PermutationsIntAry(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsIntAry(n, r) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }
}

func TestProbability_PermutationsINumMgr_01(t *testing.T) {

  numOfItems := Decimal{}.NewInt(3, 0)

  numOfItemsPicked := IntAry{}.NewInt(2, 0)

  allowRepetitions := false

  expectedResultStr := "6"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_02(t *testing.T) {

  numOfItems := IntAry{}.NewInt(3, 0)

  numOfItemsPicked := NumStrDto{}.NewInt(2, 0)

  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsINumMgr_03(t *testing.T) {

  numOfItems := BigIntNum{}.NewIntExponent(10, 0)

  numOfItemsPicked := IntAry{}.NewInt(3, 0)

  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_04(t *testing.T) {

  numOfItems := Decimal{}.NewInt(20, 0)

  numOfItemsPicked := NumStrDto{}.NewInt(5, 0)

  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsINumMgr_05(t *testing.T) {

  numOfItems := IntAry{}.NewInt(52, 0)

  numOfItemsPicked := IntAry{}.NewInt(5, 0)

  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsINumMgr_06(t *testing.T) {

  numOfItems := BigIntNum{}.NewInt(5, 0)

  numOfItemsPicked := Decimal{}.NewInt(3, 0)

  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_07(t *testing.T) {

  numOfItems := Decimal{}.NewInt(20, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(5, 0)

  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_08(t *testing.T) {

  numOfItems := Decimal{}.NewInt(5, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(11, 0)

  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_09(t *testing.T) {

  numOfItems := Decimal{}.NewInt(56, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(5, 0)

  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_10(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(3, 0)

  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_11(t *testing.T) {

  numOfItems := Decimal{}.NewInt(12, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(7, 0)

  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_12(t *testing.T) {

  numOfItems := Decimal{}.NewInt(18, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(8, 0)

  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_13(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(9, 0)

  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_14(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(9, 0)

  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_15(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(1, 0)

  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_16(t *testing.T) {

  numOfItems := Decimal{}.NewInt(9, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(1, 0)

  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsINumMgr(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsINumMgr_17(t *testing.T) {
  nINumMgr := 0
  rINumMgr := 4

  numOfItems := IntAry{}.NewInt(nINumMgr, 0)

  numOfItemsPicked := NumStrDto{}.NewInt(rINumMgr, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nINumMgr, rINumMgr)
  }

}

func TestProbability_PermutationsINumMgr_18(t *testing.T) {
  nInt := 15
  rInt := 0

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(rInt, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_19(t *testing.T) {
  nInt := -15
  rInt := 2

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  allowRepetitions := true

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_20(t *testing.T) {
  nInt := 15
  rInt := -2
  allowRepetitions := true

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_21(t *testing.T) {
  nInt := 5
  rInt := 11
  allowRepetitions := false

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr("+
      "&numOfItems, &numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_22(t *testing.T) {
  nINumMgr := 0
  rINumMgr := 4
  allowRepetitions := false

  numOfItems := IntAry{}.NewInt(nINumMgr, 0)

  numOfItemsPicked := NumStrDto{}.NewInt(rINumMgr, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nINumMgr, rINumMgr)
  }

}

func TestProbability_PermutationsINumMgr_23(t *testing.T) {
  nInt := 15
  rInt := 0
  allowRepetitions := false

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := BigIntNum{}.NewInt(rInt, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_24(t *testing.T) {
  nInt := -15
  rInt := 2
  allowRepetitions := false

  numOfItems := IntAry{}.NewInt(nInt, 0)

  numOfItemsPicked := IntAry{}.NewInt(rInt, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsINumMgr_25(t *testing.T) {
  nInt := 15
  rInt := -2
  allowRepetitions := false

  numOfItems := Decimal{}.NewInt(nInt, 0)

  numOfItemsPicked := Decimal{}.NewInt(rInt, 0)

  _, err := Probability{}.PermutationsINumMgr(&numOfItems, &numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsINumMgr(n, r) "+
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
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
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt_08(t *testing.T) {

  numOfItems := int(5)
  numOfItemsPicked := int(11)
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_09(t *testing.T) {

  numOfItems := int(56)
  numOfItemsPicked := int(5)
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_10(t *testing.T) {

  numOfItems := int(9)
  numOfItemsPicked := int(3)
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_11(t *testing.T) {

  numOfItems := int(12)
  numOfItemsPicked := int(7)
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_12(t *testing.T) {

  numOfItems := int(18)
  numOfItemsPicked := int(8)
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_13(t *testing.T) {

  numOfItems := int(9)
  numOfItemsPicked := int(9)
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_14(t *testing.T) {

  numOfItems := int(9)
  numOfItemsPicked := int(9)
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_15(t *testing.T) {

  numOfItems := int(9)
  numOfItemsPicked := int(1)
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_16(t *testing.T) {

  numOfItems := int(9)
  numOfItemsPicked := int(1)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected permutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }
}

func TestProbability_PermutationsInt_17(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_18(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_19(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_20(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_21(t *testing.T) {
  nInt := 5
  rInt := 11
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_22(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_23(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_24(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt_25(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int(nInt)
  numOfItemsPicked := int(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsWithRepsBigInt(n, r) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}
