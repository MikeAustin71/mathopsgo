package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestProbability_PermutationsUint_01(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_01"

	numOfItemsUint := uint(3)

	numOfItemsChosenUint := uint(2)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_02(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_02"

	numOfItemsUint := uint(3)

	numOfItemsChosenUint := uint(2)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_03(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_03"

	numOfItemsUint := uint(10)

	numOfItemsChosenUint := uint(3)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_04(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_04"

	numOfItemsUint := uint(20)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_05(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_05"

	numOfItemsUint := uint(52)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_06(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_06"

	numOfItemsUint := uint(5)

	numOfItemsChosenUint := uint(3)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_07(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_07"

	numOfItemsUint := uint(20)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_08(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_08"

	numOfItemsUint := uint(5)

	numOfItemsChosenUint := uint(11)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_09(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_09"

	numOfItemsUint := uint(11)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_10(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_10"

	numOfItemsUint := uint(11)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_11(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_11"

	numOfItemsUint := uint(56)

	numOfItemsChosenUint := uint(5)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_12(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_12"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(3)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_13(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_13"

	numOfItemsUint := uint(12)

	numOfItemsChosenUint := uint(7)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_14(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_14"

	numOfItemsUint := uint(18)

	numOfItemsChosenUint := uint(8)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_15(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_15"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(9)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_16(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_16"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(9)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_17(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_17"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(1)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_18(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_18"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(1)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_19(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_19"

	numOfItemsUint := uint(9)

	numOfItemsChosenUint := uint(4)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "6561"

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

	bigINumResult, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint_20(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_20"

	numOfItemsUint := uint(0)

	numOfItemsChosenUint := uint(4)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint <= 0\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint_21(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_21"

	numOfItemsUint := uint(15)

	numOfItemsChosenUint := uint(0)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint <= 0\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint_22(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_22"

	numOfItemsUint := uint(5)

	numOfItemsChosenUint := uint(11)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint < numOfItemsChosenUint\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint_23(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_23"

	numOfItemsUint := uint(0)

	numOfItemsChosenUint := uint(4)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint <= 0\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint_24(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_24"

	numOfItemsUint := uint(15)

	numOfItemsChosenUint := uint(0)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint < numOfItemsChosenUint\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint_25(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint_25"

	numOfItemsUint := uint(0)

	numOfItemsChosenUint := uint(0)

	numOfItemsUintStr := strconv.FormatUint(uint64(numOfItemsUint), 10)

	numOfItemsChosenUintStr := strconv.FormatUint(uint64(numOfItemsChosenUint), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint(numOfItemsUint, numOfItemsChosenUint, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint, numOfItemsChosenUint, allowRepetitions)\n"+
			"numOfItemsUint= '%v'\n"+
			"numOfItemsChosenUint= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint <= 0 & numOfItemsChosenUint <= 0\n\n",
			ePrefix,
			numOfItemsUintStr,
			numOfItemsChosenUintStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_01(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_01"

	numOfItemsUint32 := uint32(3)

	numOfItemsChosenUint32 := uint32(2)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_02(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_02"

	numOfItemsUint32 := uint32(3)

	numOfItemsChosenUint32 := uint32(2)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_03(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_03"

	numOfItemsUint32 := uint32(10)

	numOfItemsChosenUint32 := uint32(3)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_04(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_04"

	numOfItemsUint32 := uint32(20)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_05(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_05"

	numOfItemsUint32 := uint32(52)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_06(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_06"

	numOfItemsUint32 := uint32(5)

	numOfItemsChosenUint32 := uint32(3)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_07(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_07"

	numOfItemsUint32 := uint32(20)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_08(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_08"

	numOfItemsUint32 := uint32(5)

	numOfItemsChosenUint32 := uint32(11)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_09(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_09"

	numOfItemsUint32 := uint32(11)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_10(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_10"

	numOfItemsUint32 := uint32(11)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_11(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_11"

	numOfItemsUint32 := uint32(56)

	numOfItemsChosenUint32 := uint32(5)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_12(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_12"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(3)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_13(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_13"

	numOfItemsUint32 := uint32(12)

	numOfItemsChosenUint32 := uint32(7)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_14(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_14"

	numOfItemsUint32 := uint32(18)

	numOfItemsChosenUint32 := uint32(8)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_15(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_15"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(9)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_16(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_16"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(9)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_17(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_17"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(1)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_18(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_18"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(1)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_19(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_19"

	numOfItemsUint32 := uint32(9)

	numOfItemsChosenUint32 := uint32(4)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "6561"

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

	bigINumResult, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
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

func TestProbability_PermutationsUint32_20(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_20"

	numOfItemsUint32 := uint32(0)

	numOfItemsChosenUint32 := uint32(4)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint32 <= 0\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_21(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_21"

	numOfItemsUint32 := uint32(15)

	numOfItemsChosenUint32 := uint32(0)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint32 <= 0\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_22(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_22"

	numOfItemsUint32 := uint32(5)

	numOfItemsChosenUint32 := uint32(11)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint32 < numOfItemsChosenUint32\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_23(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_23"

	numOfItemsUint32 := uint32(0)

	numOfItemsChosenUint32 := uint32(4)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint32 <= 0\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_24(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_24"

	numOfItemsUint32 := uint32(15)

	numOfItemsChosenUint32 := uint32(0)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint32 <= 0\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint32_25(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint32_25"

	numOfItemsUint32 := uint32(0)

	numOfItemsChosenUint32 := uint32(0)

	numOfItemsUint32Str := strconv.FormatUint(uint64(numOfItemsUint32), 10)

	numOfItemsChosenUint32Str := strconv.FormatUint(uint64(numOfItemsChosenUint32), 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint32(numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint32(\n"+
			"  numOfItemsUint32, numOfItemsChosenUint32, allowRepetitions)\n"+
			"numOfItemsUint32= '%v'\n"+
			"numOfItemsChosenUint32= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint32 <= 0 & numOfItemsChosenUint32 <= 0\n\n",
			ePrefix,
			numOfItemsUint32Str,
			numOfItemsChosenUint32Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_01(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_01"

	numOfItemsUint64 := uint64(3)

	numOfItemsChosenUint64 := uint64(2)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_02(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_02"

	numOfItemsUint64 := uint64(3)

	numOfItemsChosenUint64 := uint64(2)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_03(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_03"

	numOfItemsUint64 := uint64(10)

	numOfItemsChosenUint64 := uint64(3)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_04(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_04"

	numOfItemsUint64 := uint64(20)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_05(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_05"

	numOfItemsUint64 := uint64(52)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_06(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_06"

	numOfItemsUint64 := uint64(5)

	numOfItemsChosenUint64 := uint64(3)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_07(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_07"

	numOfItemsUint64 := uint64(20)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_08(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_08"

	numOfItemsUint64 := uint64(5)

	numOfItemsChosenUint64 := uint64(11)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_09(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_09"

	numOfItemsUint64 := uint64(11)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_10(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_10"

	numOfItemsUint64 := uint64(11)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_11(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_11"

	numOfItemsUint64 := uint64(56)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_12(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_12"

	numOfItemsUint64 := uint64(9)

	numOfItemsChosenUint64 := uint64(3)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_13(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_13"

	numOfItemsUint64 := uint64(12)

	numOfItemsChosenUint64 := uint64(7)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_14(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_14"

	numOfItemsUint64 := uint64(18)

	numOfItemsChosenUint64 := uint64(8)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_15(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_15"

	numOfItemsUint64 := uint64(9)

	numOfItemsChosenUint64 := uint64(9)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_16(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_16"

	numOfItemsUint64 := uint64(9)

	numOfItemsChosenUint64 := uint64(9)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_17(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_17"

	numOfItemsUint64 := uint64(9)

	numOfItemsChosenUint64 := uint64(1)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_18(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_18"

	numOfItemsUint64 := uint64(9)

	numOfItemsChosenUint64 := uint64(1)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_19(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_19"

	numOfItemsUint64 := uint64(11)

	numOfItemsChosenUint64 := uint64(5)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

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

	bigINumResult, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumResult, err := new(Probability).PermutationsUint(\n"+
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

func TestProbability_PermutationsUint64_20(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_20"

	numOfItemsUint64 := uint64(0)

	numOfItemsChosenUint64 := uint64(4)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint64 <= 0\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_21(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_21"

	numOfItemsUint64 := uint64(15)

	numOfItemsChosenUint64 := uint64(0)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint64 <= 0\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_22(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_22"

	numOfItemsUint64 := uint64(5)

	numOfItemsChosenUint64 := uint64(11)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint64 < numOfItemsChosenUint64\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_23(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_23"

	numOfItemsUint64 := uint64(0)

	numOfItemsChosenUint64 := uint64(4)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint64 <= 0\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_24(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_24"

	numOfItemsUint64 := uint64(15)

	numOfItemsChosenUint64 := uint64(0)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsChosenUint64 <= 0\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsUint64_25(t *testing.T) {

	ePrefix := "TestProbability_PermutationsUint64_25"

	numOfItemsUint64 := uint64(0)

	numOfItemsChosenUint64 := uint64(0)

	numOfItemsUint64Str := strconv.FormatUint(numOfItemsUint64, 10)

	numOfItemsChosenUint64Str := strconv.FormatUint(numOfItemsChosenUint64, 10)

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).PermutationsUint64(numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err := new(Probability).PermutationsUint64(\n"+
			"  numOfItemsUint64, numOfItemsChosenUint64, allowRepetitions)\n"+
			"numOfItemsUint64= '%v'\n"+
			"numOfItemsChosenUint64= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because:\n"+
			"  numOfItemsUint64 <= 0  &  numOfItemsChosenUint64 <= 0\n\n",
			ePrefix,
			numOfItemsUint64Str,
			numOfItemsChosenUint64Str,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}
