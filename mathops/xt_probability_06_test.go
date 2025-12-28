package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_CombinationsUint_01(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_01"

  numOfItemsUint := uint(16)

  numOfItemsChosenUint := uint(3)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_02(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_01"

  numOfItemsUint := uint(16)

  numOfItemsChosenUint := uint(12)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_03(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_03"

  numOfItemsUint := uint(52)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_04(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_01"

  numOfItemsUint := uint(52)

  numOfItemsChosenUint := uint(26)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_05(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_05"

  numOfItemsUint := uint(18)

  numOfItemsChosenUint := uint(7)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_06(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_06"

  numOfItemsUint := uint(22)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_07(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_07"

  numOfItemsUint := uint(56)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_08(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_08"
  numOfItemsUint := uint(56)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_09(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_09"

  numOfItemsUint := uint(25)

  numOfItemsChosenUint := uint(25)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_10(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_10"

  numOfItemsUint := uint(25)

  numOfItemsChosenUint := uint(1)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_11(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_11"

  numOfItemsUint := uint(26)

  numOfItemsChosenUint := uint(52)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint < numOfItemsChosenUint\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint_12(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_11"

  numOfItemsUint := uint(52)

  numOfItemsChosenUint := uint(0)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenUint <= 0\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint_13(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_11"

  numOfItemsUint := uint(0)

  numOfItemsChosenUint := uint(26)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint <= 0\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint_14(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_14"

  numOfItemsUint := uint(5)

  numOfItemsChosenUint := uint(3)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "35"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_15(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_15"

  numOfItemsUint := uint(12)

  numOfItemsChosenUint := uint(11)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "705432"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_16(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_16"

  numOfItemsUint := uint(26)

  numOfItemsChosenUint := uint(2)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "351"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_17(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_17"

  numOfItemsUint := uint(26)

  numOfItemsChosenUint := uint(24)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "63205303218876"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_18(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_18"

  numOfItemsUint := uint(10)

  numOfItemsChosenUint := uint(14)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "817190"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_19(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_19"

  numOfItemsUint := uint(12)

  numOfItemsChosenUint := uint(15)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "7726160"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_20(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_20"

  numOfItemsUint := uint(7)

  numOfItemsChosenUint := uint(3)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "84"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_21(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_21"

  numOfItemsUint := uint(3)

  numOfItemsChosenUint := uint(7)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "36"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_22(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_22"

  numOfItemsUint := uint(62)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "8936928"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_23(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_23"

  numOfItemsUint := uint(97)

  numOfItemsChosenUint := uint(5)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "79208745"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_24(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_24"

  numOfItemsUint := uint(15)

  numOfItemsChosenUint := uint(15)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "77558760"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_25(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_25"

  numOfItemsUint := uint(12)

  numOfItemsChosenUint := uint(1)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = false

  expectedNumStr := "12"

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

  bigIntNumResult, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions),
      err.Error())

    return
  }

  err = bigIntNumResult.IsValid("Validating bigIntNumResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := bigIntNumResult.IsValid('Validating bigIntNumResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumStr, err := bigIntNumResult.GetNumStr()\n"+
      "bigIntNumResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionInt, err :=\n"+
      "  bigIntNumResult.GetPrecisionInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultPrecisionUint, err :=\n"+
      "  bigIntNumResult.GetPrecisionUint()\n"+
      "bigIntNumResultResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()\n"+
      "bigIntNumResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultScaleFactorBigInt, err := bigIntNumResult.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultScaleFactorBigInt, err :=\n"+
      "  bigIntNumResult.GetScaleFactor()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResultBigInt, err :=\n"+
      "  bigIntNumResult.GetBigInt()\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedNumStr != bigIntNumResultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Number Strings DON'T MATCH!\n"+
      "Because expectedNumStr != bigIntNumResultNumStr\n"+
      "Expected bigIntNumResultNumStr = '%v'\n"+
      "  Actual bigIntNumResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, bigIntNumResultNumStr)

    return
  }

  expectedAndResultBigINumsAreEqual, err := expectedBigINum.Equal(bigIntNumResult)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAndResultBigINumsAreEqual, err :=\n"+
      "  expectedBigINum.Equal(bigIntNumResult)\n"+
      "expectedBigINum= '%v'\n"+
      "bigIntNumResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr, err.Error())
    return
  }

  if expectedAndResultBigINumsAreEqual == false {
    t.Errorf("%v\n"+
      "Error: bigIntNumResult is INVALID!\n"+
      "Because expectedAndResultBigINumsAreEqual == false\n"+
      "Expected bigIntNumResult = '%v'\n"+
      "  Actual bigIntNumResult = '%v'\n\n",
      ePrefix, expectedBigINumStr, bigIntNumResultNumStr)

    return
  }

  if expectedPrecisionInt != bigIntNumResultPrecisionInt {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Integers DON'T MATCH!\n"+
      "Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
      "Expected bigIntNumResultPrecisionInt = '%v'\n"+
      "  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

    return
  }

  if expectedPrecisionUint != bigIntNumResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Precision Uint's DON'T MATCH!\n"+
      "Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
      "Expected bigIntNumResultPrecisionUint = '%v'\n"+
      "  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

    return
  }

  if expectedSignValue != bigIntNumResultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Sign Values DON'T MATCH!\n"+
      "Because expectedSignValue != bigIntNumResultSignValue\n"+
      "Expected bigIntNumResultSignValue = '%v'\n"+
      "  Actual bigIntNumResultSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, bigIntNumResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Numeric Separators DON'T MATCH!\n"+
      "Because expectedNumSeps != bigIntNumResultNumSeps \n"+
      "Expected bigIntNumResultNumSeps = '%v'\n"+
      "  Actual bigIntNumResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

    return
  }

  if expectedScaleFactorBigInt.Cmp(bigIntNumResultScaleFactorBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult Scale Factors DON'T MATCH!\n"+
      "Because expectedScaleFactorBigInt!=bigIntNumResultScaleFactorBigInt\n"+
      "Expected bigIntNumResultScaleFactorBigInt = '%v'\n"+
      "  Actual bigIntNumResultScaleFactorBigInt = '%v'\n\n",
      ePrefix,
      expectedScaleFactorBigInt.Text(10),
      bigIntNumResultScaleFactorBigInt.Text(10))

    return
  }

  if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Expected vs Actual bigIntNumResult BigInt Values DON'T MATCH!\n"+
      "Because expectedBigInt.Cmp(bigIntNumResultBigInt) != 0\n"+
      "Expected bigIntNumResultBigInt = '%v'\n"+
      "  Actual bigIntNumResultBigInt = '%v'\n\n",
      ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

    return
  }

  return
}

func TestProbability_CombinationsUint_26(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_26"

  numOfItemsUint := uint(0)

  numOfItemsChosenUint := uint(15)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint <= 0\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint_27(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint_27"

  numOfItemsUint := uint(12)

  numOfItemsChosenUint := uint(0)

  numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

  numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint(\n"+
      "  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
      "numOfItemsUint= '%v'\n"+
      "numOfItemsChosenUint= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenUint <= 0\n\n",
      ePrefix,
      numOfItemsUintStr,
      numOfItemsChosenUintStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint32_01(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 3
  expectedResultStr := "560"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_02(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 12
  expectedResultStr := "1820"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_03(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 5
  expectedResultStr := "2598960"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_04(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 26
  expectedResultStr := "495918532948104"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_05(t *testing.T) {

  numOfItemsInt := 18
  numOfItemsChosenInt := 7
  expectedResultStr := "31824"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_06(t *testing.T) {

  numOfItemsInt := 22
  numOfItemsChosenInt := 5
  expectedResultStr := "26334"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_07(t *testing.T) {

  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_08(t *testing.T) {
  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_09(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 25
  expectedResultStr := "1"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_10(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 1
  expectedResultStr := "25"
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_11(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 52
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint32_12(t *testing.T) {
  numOfItemsInt := 52
  numOfItemsChosenInt := 0
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint32_13(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 26
  allowRepetitions := false

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint32_14(t *testing.T) {
  numOfItemsInt := 5
  numOfItemsChosenInt := 3
  expectedResultStr := "35"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_15(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 11
  expectedResultStr := "705432"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_16(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 2
  expectedResultStr := "351"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_17(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 24
  expectedResultStr := "63205303218876"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_18(t *testing.T) {
  numOfItemsInt := 10
  numOfItemsChosenInt := 14
  expectedResultStr := "817190"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_19(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 15
  expectedResultStr := "7726160"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_20(t *testing.T) {
  numOfItemsInt := 7
  numOfItemsChosenInt := 3
  expectedResultStr := "84"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_21(t *testing.T) {
  numOfItemsInt := 3
  numOfItemsChosenInt := 7
  expectedResultStr := "36"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_22(t *testing.T) {
  numOfItemsInt := 62
  numOfItemsChosenInt := 5
  expectedResultStr := "8936928"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_23(t *testing.T) {
  numOfItemsInt := 97
  numOfItemsChosenInt := 5
  expectedResultStr := "79208745"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_24(t *testing.T) {
  numOfItemsInt := 15
  numOfItemsChosenInt := 15
  expectedResultStr := "77558760"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_25(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 1
  expectedResultStr := "12"
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint32("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint32_26(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 15
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint32_27(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 0
  allowRepetitions := true

  numOfItems := uint32(numOfItemsInt)
  numOfItemsChosen := uint32(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint32(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint64_01(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 3
  expectedResultStr := "560"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_02(t *testing.T) {

  numOfItemsInt := 16
  numOfItemsChosenInt := 12
  expectedResultStr := "1820"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_03(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 5
  expectedResultStr := "2598960"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_04(t *testing.T) {

  numOfItemsInt := 52
  numOfItemsChosenInt := 26
  expectedResultStr := "495918532948104"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_05(t *testing.T) {

  numOfItemsInt := 18
  numOfItemsChosenInt := 7
  expectedResultStr := "31824"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_06(t *testing.T) {

  numOfItemsInt := 22
  numOfItemsChosenInt := 5
  expectedResultStr := "26334"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_07(t *testing.T) {

  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_08(t *testing.T) {
  numOfItemsInt := 56
  numOfItemsChosenInt := 5
  expectedResultStr := "3819816"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_09(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 25
  expectedResultStr := "1"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_10(t *testing.T) {
  numOfItemsInt := 25
  numOfItemsChosenInt := 1
  expectedResultStr := "25"
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_11(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 52
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint64_12(t *testing.T) {
  numOfItemsInt := 52
  numOfItemsChosenInt := 0
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint64_13(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 26
  allowRepetitions := false

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
      "numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
      "allowRepetitions='%v'",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint64_14(t *testing.T) {
  numOfItemsInt := 5
  numOfItemsChosenInt := 3
  expectedResultStr := "35"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_15(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 11
  expectedResultStr := "705432"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_16(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 2
  expectedResultStr := "351"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_17(t *testing.T) {
  numOfItemsInt := 26
  numOfItemsChosenInt := 24
  expectedResultStr := "63205303218876"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_18(t *testing.T) {
  numOfItemsInt := 10
  numOfItemsChosenInt := 14
  expectedResultStr := "817190"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_19(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 15
  expectedResultStr := "7726160"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_20(t *testing.T) {
  numOfItemsInt := 7
  numOfItemsChosenInt := 3
  expectedResultStr := "84"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_21(t *testing.T) {
  numOfItemsInt := 3
  numOfItemsChosenInt := 7
  expectedResultStr := "36"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_22(t *testing.T) {
  numOfItemsInt := 62
  numOfItemsChosenInt := 5
  expectedResultStr := "8936928"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_23(t *testing.T) {
  numOfItemsInt := 97
  numOfItemsChosenInt := 5
  expectedResultStr := "79208745"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_24(t *testing.T) {
  numOfItemsInt := 15
  numOfItemsChosenInt := 15
  expectedResultStr := "77558760"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_25(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 1
  expectedResultStr := "12"
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  result, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.CombinationsUint64("+
      "numOfItems, numOfItemsChosen, allowRepetitions). "+
      "Error='%v' ", err.Error())
  }

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResultStr, actualResultStr)
  }
}

func TestProbability_CombinationsUint64_26(t *testing.T) {
  numOfItemsInt := 0
  numOfItemsChosenInt := 15
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}

func TestProbability_CombinationsUint64_27(t *testing.T) {
  numOfItemsInt := 12
  numOfItemsChosenInt := 0
  allowRepetitions := true

  numOfItems := uint64(numOfItemsInt)
  numOfItemsChosen := uint64(numOfItemsChosenInt)

  _, err := Probability{}.CombinationsUint64(numOfItems, numOfItemsChosen, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
      "numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
      numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
  }
}
