package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_PermutationsInt_01(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_01"

  numOfItemsInt := 3

  numOfItemsChosenInt := 2

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_02(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_02"

  numOfItemsInt := 3

  numOfItemsChosenInt := 2

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_03(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_03"

  numOfItemsInt := 10

  numOfItemsChosenInt := 3

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_04(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_04"

  numOfItemsInt := 20

  numOfItemsChosenInt := 5

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_05(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_05"

  numOfItemsInt := 52

  numOfItemsChosenInt := 5

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_06(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_06"

  numOfItemsInt := 5

  numOfItemsChosenInt := 3

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_07(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_07"

  numOfItemsInt := 20

  numOfItemsChosenInt := 5

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_08(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_08"

  numOfItemsInt := 5

  numOfItemsChosenInt := 11

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_09(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_09"

  numOfItemsInt := 56

  numOfItemsChosenInt := 5

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_10(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_10"

  numOfItemsInt := 9

  numOfItemsChosenInt := 3

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_11(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_11"

  numOfItemsInt := 12

  numOfItemsChosenInt := 7

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_12(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_12"

  numOfItemsInt := 18

  numOfItemsChosenInt := 8

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_13(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_13"

  numOfItemsInt := 9

  numOfItemsChosenInt := 9

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_14(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_14"

  numOfItemsInt := 9

  numOfItemsChosenInt := 9

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_15(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_15"

  numOfItemsInt := 9

  numOfItemsChosenInt := 1

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_16(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_16"

  numOfItemsInt := 9

  numOfItemsChosenInt := 1

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

  bigINumResult, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigINumResult, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
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

func TestProbability_PermutationsInt_17(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_17"

  numOfItemsInt := 0

  numOfItemsChosenInt := 4

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_18(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_18"

  numOfItemsInt := 15

  numOfItemsChosenInt := 0

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsChosenInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_19(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_19"

  numOfItemsInt := -15

  numOfItemsChosenInt := 2

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_20(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_20"

  numOfItemsInt := 15

  numOfItemsChosenInt := -2

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsChosenInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_21(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_21"

  numOfItemsInt := 5

  numOfItemsChosenInt := 11

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsInt < numOfItemsChosenInt\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_22(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_22"

  numOfItemsInt := 0

  numOfItemsChosenInt := 4

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsInt < numOfItemsChosenInt\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_23(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_23"

  numOfItemsInt := 15

  numOfItemsChosenInt := 0

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsChosenInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_24(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_24"

  numOfItemsInt := -15

  numOfItemsChosenInt := 2

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt_25(t *testing.T) {

  ePrefix := "TestProbability_PermutationsInt_25"

  numOfItemsInt := 15

  numOfItemsChosenInt := -2

  numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

  numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).PermutationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "_, err := new(Probability).PermutationsInt(\n"+
      "  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
      "numOfItemsInt= '%v'\n"+
      "numOfItemsChosenInt= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because:\n"+
      "  numOfItemsChosenInt <= 0\n\n",
      ePrefix,
      numOfItemsIntStr,
      numOfItemsChosenIntStr,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_PermutationsInt32_01(t *testing.T) {

  numOfItems := int32(3)
  numOfItemsPicked := int32(2)
  allowRepetitions := false
  expectedResultStr := "6"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_02(t *testing.T) {

  numOfItems := int32(3)
  numOfItemsPicked := int32(2)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_03(t *testing.T) {

  numOfItems := int32(10)
  numOfItemsPicked := int32(3)
  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_04(t *testing.T) {

  numOfItems := int32(20)
  numOfItemsPicked := int32(5)
  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_05(t *testing.T) {

  numOfItems := int32(52)
  numOfItemsPicked := int32(5)
  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_06(t *testing.T) {

  numOfItems := int32(5)
  numOfItemsPicked := int32(3)
  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_07(t *testing.T) {

  numOfItems := int32(20)
  numOfItemsPicked := int32(5)
  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_08(t *testing.T) {

  numOfItems := int32(5)
  numOfItemsPicked := int32(11)
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_09(t *testing.T) {

  numOfItems := int32(56)
  numOfItemsPicked := int32(5)
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_10(t *testing.T) {

  numOfItems := int32(9)
  numOfItemsPicked := int32(3)
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_11(t *testing.T) {

  numOfItems := int32(12)
  numOfItemsPicked := int32(7)
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_12(t *testing.T) {

  numOfItems := int32(18)
  numOfItemsPicked := int32(8)
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_13(t *testing.T) {

  numOfItems := int32(9)
  numOfItemsPicked := int32(9)
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_14(t *testing.T) {

  numOfItems := int32(9)
  numOfItemsPicked := int32(9)
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_15(t *testing.T) {

  numOfItems := int32(9)
  numOfItemsPicked := int32(1)
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_16(t *testing.T) {

  numOfItems := int32(9)
  numOfItemsPicked := int32(1)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt32_17(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_18(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_19(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_20(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_21(t *testing.T) {
  nInt := 5
  rInt := 11
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_22(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_23(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_24(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt32_25(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int32(nInt)
  numOfItemsPicked := int32(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt32(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt32("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_01(t *testing.T) {

  numOfItems := int64(3)
  numOfItemsPicked := int64(2)
  allowRepetitions := false
  expectedResultStr := "6"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_02(t *testing.T) {

  numOfItems := int64(3)
  numOfItemsPicked := int64(2)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_03(t *testing.T) {

  numOfItems := int64(10)
  numOfItemsPicked := int64(3)
  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_04(t *testing.T) {

  numOfItems := int64(20)
  numOfItemsPicked := int64(5)
  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_05(t *testing.T) {

  numOfItems := int64(52)
  numOfItemsPicked := int64(5)
  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_06(t *testing.T) {

  numOfItems := int64(5)
  numOfItemsPicked := int64(3)
  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_07(t *testing.T) {

  numOfItems := int64(20)
  numOfItemsPicked := int64(5)
  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_08(t *testing.T) {

  numOfItems := int64(5)
  numOfItemsPicked := int64(11)
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_09(t *testing.T) {

  numOfItems := int64(56)
  numOfItemsPicked := int64(5)
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_10(t *testing.T) {

  numOfItems := int64(9)
  numOfItemsPicked := int64(3)
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_11(t *testing.T) {

  numOfItems := int64(12)
  numOfItemsPicked := int64(7)
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_12(t *testing.T) {

  numOfItems := int64(18)
  numOfItemsPicked := int64(8)
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_13(t *testing.T) {

  numOfItems := int64(9)
  numOfItemsPicked := int64(9)
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_14(t *testing.T) {

  numOfItems := int64(9)
  numOfItemsPicked := int64(9)
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_15(t *testing.T) {

  numOfItems := int64(9)
  numOfItemsPicked := int64(1)
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_16(t *testing.T) {

  numOfItems := int64(9)
  numOfItemsPicked := int64(1)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsInt64_17(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_18(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_19(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_20(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_21(t *testing.T) {
  nInt := 5
  rInt := 11
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }
}

func TestProbability_PermutationsInt64_22(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }
}

func TestProbability_PermutationsInt64_23(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_24(t *testing.T) {
  nInt := -15
  rInt := 2
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsInt64_25(t *testing.T) {
  nInt := 15
  rInt := -2
  numOfItems := int64(nInt)
  numOfItemsPicked := int64(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsInt64(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsInt64("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r < 0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_01(t *testing.T) {

  numOfItems := uint(3)
  numOfItemsPicked := uint(2)
  allowRepetitions := false
  expectedResultStr := "6"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_02(t *testing.T) {

  numOfItems := uint(3)
  numOfItemsPicked := uint(2)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_03(t *testing.T) {

  numOfItems := uint(10)
  numOfItemsPicked := uint(3)
  allowRepetitions := true
  expectedResultStr := "1000"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_04(t *testing.T) {

  numOfItems := uint(20)
  numOfItemsPicked := uint(5)
  allowRepetitions := false
  expectedResultStr := "1860480"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_05(t *testing.T) {

  numOfItems := uint(52)
  numOfItemsPicked := uint(5)
  allowRepetitions := false
  expectedResultStr := "311875200"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_06(t *testing.T) {

  numOfItems := uint(5)
  numOfItemsPicked := uint(3)
  allowRepetitions := true
  expectedResultStr := "125"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_07(t *testing.T) {

  numOfItems := uint(20)
  numOfItemsPicked := uint(5)
  allowRepetitions := true
  expectedResultStr := "3200000"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_08(t *testing.T) {

  numOfItems := uint(5)
  numOfItemsPicked := uint(11)
  allowRepetitions := true
  expectedResultStr := "48828125"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_09(t *testing.T) {

  numOfItems := uint(11)
  numOfItemsPicked := uint(5)
  allowRepetitions := false
  expectedResultStr := "55440"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_10(t *testing.T) {

  numOfItems := uint(11)
  numOfItemsPicked := uint(5)
  allowRepetitions := true
  expectedResultStr := "161051"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_11(t *testing.T) {

  numOfItems := uint(56)
  numOfItemsPicked := uint(5)
  allowRepetitions := false
  expectedResultStr := "458377920"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_12(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(3)
  allowRepetitions := false
  expectedResultStr := "504"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_13(t *testing.T) {

  numOfItems := uint(12)
  numOfItemsPicked := uint(7)
  allowRepetitions := false
  expectedResultStr := "3991680"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_14(t *testing.T) {

  numOfItems := uint(18)
  numOfItemsPicked := uint(8)
  allowRepetitions := false
  expectedResultStr := "1764322560"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_15(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(9)
  allowRepetitions := false
  expectedResultStr := "362880"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_16(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(9)
  allowRepetitions := true
  expectedResultStr := "387420489"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_17(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(1)
  allowRepetitions := false
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_18(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(1)
  allowRepetitions := true
  expectedResultStr := "9"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_19(t *testing.T) {

  numOfItems := uint(9)
  numOfItemsPicked := uint(4)
  allowRepetitions := true
  expectedResultStr := "6561"

  result, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err != nil {
    t.Errorf("Error returned by Probability{}.PermutationsUint(numOfItems, numOfItemsPicked). "+
      "numOfItems='%v' numOfItemsPicked='%v' Error='%v'",
      numOfItems, numOfItemsPicked, err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedResultStr != actualNumStr {
    t.Errorf("Error: Expected pemutations='%v'. Instead, permutations='%v'. ",
      expectedResultStr, actualNumStr)
  }

}

func TestProbability_PermutationsUint_20(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_21(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := true

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_22(t *testing.T) {
  nInt := 5
  rInt := 11
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsBigIntNum(n, r) "+
      "However no error was generated. r > n;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_23(t *testing.T) {
  nInt := 0
  rInt := 4
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_24(t *testing.T) {
  nInt := 15
  rInt := 0
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from PermutationsUint("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. r==0;  n='%v' r='%v' ", nInt, rInt)
  }

}

func TestProbability_PermutationsUint_25(t *testing.T) {
  nInt := 0
  rInt := 0
  numOfItems := uint(nInt)
  numOfItemsPicked := uint(rInt)
  allowRepetitions := false

  _, err := Probability{}.PermutationsUint(numOfItems, numOfItemsPicked, allowRepetitions)

  if err == nil {
    t.Errorf("Error: Expected error return from Probability{}.PermutationsUint("+
      "numOfItems, numOfItemsPicked, allowRepetitions) "+
      "However no error was generated. n==0;  n='%v' r='%v' ", nInt, rInt)
  }

}
