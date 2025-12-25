package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestProbability_CombinationsNoRepsBigInt_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_01"

	numOfItemsInt := 16

	numOfItemsChosenInt := 3

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "560"

	expectedBigInt := big.NewInt(560)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = false

	expectedNumAbsIntStr := "560"

	expectedNumAbsFracStr := ""

	expectedAbsAllRunesNumStr := "560"

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

	numStrDtoBigIntNumResult, err := bigIntNumResult.GetNumStrDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigIntNumResult, err :=\n"+
			"  bigIntNumResult.GetNumStrDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumStr, err.Error())
		return
	}

	err = numStrDtoBigIntNumResult.IsValid("Validating numStrDtoBigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoBigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoBigIntNumResultNumStr, err := numStrDtoBigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigIntNumResultNumStr, err :=\n"+
			"  numStrDtoBigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != numStrDtoBigIntNumResultNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDto Extracted from bigIntNumResult is INVALID!\n"+
			"Because expectedNumStr != numStrDtoBigIntNumResultNumStr\n"+
			"Expected numStrDtoBigIntNumResultNumStr = '%v'\n"+
			"  Actual numStrDtoBigIntNumResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoBigIntNumResultNumStr)

		return
	}

	bigIntNumResultHasNumericDigits := numStrDtoBigIntNumResult.HasNumericDigits()

	bigIntNumResultIsFractionalValue := numStrDtoBigIntNumResult.IsFractionalValue()

	bigIntNumResultAbsIntRunes, err := numStrDtoBigIntNumResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultAbsIntRunes, err := \n"+
			"  numStrDtoBigIntNumResult.GetAbsIntRunes()\n"+
			"numStrDtoBigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoBigIntNumResultNumStr, err.Error())
		return
	}

	bigIntNumResultAbsIntStr := string(bigIntNumResultAbsIntRunes)

	bigIntNumResultAbsFracRunes, err := numStrDtoBigIntNumResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultAbsFracRunes, err :=\n"+
			"  numStrDtoBigIntNumResult.GetAbsFracRunes()\n"+
			"numStrDtoBigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoBigIntNumResultNumStr, err.Error())
		return
	}

	bigIntNumResultAbsAllRunes, err := numStrDtoBigIntNumResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultAbsAllRunes, err := \n"+
			"  numStrDtoBigIntNumResult.GetAbsAllNumRunes()\n"+
			"numStrDtoBigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoBigIntNumResultNumStr, err.Error())
		return
	}

	bigIntNumResultAbsAllRunesNumStr := string(bigIntNumResultAbsAllRunes)

	bigIntNumResultAbsFracStr := string(bigIntNumResultAbsFracRunes)

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

	if expectedNumHasNumericDigits != bigIntNumResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Has Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != bigIntNumResultHasNumericDigits\n"+
			"Expected bigIntNumResultHasNumericDigits = '%v'\n"+
			"  Actual bigIntNumResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, bigIntNumResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != bigIntNumResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != bigIntNumResultIsFractionalValue\n"+
			"Expected bigIntNumResultIsFractionalValue = '%v'\n"+
			"  Actual bigIntNumResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, bigIntNumResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != bigIntNumResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual bigIntNumResult Absolute Integer Strings DON'T MATCH!\n"+
			"Because expectedNumAbsIntStr != bigIntNumResultAbsIntStr\n"+
			"Expected bigIntNumResultAbsIntStr = '%v'\n"+
			"  Actual bigIntNumResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, bigIntNumResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != bigIntNumResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute Fractional Strings DON'T MATCH\n"+
			"Because expectedNumAbsFracStr != bigIntNumResultAbsFracStr\n"+
			"Expected bigIntNumResultAbsFracStr = '%v'\n"+
			"  Actual bigIntNumResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, bigIntNumResultAbsFracStr)

		return
	}

	if expectedAbsAllRunesNumStr != bigIntNumResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Absolute All Runes Number Strings DON'T MATCH!\n"+
			"Because expectedAbsAllRunesNumStr != bigIntNumResultAbsAllRunesNumStr\n"+
			"Expected bigIntNumResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual bigIntNumResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, bigIntNumResultAbsAllRunesNumStr)

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

func TestProbability_CombinationsNoRepsBigInt_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_02"

	numOfItemsInt := 16

	numOfItemsChosenInt := 12

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "1820"

	expectedBigInt := big.NewInt(1820)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_03"

	numOfItemsInt := 52

	numOfItemsChosenInt := 5

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "2598960"

	expectedBigInt := big.NewInt(2598960)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_04"

	numOfItemsInt := 52

	numOfItemsChosenInt := 26

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "495918532948104"

	expectedBigInt := big.NewInt(495918532948104)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_05"

	numOfItemsInt := 18

	numOfItemsChosenInt := 7

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "31824"

	expectedBigInt := big.NewInt(31824)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_06"

	numOfItemsInt := 22

	numOfItemsChosenInt := 5

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "26334"

	expectedBigInt := big.NewInt(26334)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_07"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "3819816"

	expectedBigInt := big.NewInt(3819816)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_08"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "12"

	expectedBigInt := big.NewInt(12)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_09"

	numOfItemsInt := 25

	numOfItemsChosenInt := 25

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "1"

	expectedBigInt := big.NewInt(1)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_10"

	numOfItemsInt := 25

	numOfItemsChosenInt := 1

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "25"

	expectedBigInt := big.NewInt(25)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).\n"+
			"  CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems.Text(10),
			numOfItemsChosen.Text(10),
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

func TestProbability_CombinationsNoRepsBigInt_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsNoRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItems < numOfItemsChosen\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsNoRepsBigInt_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsNoRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItemsChosen == 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsNoRepsBigInt_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 26

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsNoRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItems = 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsNoRepsBigInt_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_14"

	numOfItemsInt := -52

	numOfItemsChosenInt := 26

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsNoRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItems < 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsNoRepsBigInt_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsNoRepsBigInt_15"

	numOfItemsInt := 52

	numOfItemsChosenInt := -26

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsNoRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItemsChosen < 0. numOfItems='%v' numOfItemsChosen='%v'",
			numOfItemsInt, numOfItemsChosenInt)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsNoRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItemsChosen < 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsWithRepsBigInt_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_01"

	numOfItemsInt := 5

	numOfItemsChosenInt := 3

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "35"

	expectedBigInt := big.NewInt(35)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_02"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "705432"

	expectedBigInt := big.NewInt(705432)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_03"

	numOfItemsInt := 26

	numOfItemsChosenInt := 2

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "351"

	expectedBigInt := big.NewInt(351)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_04"

	numOfItemsInt := 26

	numOfItemsChosenInt := 24

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "63205303218876"

	expectedBigInt := big.NewInt(63205303218876)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_05"

	numOfItemsInt := 10

	numOfItemsChosenInt := 14

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "817190"

	expectedBigInt := big.NewInt(817190)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_06"

	numOfItemsInt := 12

	numOfItemsChosenInt := 15

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "7726160"

	expectedBigInt := big.NewInt(7726160)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_07"

	numOfItemsInt := 7

	numOfItemsChosenInt := 3

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "84"

	expectedBigInt := big.NewInt(84)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_08"

	numOfItemsInt := 3

	numOfItemsChosenInt := 7

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "36"

	expectedBigInt := big.NewInt(36)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_09"

	numOfItemsInt := 62

	numOfItemsChosenInt := 5

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "8936928"

	expectedBigInt := big.NewInt(8936928)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_10"

	numOfItemsInt := 97

	numOfItemsChosenInt := 5

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "79208745"

	expectedBigInt := big.NewInt(79208745)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_11"

	numOfItemsInt := 15

	numOfItemsChosenInt := 15

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "77558760"

	expectedBigInt := big.NewInt(77558760)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_12"

	numOfItemsInt := 12

	numOfItemsChosenInt := 1

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	expectedNumStr := "12"

	expectedBigInt := big.NewInt(12)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"  numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItems,
			numOfItemsChosen,
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

func TestProbability_CombinationsWithRepsBigInt_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 15

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItems == 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsWithRepsBigInt_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_14"

	numOfItemsInt := 12

	numOfItemsChosenInt := 0

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItemsChosen == 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsWithRepsBigInt_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_15"

	numOfItemsInt := -12

	numOfItemsChosenInt := 6

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItems < 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsWithRepsBigInt_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsWithRepsBigInt_16"

	numOfItemsInt := 12

	numOfItemsChosenInt := -6

	numOfItems := big.NewInt(int64(numOfItemsInt))

	numOfItemsChosen := big.NewInt(int64(numOfItemsChosenInt))

	_, err := new(Probability).CombinationsWithRepsBigInt(numOfItems, numOfItemsChosen)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := new(Probability).CombinationsWithRepsBigInt(\n"+
			"numOfItems, numOfItemsChosen)\n"+
			"numOfItems= '%v'\n"+
			"numOfItemsChosen= '%v'\n"+
			"Error should have triggered because: numOfItemsChosen < 0\n\n",
			ePrefix, numOfItems.Text(10), numOfItemsChosen.Text(10))

		return
	}

	return
}

func TestProbability_CombinationsBigIntNum_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_01"

	numOfItemsInt := 16

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenInt := 3

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_02"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_03"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_04"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_05"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_06"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_07"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_08"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_09"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_10"

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen == 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsBigIntNum_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems == 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsBigIntNum_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_14"

	numOfItemsInt := -52

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_15"

	numOfItemsInt := 52

	numOfItemsChosenInt := -26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_16"

	numOfItemsInt := 5

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_17(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_17"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_18(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_18"

	numOfItemsInt := 26

	numOfItemsChosenInt := 2

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_19(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_19"

	numOfItemsInt := 26

	numOfItemsChosenInt := 24

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_20(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_20"

	numOfItemsInt := 10

	numOfItemsChosenInt := 14

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_21(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_21"

	numOfItemsInt := 12

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_22(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_22"

	numOfItemsInt := 7

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_23(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_23"

	numOfItemsInt := 3

	numOfItemsChosenInt := 7

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_24(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_24"

	numOfItemsInt := 62

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_25(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_25"

	numOfItemsInt := 97

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_26(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_26"

	numOfItemsInt := 15

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_27(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_27"

	numOfItemsInt := 12

	numOfItemsChosenInt := 1

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

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

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_28(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_28"

	numOfItemsInt := 0

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems == 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsBigIntNum_29(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_29"

	numOfItemsInt := 12

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
			"  bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen == 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsBigIntNum_30(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_30"

	numOfItemsInt := -12

	numOfItemsChosenInt := 6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
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

func TestProbability_CombinationsBigIntNum_31(t *testing.T) {

	ePrefix := "TestProbability_CombinationsBigIntNum_31"

	numOfItemsInt := 12

	numOfItemsChosenInt := -6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	bigINumNumOfItems, err := new(BigIntNum).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItems, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigINumNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsBigIntNum(bigINumNumOfItems, bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsBigIntNum(\n"+
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
