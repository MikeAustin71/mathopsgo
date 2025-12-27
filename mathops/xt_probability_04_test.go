package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestProbability_CombinationsInt32_01(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_01"

  numOfItemsInt32 := int32(16)

  numOfItemsChosenInt32 := int32(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_02(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_02"

  numOfItemsInt32 := int32(16)

  numOfItemsChosenInt32 := int32(12)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_03(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_03"

  numOfItemsInt32 := int32(52)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_04(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_04"

  numOfItemsInt32 := int32(52)

  numOfItemsChosenInt32 := int32(26)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_05(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_05"

  numOfItemsInt32 := int32(18)

  numOfItemsChosenInt32 := int32(7)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_06(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_06"

  numOfItemsInt32 := int32(22)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_07(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_07"

  numOfItemsInt32 := int32(56)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_08(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_08"

  numOfItemsInt32 := int32(56)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_09(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_09"

  numOfItemsInt32 := int32(25)

  numOfItemsChosenInt32 := int32(25)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_10(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_10"

  numOfItemsInt32 := int32(25)

  numOfItemsChosenInt32 := int32(1)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_11(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_11"

  numOfItemsInt32 := int32(26)

  numOfItemsChosenInt32 := int32(52)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt32 < numOfItemsChosenInt32\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_12(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_12"

  numOfItemsInt32 := int32(52)

  numOfItemsChosenInt32 := int32(0)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_13(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_13"

  numOfItemsInt32 := int32(0)

  numOfItemsChosenInt32 := int32(26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_14(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_14"

  numOfItemsInt32 := int32(-52)

  numOfItemsChosenInt32 := int32(26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_15(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_15"

  numOfItemsInt32 := int32(52)

  numOfItemsChosenInt32 := int32(-26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_16(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_16"

  numOfItemsInt32 := int32(5)

  numOfItemsChosenInt32 := int32(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_17(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_17"

  numOfItemsInt32 := int32(12)

  numOfItemsChosenInt32 := int32(11)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_18(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_18"

  numOfItemsInt32 := int32(26)

  numOfItemsChosenInt32 := int32(2)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_19(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_19"

  numOfItemsInt32 := int32(26)

  numOfItemsChosenInt32 := int32(24)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_20(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_20"

  numOfItemsInt32 := int32(10)

  numOfItemsChosenInt32 := int32(14)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_21(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_21"

  numOfItemsInt32 := int32(12)

  numOfItemsChosenInt32 := int32(15)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_22(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_22"

  numOfItemsInt32 := int32(7)

  numOfItemsChosenInt32 := int32(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_23(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_23"

  numOfItemsInt32 := int32(3)

  numOfItemsChosenInt32 := int32(7)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_24(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_24"

  numOfItemsInt32 := int32(62)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_25(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_25"

  numOfItemsInt32 := int32(97)

  numOfItemsChosenInt32 := int32(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_26(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_26"

  numOfItemsInt32 := int32(15)

  numOfItemsChosenInt32 := int32(15)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_27(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_27"

  numOfItemsInt32 := int32(12)

  numOfItemsChosenInt32 := int32(1)

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

  bigIntNumResult, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
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

func TestProbability_CombinationsInt32_28(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_28"

  numOfItemsInt32 := int32(0)

  numOfItemsChosenInt32 := int32(15)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_29(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_29"

  numOfItemsInt32 := int32(12)

  numOfItemsChosenInt32 := int32(0)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_30(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_30"

  numOfItemsInt32 := int32(-12)

  numOfItemsChosenInt32 := int32(6)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt32_31(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt32_31"

  numOfItemsInt32 := int32(12)

  numOfItemsChosenInt32 := int32(-6)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt32(numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt32(\n"+
      "  numOfItemsInt32, numOfItemsChosenInt32, allowRepetitions)\n"+
      "numOfItemsInt32= '%v'\n"+
      "numOfItemsChosenInt32= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt32 <= 0\n\n",
      ePrefix,
      numOfItemsInt32,
      numOfItemsChosenInt32,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_01(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_01"

  numOfItemsInt64 := int64(16)

  numOfItemsChosenInt64 := int64(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_02(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_02"

  numOfItemsInt64 := int64(16)

  numOfItemsChosenInt64 := int64(12)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_03(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_03"

  numOfItemsInt64 := int64(52)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_04(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_04"

  numOfItemsInt64 := int64(52)

  numOfItemsChosenInt64 := int64(26)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_05(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_05"

  numOfItemsInt64 := int64(18)

  numOfItemsChosenInt64 := int64(7)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_06(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_06"

  numOfItemsInt64 := int64(22)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_07(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_07"

  numOfItemsInt64 := int64(56)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_08(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_08"

  numOfItemsInt64 := int64(56)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_09(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_09"

  numOfItemsInt64 := int64(25)

  numOfItemsChosenInt64 := int64(25)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_10(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_10"

  numOfItemsInt64 := int64(25)

  numOfItemsChosenInt64 := int64(1)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_11(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_11"

  numOfItemsInt64 := int64(26)

  numOfItemsChosenInt64 := int64(52)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt64 < numOfItemsInt64\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_12(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_12"

  numOfItemsInt64 := int64(52)

  numOfItemsChosenInt64 := int64(0)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_13(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_13"

  numOfItemsInt64 := int64(0)

  numOfItemsChosenInt64 := int64(26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt64 <= 0 \n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_14(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_14"

  numOfItemsInt64 := int64(-52)

  numOfItemsChosenInt64 := int64(26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_15(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_15"

  numOfItemsInt64 := int64(52)

  numOfItemsChosenInt64 := int64(-26)

  var allowRepetitions bool

  allowRepetitions = false

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_16(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_16"

  numOfItemsInt64 := int64(5)

  numOfItemsChosenInt64 := int64(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_17(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_17"

  numOfItemsInt64 := int64(12)

  numOfItemsChosenInt64 := int64(11)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_18(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_18"

  numOfItemsInt64 := int64(26)

  numOfItemsChosenInt64 := int64(2)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_19(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_19"

  numOfItemsInt64 := int64(26)

  numOfItemsChosenInt64 := int64(24)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_20(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_20"

  numOfItemsInt64 := int64(10)

  numOfItemsChosenInt64 := int64(14)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_21(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_21"

  numOfItemsInt64 := int64(12)

  numOfItemsChosenInt64 := int64(15)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_22(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_22"

  numOfItemsInt64 := int64(7)

  numOfItemsChosenInt64 := int64(3)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_23(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_23"

  numOfItemsInt64 := int64(3)

  numOfItemsChosenInt64 := int64(7)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_24(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_24"

  numOfItemsInt64 := int64(62)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_25(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_25"

  numOfItemsInt64 := int64(97)

  numOfItemsChosenInt64 := int64(5)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_26(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_26"

  numOfItemsInt64 := int64(15)

  numOfItemsChosenInt64 := int64(15)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_27(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_27"

  numOfItemsInt64 := int64(12)

  numOfItemsChosenInt64 := int64(1)

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

  bigIntNumResult, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigIntNumResult, err := new(Probability).CombinationsInt64(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
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

func TestProbability_CombinationsInt64_28(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_28"

  numOfItemsInt64 := int64(0)

  numOfItemsChosenInt64 := int64(15)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_29(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_29"

  numOfItemsInt64 := int64(12)

  numOfItemsChosenInt64 := int64(0)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_30(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_30"

  numOfItemsInt64 := int64(-12)

  numOfItemsChosenInt64 := int64(6)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}

func TestProbability_CombinationsInt64_31(t *testing.T) {

  ePrefix := "TestProbability_CombinationsInt64_31"

  numOfItemsInt64 := int64(12)

  numOfItemsChosenInt64 := int64(-6)

  var allowRepetitions bool

  allowRepetitions = true

  _, err := new(Probability).CombinationsInt64(numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "Function Call:\n"+
      "_, err = new(Probability).CombinationsInt(\n"+
      "  numOfItemsInt64, numOfItemsChosenInt64, allowRepetitions)\n"+
      "numOfItemsInt64= '%v'\n"+
      "numOfItemsChosenInt64= '%v'\n"+
      "allowRepetitions= '%v'\n"+
      "Error should have triggered because: numOfItemsChosenInt64 <= 0\n\n",
      ePrefix,
      numOfItemsInt64,
      numOfItemsChosenInt64,
      strconv.FormatBool(allowRepetitions))

    return
  }

  return
}
