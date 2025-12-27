package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestProbability_CombinationsINumMgr_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_01"

	numOfItemsInt := 16

	numOfItemsChosenInt := 3

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_02"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&intAryNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
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

func TestProbability_CombinationsINumMgr_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_03"

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

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsBigIntNum(\n"+
			"  &numStrDtoNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"numStrDtoNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoNumOfItemsNumStr,
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

func TestProbability_CombinationsINumMgr_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_04"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_05"

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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_06"

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

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	numStrDtoNumOfItemsChosen, err := new(BigIntNum).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(BigIntNum).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_07"

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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)\n"+
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

func TestProbability_CombinationsINumMgr_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_08"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
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

func TestProbability_CombinationsINumMgr_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_09"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
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

func TestProbability_CombinationsINumMgr_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_10"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&intAryNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
			"numStrDtoNumOfItems= '%v'\n"+
			"numStrDtoNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numStrDtoNumOfItems < numStrDtoNumOfItemsChosen\n\n",
			ePrefix,
			numStrDtoNumOfItemsNumStr,
			numStrDtoNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	numOfItemsIntStr := strconv.Itoa(numOfItemsInt)

	numOfItemsChosenIntStr := strconv.Itoa(numOfItemsChosenInt)

	var allowRepetitions bool

	allowRepetitions = false

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &numStrDtoNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"numStrDtoNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			numStrDtoNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_13"

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

	intAryNumOfItemsChosen, err := new(IntAry).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItemsChosen, err := new(IntAry).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&decimalNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"intAryNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: decimalNumOfItems <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			intAryNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_14"

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
			ePrefix, numOfItemsIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: intAryNumOfItems <= 0\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_15"

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

	_, err = new(Probability).CombinationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			bigINumNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_16"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_17(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_17"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&intAryNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			bigINumNumOfItemsChosen,
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

func TestProbability_CombinationsINumMgr_18(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_18"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_19(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_19"

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

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"numStrDtoNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
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

func TestProbability_CombinationsINumMgr_20(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_20"

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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
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

func TestProbability_CombinationsINumMgr_21(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_21"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			bigINumNumOfItemsChosen,
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

func TestProbability_CombinationsINumMgr_22(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_22"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&decimalNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
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

func TestProbability_CombinationsINumMgr_23(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_23"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&intAryNumOfItems, &intAryNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_24(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_24"

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

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
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

func TestProbability_CombinationsINumMgr_25(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_25"

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

	numStrDtoNumOfItems, err := new(NumStrDto).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItems, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&numStrDtoNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"numStrDtoNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
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

func TestProbability_CombinationsINumMgr_26(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_26"

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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"intAryNumOfItemsChosen= '%v'\n"+
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

func TestProbability_CombinationsINumMgr_27(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_27"

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

	intAryNumOfItems, err := new(IntAry).NewInt(numOfItemsInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumOfItems, err := new(IntAry).\n"+
			"  NewInt(numOfItemsInt, 0)\n"+
			"numOfItemsInt= '%v'\n"+
			"precision= '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsIntStr, err.Error())
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

	bigIntNumResult, err := new(Probability).CombinationsINumMgr(&intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
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

func TestProbability_CombinationsINumMgr_28(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_28"

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

	_, err = new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("Error: Expected an error to be returned. Instead err==nil "+
			"numOfItems == 0. numOfItems='%v' numOfItemsChosen='%v' allowRepetitions='%v' ",
			numOfItemsInt, numOfItemsChosenInt, allowRepetitions)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &bigINumNumOfItems, &decimalNumOfItemsChosen, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"decimalNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: bigINumNumOfItems <= 0\n\n",
			ePrefix,
			bigINumNumOfItemsNumStr,
			decimalNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return

}

func TestProbability_CombinationsINumMgr_29(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_29"

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

	_, err = new(Probability).CombinationsINumMgr(&bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &bigINumNumOfItems, &bigINumNumOfItemsChosen, allowRepetitions)\n"+
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

func TestProbability_CombinationsINumMgr_30(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_30"

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
			ePrefix, numOfItemsIntStr, err.Error())
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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&intAryNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &intAryNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
			"intAryNumOfItems= '%v'\n"+
			"numStrDtoNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numStrDtoNumOfItems <= 0 \n\n",
			ePrefix,
			intAryNumOfItemsNumStr,
			numStrDtoNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsINumMgr_31(t *testing.T) {

	ePrefix := "TestProbability_CombinationsINumMgr_31"

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

	numStrDtoNumOfItemsChosen, err := new(NumStrDto).NewInt(numOfItemsChosenInt, 0)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoNumOfItemsChosen, err := new(NumStrDto).\n"+
			"  NewInt(numOfItemsChosenInt, 0)\n"+
			"numOfItemsChosenInt = '%v'\n"+
			"precision = '0'\n"+
			"Error= '%v'\n\n",
			ePrefix, numOfItemsChosenIntStr, err.Error())
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

	_, err = new(Probability).CombinationsINumMgr(&decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsINumMgr(\n"+
			"  &decimalNumOfItems, &numStrDtoNumOfItemsChosen, allowRepetitions)\n"+
			"decimalNumOfItems= '%v'\n"+
			"numStrDtoNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numStrDtoNumOfItemsChosen <= 0\n\n",
			ePrefix,
			decimalNumOfItemsNumStr,
			numStrDtoNumOfItemsChosenNumStr,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_01(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_01"

	numOfItemsInt := 16

	numOfItemsChosenInt := 3

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_02(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_02"

	numOfItemsInt := 16

	numOfItemsChosenInt := 12

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_03(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_03"

	numOfItemsInt := 52

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_04(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_04"

	numOfItemsInt := 52

	numOfItemsChosenInt := 26

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_05(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_05"

	numOfItemsInt := 18

	numOfItemsChosenInt := 7

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_06(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_01"

	numOfItemsInt := 22

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_07(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_07"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_08(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_08"

	numOfItemsInt := 56

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_09(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_09"

	numOfItemsInt := 25

	numOfItemsChosenInt := 25

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_10(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_10"

	numOfItemsInt := 25

	numOfItemsChosenInt := 1

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_11(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_11"

	numOfItemsInt := 26

	numOfItemsChosenInt := 52

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt < numOfItemsChosenInt\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_12(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_12"

	numOfItemsInt := 52

	numOfItemsChosenInt := 0

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt <= 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_13(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_13"

	numOfItemsInt := 0

	numOfItemsChosenInt := 26

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt < numOfItemsChosenInt\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_14(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_14"

	numOfItemsInt := -52

	numOfItemsChosenInt := 26

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt <= 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_15(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_15"

	numOfItemsInt := 52

	numOfItemsChosenInt := -26

	var allowRepetitions bool

	allowRepetitions = false

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt <= 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_16(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_16"

	numOfItemsInt := 5

	numOfItemsChosenInt := 3

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_17(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_17"

	numOfItemsInt := 12

	numOfItemsChosenInt := 11

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_18(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_18"

	numOfItemsInt := 26

	numOfItemsChosenInt := 2

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_19(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_19"

	numOfItemsInt := 26

	numOfItemsChosenInt := 24

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_20(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_20"

	numOfItemsInt := 10

	numOfItemsChosenInt := 14

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_21(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_21"

	numOfItemsInt := 12

	numOfItemsChosenInt := 15

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_22(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_22"

	numOfItemsInt := 7

	numOfItemsChosenInt := 3

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_23(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_23"

	numOfItemsInt := 3

	numOfItemsChosenInt := 7

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_24(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_24"

	numOfItemsInt := 62

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_25(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_25"

	numOfItemsInt := 97

	numOfItemsChosenInt := 5

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_26(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_26"

	numOfItemsInt := 15

	numOfItemsChosenInt := 15

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_27(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_27"

	numOfItemsInt := 12

	numOfItemsChosenInt := 1

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

	bigIntNumResult, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"bigINumNumOfItems= '%v'\n"+
			"bigINumNumOfItemsChosen= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
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

func TestProbability_CombinationsInt_28(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_28"

	numOfItemsInt := 0

	numOfItemsChosenInt := 15

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsInt == 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_29(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_29"

	numOfItemsInt := 12

	numOfItemsChosenInt := 0

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt == 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_30(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_30"

	numOfItemsInt := -12

	numOfItemsChosenInt := 6

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsInt <= 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}

func TestProbability_CombinationsInt_31(t *testing.T) {

	ePrefix := "TestProbability_CombinationsInt_31"

	numOfItemsInt := 12

	numOfItemsChosenInt := -6

	var allowRepetitions bool

	allowRepetitions = true

	_, err := new(Probability).CombinationsInt(numOfItemsInt, numOfItemsChosenInt, allowRepetitions)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"_, err = new(Probability).CombinationsInt(\n"+
			"  numOfItemsInt, numOfItemsChosenInt, allowRepetitions)\n"+
			"numOfItemsInt= '%v'\n"+
			"numOfItemsChosenInt= '%v'\n"+
			"allowRepetitions= '%v'\n"+
			"Error should have triggered because: numOfItemsChosenInt <= 0\n\n",
			ePrefix,
			numOfItemsInt,
			numOfItemsChosenInt,
			strconv.FormatBool(allowRepetitions))

		return
	}

	return
}
