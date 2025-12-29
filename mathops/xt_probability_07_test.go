package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_CombinationsUint64_01(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_01"

  numOfItemsUint64 := uint64(16)

  numOfItemsChosenUint64 := uint64(3)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_02(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_02"

  numOfItemsUint64 := uint64(16)

  numOfItemsChosenUint64 := uint64(12)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_03(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_03"

  numOfItemsUint64 := uint64(52)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_04(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_04"

  numOfItemsUint64 := uint64(52)

  numOfItemsChosenUint64 := uint64(26)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_05(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_05"

  numOfItemsUint64 := uint64(18)

  numOfItemsChosenUint64 := uint64(7)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_06(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_06"

  numOfItemsUint64 := uint64(22)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_07(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_07"

  numOfItemsUint64 := uint64(56)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_08(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_08"

  numOfItemsUint64 := uint64(56)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_09(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_09"

  numOfItemsUint64 := uint64(25)

  numOfItemsChosenUint64 := uint64(25)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_10(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_10"

  numOfItemsUint64 := uint64(25)

  numOfItemsChosenUint64 := uint64(1)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_11(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_11"

  numOfItemsUint64 := uint64(26)

  numOfItemsChosenUint64 := uint64(52)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint64 < numOfItemsChosenUint64\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint64_12(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_12"

  numOfItemsUint64 := uint64(52)

  numOfItemsChosenUint64 := uint64(0)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenUint64 <= 0\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint64_13(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_13"

  numOfItemsUint64 := uint64(0)

  numOfItemsChosenUint64 := uint64(26)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint64 <= 0\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint64_14(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_14"

  numOfItemsUint64 := uint64(5)

  numOfItemsChosenUint64 := uint64(3)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_15(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_15"

  numOfItemsUint64 := uint64(12)

  numOfItemsChosenUint64 := uint64(11)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_16(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_16"

  numOfItemsUint64 := uint64(26)

  numOfItemsChosenUint64 := uint64(2)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_17(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_17"

  numOfItemsUint64 := uint64(26)

  numOfItemsChosenUint64 := uint64(24)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_18(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_18"

  numOfItemsUint64 := uint64(10)

  numOfItemsChosenUint64 := uint64(14)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_19(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_19"

  numOfItemsUint64 := uint64(12)

  numOfItemsChosenUint64 := uint64(15)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_20(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_20"

  numOfItemsUint64 := uint64(7)

  numOfItemsChosenUint64 := uint64(3)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_21(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_21"

  numOfItemsUint64 := uint64(3)

  numOfItemsChosenUint64 := uint64(7)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_22(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_22"

  numOfItemsUint64 := uint64(62)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_23(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_23"

  numOfItemsUint64 := uint64(97)

  numOfItemsChosenUint64 := uint64(5)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_24(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_24"

  numOfItemsUint64 := uint64(15)

  numOfItemsChosenUint64 := uint64(15)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_25(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_25"

  numOfItemsUint64 := uint64(12)

  numOfItemsChosenUint64 := uint64(1)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

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

  bigIntNumResult, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
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

func TestProbability_CombinationsUint64_26(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_26"

  numOfItemsUint64 := uint64(0)

  numOfItemsChosenUint64 := uint64(15)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsUint64 <= 0\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsUint64_27(t *testing.T) {

  ePrefix := "TestProbability_CombinationsUint64_27"

  numOfItemsUint64 := uint64(12)

  numOfItemsChosenUint64 := uint64(0)

  numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

  numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsUint64(\n"+
      "  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
      "numOfItemsUint64= '%v'\n"+
      "numOfItemsChosenUint64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenUint64 <= 0\n\n",
      ePrefix,
      numOfItemsUint64Str,
      numOfItemsChosenUint64Str,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}
