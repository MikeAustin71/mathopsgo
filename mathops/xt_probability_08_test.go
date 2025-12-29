package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestProbability_PermutationsBigIntNum_01(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_01"

	numOfItemsInt := 3
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 2
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_02(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_02"

	numOfItemsInt := 3
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 2
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_03(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_03"

	numOfItemsInt := 10
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 3
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_04(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_04"

	numOfItemsInt := 20
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 5
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_05(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_05"

	numOfItemsInt := 52
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 5
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_06(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_06"

	numOfItemsInt := 5
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 3
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_07(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_07"

	numOfItemsInt := 20
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 5
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_08(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_08"

	numOfItemsInt := 5
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 11
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_09(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_09"

	numOfItemsInt := 56
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 5
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_10(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_10"

	numOfItemsInt := 9
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 3
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_11(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_11"

	numOfItemsInt := 12
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 7
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_12(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_12"

	numOfItemsInt := 18
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 8
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_13(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_13"

	numOfItemsInt := 9
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 9
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_14(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_14"

	numOfItemsInt := 9
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 9
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_15(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_15"

	numOfItemsInt := 9
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 1
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_16(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_16"

	numOfItemsInt := 9
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 1
	numOfItemsChosenExponentInt := 0

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

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	bigIntNumResult, err := new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
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

func TestProbability_PermutationsBigIntNum_17(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_17"

	numOfItemsInt := 0
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 4
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_18(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_18"

	numOfItemsInt := 15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 0
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_19(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_19"

	numOfItemsInt := -15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 2
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_20(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_20"

	numOfItemsInt := 15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := -2
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_21(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_21"

	numOfItemsInt := 5
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 11
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems < bigINumNumOfItemsChosen\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_22(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_22"

	numOfItemsInt := 0
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 4
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_23(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_23"

	numOfItemsInt := 15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 0
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_24(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_24"

	numOfItemsInt := -15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := 2
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsBigIntNum_25(t *testing.T) {

	ePrefix := "TestProbability_PermutationsBigIntNum_25"

	numOfItemsInt := 15
	numOfItemsExponentInt := 0

	numOfItemsChosenInt := -2
	numOfItemsChosenExponentInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewIntExponent(numOfItemsInt, numOfItemsExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"BigINumNumOfItems, err := new(BigIntNum).NewIntExponent(\n"+
			"  numOfItemsInt, numOfItemsExponentInt)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsExponentInt,
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewIntExponent(numOfItemsChosenInt, numOfItemsChosenExponentInt)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, numOfItemsChosenExponentInt, err.Error())
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

	_, err = new(Probability).PermutationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_01(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_01"

	numOfItemsInt := 3
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "6"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_02(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_02"

	numOfItemsInt := 3
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "9"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_03(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_03"

	numOfItemsInt := 10
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 3
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "1000"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_04(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_04"

	numOfItemsInt := 20
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 5
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1860480"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_05(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_05"

	numOfItemsInt := 52
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 5
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "311875200"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_06(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_06"

	numOfItemsInt := 5
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 3
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "125"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_07(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_07"

	numOfItemsInt := 20
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 5
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "3200000"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_08(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_08"

	numOfItemsInt := 5
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 11
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "48828125"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_09(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_09"

	numOfItemsInt := 56
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 5
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "458377920"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_10(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_10"

	numOfItemsInt := 9
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 3
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "504"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_11(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_11"

	numOfItemsInt := 12
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 7
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "3991680"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_12(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_12"

	numOfItemsInt := 18
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 8
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1764322560"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_13(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_13"

	numOfItemsInt := 9
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 9
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "362880"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_14(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_14"

	numOfItemsInt := 9
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 9
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "387420489"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_15(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_15"

	numOfItemsInt := 9
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 1
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "9"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_16(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_16"

	numOfItemsInt := 9
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 1
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "9"

	expectedAbsoluteAllDigitsStr := expectedNumStr

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	decimalResult, err := new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions),
			err.Error())

		return
	}

	err = decimalResult.IsValid("Validating decimalResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := decimalResult.IsValid('Validating decimalResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultNumStr, err := decimalResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumStr, err := decimalResult.GetNumStr()\n"+
			"decimalResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decimalResultPrecisionInt, err := decimalResult.GetPrecision()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionInt, err :=\n"+
			"  decimalResult.GetPrecisionInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultPrecisionUint, err := decimalResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultPrecisionUint, err :=\n"+
			"  decimalResult.GetPrecisionUint()\n"+
			"decimalResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultScaleFactorBigInt, err := decimalResult.GetScaleVal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultScaleFactorBigInt, err :=\n"+
			"  decimalResult.GetScaleVal()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultSignValue, err := decimalResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultSignValue, err := decimalResult.GetSign()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultNumSeps, err := decimalResult.GetNumericSeparatorsDto()\n"+
			"decimalResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultBigInt, err := decimalResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultBigInt, err :=\n"+
			"  decimalResult.GetBigInt()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	decimalResultAbsoluteAllDigitsStr, err := decimalResult.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResultAbsoluteAllDigitsStr, err :=\n"+
			"  decimalResult.GetAbsoluteAllDigitsStr()\n"+
			"decimalResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decimalResultNumStr, err.Error())
		return
	}

	if expectedNumStr != decimalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != decimalResultNumStr\n"+
			"Expected decimalResultNumStr = '%v'\n"+
			"  Actual decimalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, decimalResultNumStr)

		return
	}

	if expectedPrecisionInt != decimalResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Integers DON'T MATCH!\n"+
			"Because expectedPrecisionInt != decimalResultPrecisionInt\n"+
			"Expected decimalResultPrecisionInt = '%v'\n"+
			"  Actual decimalResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, decimalResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != decimalResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Precision Uint's DON'T MATCH!\n"+
			"Because expectedPrecisionUint != decimalResultPrecisionUint\n"+
			"Expected decimalResultPrecisionUint = '%v'\n"+
			"  Actual decimalResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decimalResultPrecisionUint)

		return
	}

	if expectedSignValue != decimalResultSignValue {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Sign Values DON'T MATCH!\n"+
			"Because expectedSignValue != decimalResultSignValue\n"+
			"Expected decimalResultSignValue = '%v'\n"+
			"  Actual decimalResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, decimalResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(decimalResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Numeric Separators DON'T MATCH!\n"+
			"Because expectedNumSeps != decimalResultNumSeps \n"+
			"Expected decimalResultNumSeps = '%v'\n"+
			"  Actual decimalResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decimalResultNumSeps.String())

		return
	}

	if expectedScaleFactorBigInt.Cmp(decimalResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult Scale Factors DON'T MATCH!\n"+
			"Because expectedScaleFactorBigInt!=decimalResultScaleFactorBigInt\n"+
			"Expected decimalResultScaleFactorBigInt = '%v'\n"+
			"  Actual decimalResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			decimalResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Digit Strings DON'T MATCH!\n"+
			"Because expectedAbsoluteAllDigitsStr != decimalResultAbsoluteAllDigitsStr\n"+
			"Expected decimalResultAbsoluteAllDigitsStr = '%v'\n"+
			"  Actual decimalResultAbsoluteAllDigitsStr = '%v'\n\n",
			ePrefix, expectedAbsoluteAllDigitsStr, decimalResultAbsoluteAllDigitsStr)

		return
	}

	if expectedBigInt.Cmp(decimalResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual decimalResult BigInt Values DON'T MATCH!\n"+
			"Because expectedBigInt.Cmp(decimalResultBigInt) != 0\n"+
			"Expected decimalResultBigInt = '%v'\n"+
			"  Actual decimalResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), decimalResultBigInt.Text(10))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_17(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_17"

	numOfItemsInt := 0
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 4
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_18(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_18"

	numOfItemsInt := 15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 0
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_19(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_19"

	numOfItemsInt := -15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_20(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_20"

	numOfItemsInt := 15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := -2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_21(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_21"

	numOfItemsInt := 5
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 11
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems < decimalNumOfItemsChosen\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_22(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_22"

	numOfItemsInt := 0
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 4
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems < decimalNumOfItemsChosen\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_23(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_23"

	numOfItemsInt := 15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 0
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_24(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_24"

	numOfItemsInt := -15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := 2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_PermutationsDecimal_25(t *testing.T) {

	ePrefix := "TestProbability_PermutationsDecimal_25"

	numOfItemsInt := 15
	numOfItemsPrecisionUint := uint(0)

	numOfItemsChosenInt := -2
	numOfItemsChosenPrecisionUint := uint(0)

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, numOfItemsPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).NewInt(\n"+
			"  numOfItemsInt, numOfItemsPrecisionUint)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, numOfItemsChosenPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).NewInt(\n"+
			"  numOfItemsChosenInt, numOfItemsChosenPrecisionUint)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"numOfItemsChosenPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsChosenInt,
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

	_, err = new(Probability).PermutationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"_, err = new(Probability).PermutationsDecimal(\n"+
			"  decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}
