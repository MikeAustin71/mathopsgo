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

  ePrefix := "TestProbability_PermutationsIntAry_17"

  numOfItemsInt := 0
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 4
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItems <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_18(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_18"

  numOfItemsInt := 15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 0
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItemsChosen <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_19(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_19"

  numOfItemsInt := -15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItems <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_20(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_20"

  numOfItemsInt := 15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := -2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItemsChosen <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_21(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_21"

  numOfItemsInt := 5
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 11
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItems < intAryNumOfItemsChosen\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_22(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_22"

  numOfItemsInt := 0
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 4
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItems <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_23(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_23"

  numOfItemsInt := 15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 0
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItemsChosen <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_24(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_24"

  numOfItemsInt := -15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := 2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItems <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsIntAry_25(t *testing.T) {

  ePrefix := "TestProbability_PermutationsIntAry_25"

  numOfItemsInt := 15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := -2
  numOfItemsChosenPrecisionUint := uint(0)

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

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

  _, err = new(Probability).PermutationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err = new(Probability).PermutationsIntAry(\n"+
      "  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: intAryNumOfItemsChosen <= 0\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_01(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_01"

  numOfItemsInt := 15
  numOfItemsPrecisionUint := uint(0)

  numOfItemsChosenInt := -2
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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &numOfItems, &numOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_02(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_02"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

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
      numOfItemsInt,
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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&intAryNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &intAryNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "intAryNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_03(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_03"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItems, err := new(BigIntNum).\n"+
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

  err = bigINumNumOfItems.IsValid("Validating bigINumNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItems.IsValid('Validating bigINumNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()\n"+
      "bigINumNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != bigINumNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != bigINumNumOfItemsNumStr\n"+
      "Expected bigINumNumOfItemsNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, bigINumNumOfItemsNumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&bigINumNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &bigINumNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)\n"+
      "bigINumNumOfItems= '%v'\n"+
      "intAryNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigINumNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_04(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_04"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_05(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_05"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

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
      numOfItemsInt,
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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&intAryNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &intAryNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)\n"+
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

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_06(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_06"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItems, err := new(BigIntNum).\n"+
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

  err = bigINumNumOfItems.IsValid("Validating bigINumNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItems.IsValid('Validating bigINumNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()\n"+
      "bigINumNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != bigINumNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != bigINumNumOfItemsNumStr\n"+
      "Expected bigINumNumOfItemsNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, bigINumNumOfItemsNumStr)

    return
  }

  decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsChosen, err := new(Decimal).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = decimalNumOfItemsChosen.IsValid("Validating decimalNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItemsChosen.IsValid('Validating decimalNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsChosenNumStr, err := decimalNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsChosenNumStr, err := decimalNumOfItemsChosen.GetNumStr()\n"+
      "decimalNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != decimalNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != decimalNumOfItemsChosenNumStr\n"+
      "Expected decimalNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, decimalNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
      "bigINumNumOfItems= '%v'\n"+
      "decimalNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigINumNumOfItemsNumStr,
      decimalNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_07(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_07"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

    return
  }

  bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = bigINumNumOfItemsChosen.IsValid("Validating bigINumNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItemsChosen.IsValid('Validating bigINumNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()\n"+
      "bigINumNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr\n"+
      "Expected bigINumNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, bigINumNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "bigINumNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      bigINumNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_08(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_08"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

    return
  }

  bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = bigINumNumOfItemsChosen.IsValid("Validating bigINumNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItemsChosen.IsValid('Validating bigINumNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()\n"+
      "bigINumNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr\n"+
      "Expected bigINumNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, bigINumNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "bigINumNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      bigINumNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_09(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_09"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

    return
  }

  bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = bigINumNumOfItemsChosen.IsValid("Validating bigINumNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItemsChosen.IsValid('Validating bigINumNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()\n"+
      "bigINumNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr\n"+
      "Expected bigINumNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, bigINumNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "numOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      bigINumNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_10(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_10"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

    return
  }

  bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = bigINumNumOfItemsChosen.IsValid("Validating bigINumNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItemsChosen.IsValid('Validating bigINumNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()\n"+
      "bigINumNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr\n"+
      "Expected bigINumNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, bigINumNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "bigINumNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      bigINumNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_11(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_11"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItems, err := new(Decimal).\n"+
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

  err = decimalNumOfItems.IsValid("Validating decimalNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decimalNumOfItems.IsValid('Validating decimalNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decimalNumOfItemsNumStr, err := decimalNumOfItems.GetNumStr()\n"+
      "decimalNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != decimalNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: decimalNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != decimalNumOfItemsNumStr\n"+
      "Expected decimalNumOfItemsNumStr = '%v'\n"+
      "  Actual decimalNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, decimalNumOfItemsNumStr)

    return
  }

  bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
      "  NewInt(numOfItemsChosenInt, 0)\n"+
      "numOfItemsChosenInt = '%v'\n"+
      "numOfItemsChosenPrecisionUint = '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsChosenIntStr,
      numOfItemsChosenPrecisionUint,
      err.Error())

    return
  }

  err = bigINumNumOfItemsChosen.IsValid("Validating bigINumNumOfItemsChosen")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItemsChosen.IsValid('Validating bigINumNumOfItemsChosen')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsChosenNumStr, err := bigINumNumOfItemsChosen.GetNumStr()\n"+
      "bigINumNumOfItemsChosen set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItemsChosen Initialization FAILED!\n"+
      "Because numOfItemsChosenIntStr != bigINumNumOfItemsChosenNumStr\n"+
      "Expected bigINumNumOfItemsChosenNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsChosenNumStr = '%v'\n\n",
      ePrefix, numOfItemsChosenIntStr, bigINumNumOfItemsChosenNumStr)

    return
  }

  bigINumResult, err := new(Probability).PermutationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
      "decimalNumOfItems= '%v'\n"+
      "bigINumNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decimalNumOfItemsNumStr,
      bigINumNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_12(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_12"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "numStrDtoNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numStrDtoNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_13(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_13"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItems, err := new(BigIntNum).\n"+
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

  err = bigINumNumOfItems.IsValid("Validating bigINumNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItems.IsValid('Validating bigINumNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()\n"+
      "bigINumNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != bigINumNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != bigINumNumOfItemsNumStr\n"+
      "Expected bigINumNumOfItemsNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, bigINumNumOfItemsNumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&bigINumNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &bigINumNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)\n"+
      "bigINumNumOfItems= '%v'\n"+
      "numOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigINumNumOfItemsNumStr,
      intAryNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_PermutationsINumMgr_14(t *testing.T) {

  ePrefix := "TestProbability_PermutationsINumMgr_1X"

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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
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

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Expected Number String Values NOT Equal\n"+
      "Because expectedNumStr != expectedBigINumStr \n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedBigINumStr)

    return
  }

  bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItems, err := new(BigIntNum).\n"+
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

  err = bigINumNumOfItems.IsValid("Validating bigINumNumOfItems")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumNumOfItems.IsValid('Validating bigINumNumOfItems')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumNumOfItemsNumStr, err := bigINumNumOfItems.GetNumStr()\n"+
      "bigINumNumOfItems set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if numOfItemsIntStr != bigINumNumOfItemsNumStr {
    t.Errorf("%v\n"+
      "Error: bigINumNumOfItems Initialization FAILED!\n"+
      "Because numOfItemsIntStr != bigINumNumOfItemsNumStr\n"+
      "Expected bigINumNumOfItemsNumStr = '%v'\n"+
      "  Actual bigINumNumOfItemsNumStr = '%v'\n\n",
      ePrefix, numOfItemsIntStr, bigINumNumOfItemsNumStr)

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

  bigINumResult, err := new(Probability).PermutationsINumMgr(&bigINumNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsINumMgr(\n"+
      "  &bigINumNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
      "bigINumNumOfItems= '%v'\n"+
      "numStrDtoNumOfItemsChosen= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      bigINumNumOfItemsNumStr,
      numStrDtoNumOfItemsChosenNumStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigINumResult.IsValid("Validating bigINumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigINumResult.IsValid('Validating bigINumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultNumStr, err := bigINumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumStr, err := bigINumResult.GetNumStr()\n"+
      "bigINumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigINumResultPrecisionInt, err := bigINumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionInt, err :=\n"+
      "  bigINumResult.GetPrecisionInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultPrecisionUint, err := bigINumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultPrecisionUint, err :=\n"+
      "  bigINumResult.GetPrecisionUint()\n"+
      "bigINumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultSignValue, err := bigINumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultSignValue, err := bigINumResult.GetSign()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultNumSeps, err := bigINumResult.GetNumericSeparatorsDto()\n"+
      "bigINumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultScaleFactorBigInt, err := bigINumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultScaleFactorBigInt, err :=\n"+
      "  bigINumResult.GetScaleFactor()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  bigINumResultBigInt, err := bigINumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResultBigInt, err :=\n"+
      "  bigINumResult.GetBigInt()\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigINumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigINumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigINumResultNumStr\n"+
      "Expected bigINumResultNumStr = '%v'\n"+
      "  Actual bigINumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigINumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigINumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigINumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigINumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigINumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigINumResult = '%v'\n"+
      "  Actual bigINumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigINumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigINumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigINumResultPrecisionInt\n"+
      "Expected bigINumResultPrecisionInt = '%v'\n"+
      "  Actual bigINumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigINumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigINumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigINumResultPrecisionUint\n"+
      "Expected bigINumResultPrecisionUint = '%v'\n"+
      "  Actual bigINumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigINumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigINumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigINumResultSignValue\n"+
      "Expected bigINumResultSignValue = '%v'\n"+
      "  Actual bigINumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigINumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigINumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigINumResultNumSeps \n"+
      "Expected bigINumResultNumSeps = '%v'\n"+
      "  Actual bigINumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigINumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigINumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigINumResultScaleFactorBigInt\n"+
      "Expected bigINumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigINumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigINumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigINumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigINumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigINumResultBigInt) != 0\n"+
      "Expected bigINumResultBigInt = '%v'\n"+
      "  Actual bigINumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigINumResultBigInt.Text(10))

    return
  }

  return
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
