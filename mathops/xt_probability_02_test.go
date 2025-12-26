package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestProbability_CombinationsDecimal_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_01"

	numOfItemsInt := 16

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "560"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_02"

	numOfItemsInt := 16

	numOfItemsChosenInt := 12

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1820"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_03"

	numOfItemsInt := 52

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "2598960"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_04"

	numOfItemsInt := 52

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "495918532948104"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_05"

	numOfItemsInt := 18

	numOfItemsChosenInt := 7

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "31824"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_06"

	numOfItemsInt := 22

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "26334"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_07"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "3819816"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_08"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "3819816"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_09"

	numOfItemsInt := 25

	numOfItemsChosenInt := 25

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_10"

	numOfItemsInt := 25

	numOfItemsChosenInt := 1

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "25"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_14"

	numOfItemsInt := -52

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_15"

	numOfItemsInt := 52

	numOfItemsChosenInt := -26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_16"

	numOfItemsInt := 5

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "35"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_17(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_17"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "705432"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_18(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_18"

	numOfItemsInt := 26

	numOfItemsChosenInt := 2

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "351"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_19(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_19"

	numOfItemsInt := 26

	numOfItemsChosenInt := 24

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "63205303218876"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_20(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_20"

	numOfItemsInt := 10

	numOfItemsChosenInt := 14

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "817190"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_21(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_21"

	numOfItemsInt := 12

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "7726160"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_22(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_22"

	numOfItemsInt := 7

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "84"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_23(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_23"

	numOfItemsInt := 3

	numOfItemsChosenInt := 7

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "36"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_24(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_24"

	numOfItemsInt := 62

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "8936928"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_25(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_25"

	numOfItemsInt := 97

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "79208745"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_26(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_26"

	numOfItemsInt := 15

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "77558760"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_27(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_27"

	numOfItemsInt := 12

	numOfItemsChosenInt := 1

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "12"

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

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	decimalResult, err := new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_28(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_28"

	numOfItemsInt := 0

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_29(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_29"

	numOfItemsInt := 12

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_30(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_30"

	numOfItemsInt := -12

	numOfItemsChosenInt := 6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsDecimal_31(t *testing.T) {

	ePrefix := "TestProbability_CombinationsDecimal_31"

	numOfItemsInt := 12

	numOfItemsChosenInt := -6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	decimalNumOfItems, err := new(Decimal).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItems, err := new(Decimal).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	decimalNumOfItemsChosen, err := new(Decimal).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decimalNumOfItemsChosen, err := new(Decimal).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsDecimal(decimalNumOfItems, decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_01"

	numOfItemsInt := 16

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "560"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_02"

	numOfItemsInt := 16

	numOfItemsChosenInt := 12

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1820"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_03"

	numOfItemsInt := 52

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "2598960"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_04"

	numOfItemsInt := 52

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "495918532948104"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_05"

	numOfItemsInt := 18

	numOfItemsChosenInt := 7

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "31824"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_06"

	numOfItemsInt := 22

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "26334"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_07"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "3819816"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_08"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "3819816"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_09"

	numOfItemsInt := 25

	numOfItemsChosenInt := 25

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "1"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_10"

	numOfItemsInt := 25

	numOfItemsChosenInt := 1

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	expectedNumStr := "25"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead  err==nil "+
			"numOfItems < numOfItemsChosen. numOfItems='%v' numOfItemsChosen='%v' "+
			"allowRepetitions='%v'",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
			"  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"intAryNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: intAryNumOfItemsChosen <=0\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			intAryNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsIntAry_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
			"  intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"intAryNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: intAryNumOfItems <=0\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			intAryNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsIntAry_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_14"

	numOfItemsInt := -52

	numOfItemsChosenInt := 26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_15"

	numOfItemsInt := 52

	numOfItemsChosenInt := -26

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_16"

	numOfItemsInt := 5

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "35"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_17(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_17"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "705432"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_18(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_18"

	numOfItemsInt := 26

	numOfItemsChosenInt := 2

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "351"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_19(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_19"

	numOfItemsInt := 26

	numOfItemsChosenInt := 24

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "63205303218876"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_20(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_20"

	numOfItemsInt := 10

	numOfItemsChosenInt := 14

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "817190"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_21(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_21"

	numOfItemsInt := 12

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "7726160"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_22(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_22"

	numOfItemsInt := 7

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "84"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_23(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_23"

	numOfItemsInt := 3

	numOfItemsChosenInt := 7

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "36"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_24(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_24"

	numOfItemsInt := 62

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "8936928"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_25(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_25"

	numOfItemsInt := 97

	numOfItemsChosenInt := 5

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "79208745"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_26(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_26"

	numOfItemsInt := 15

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "77558760"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_27(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_27"

	numOfItemsInt := 12

	numOfItemsChosenInt := 1

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	expectedNumStr := "12"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	intAryResult, err := new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(Probability).CombinationsDecimal(\n"+
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

func TestProbability_CombinationsIntAry_28(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_28"

	numOfItemsInt := 0

	numOfItemsChosenInt := 15

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_29(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_29"

	numOfItemsInt := 12

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_30(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_30"

	numOfItemsInt := -12

	numOfItemsChosenInt := 6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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

func TestProbability_CombinationsIntAry_31(t *testing.T) {

	ePrefix := "TestProbability_CombinationsIntAry_31"

	numOfItemsInt := 12

	numOfItemsChosenInt := -6

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = true

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenInt, err.Error())
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

	_, err = new(Probability).CombinationsIntAry(intAryNumOfItems, intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsIntAry(\n"+
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
