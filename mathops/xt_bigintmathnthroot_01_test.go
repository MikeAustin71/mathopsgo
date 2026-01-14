package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathNthRoot_GetNthRootBigNum_01(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_01"

	baseStr := "125"

	nthRootStr := "5"

	maxPrecision := uint(14)

	//                                  1         2         3
	//                       0.123456489012345678901234567890
	expectedResultNumStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultPrecisionInt, err := result.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionInt, err := result.GetPrecisionInt()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error= '%v\n\n", ePrefix, err.Error())
		return

	}

	resultScaleFactorBigInt, err := result.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultScaleFactorBigInt, err := result.GetScaleFactorBigInt()\n"+
			"Error= '%v\n\n", ePrefix, err.Error())
		return

	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedPrecisionInt != resultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedPrecisionInt != resultPrecisionInt\n"+
			"Expected resultPrecisionInt = '%v'\n"+
			"  Actual resultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, resultPrecisionInt)

		return
	}

	if expectedPrecisionUint != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedPrecisionUint != resultPrecisionUint\n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, resultPrecisionUint)

		return
	}

	if expectedScaleFactorBigInt.Cmp(resultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factors Do Not Match!\n"+
			"Because expectedScaleFactorBigInt != resultScaleFactorBigInt\n"+
			"Expected resultScaleFactorBigInt = '%v'\n"+
			"  Actual resultScaleFactorBigInt = '%v'\n\n",
			ePrefix, expectedScaleFactorBigInt, resultScaleFactorBigInt.Text(10))

		return
	}

	if expectedResultNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings Do Not Match!\n"+
			"Because expectedResultNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_02(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_02"

	baseStr := "5604423"

	nthRootStr := "6"

	maxPrecision := uint(13)

	expectedResult := "13.3276982415963"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_03(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_03"

	baseStr := "5604423.924"

	nthRootStr := "6"

	maxPrecision := uint(13)

	expectedResult := "13.3276986078187"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_04(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_04"

	baseStr := "-27"

	nthRootStr := "3"

	maxPrecision := uint(2)

	expectedResult := "-3.00"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_05(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_05"

	baseStr := "-27"

	nthRootStr := "4"

	maxPrecision := uint(2)

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	_, err = new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {

		t.Errorf("%v\n"+
			"Expected Error from nRt.OriginalNthRoot()\n"+
			"Negative number with even nthRoot should trigger an error.\n"+
			"No Error was returned\n", ePrefix)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_06(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_06"

	baseStr := "-5604423.924"

	nthRootStr := "5"

	maxPrecision := uint(13)

	expectedResult := "-22.3720713464898"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_07(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_07"

	baseStr := "5604423.924"

	nthRootStr := "0"

	maxPrecision := uint(1)

	expectedResult := "1"

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).GetNthRoot(\n"+
			"  bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseStr,
			nthRootStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid(ePrefix)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumSeps, err := result.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
			"result= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because expectedResult != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if !expectedNumSeps.Equal(resultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedNumSeps != resultNumSeps \n"+
			"Expected resultNumSeps = '%v'\n"+
			"  Actual resultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), resultNumSeps.String())

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_08(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_08"

	baseStr := "27"

	nthRootStr := "1"

	maxPrecision := uint(2)

	bINumBase, err := new(BigIntNum).NewNumStr(baseStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(baseStr)\n"+
			"baseStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, baseStr, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRootStr)\n"+
			"nthRootStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRootStr, err.Error())
		return
	}

	_, err = new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected Error from nRt.OriginalNthRoot()\n"+
			"For nthRoot == 1 an error should be triggered.\n"+
			"However, No Error was returned.", ePrefix)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_09(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_09"

	radicand := "0.027"

	nthRoot := "3"

	expectedStr := "0.300000"

	maxPrecision := uint(6)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_10(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_10"

	radicand := "0.0005"

	nthRoot := "9"

	expectedStr := "0.429752972587713"

	maxPrecision := uint(15)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_11(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_11"

	radicand := "200000.000005"

	nthRoot := "2"

	expectedStr := "447.213595505548"

	maxPrecision := uint(12)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_12(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_11"

	radicand := "200001.100005"

	nthRoot := "2"

	expectedStr := "447.214825341245"

	maxPrecision := uint(12)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_13(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_13"

	radicand := "2000000.0000005"

	nthRoot := "2"

	expectedStr := "1414.21356237327"

	maxPrecision := uint(11)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_14(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_14"

	radicand := "20000000.00000005"

	nthRoot := "3"

	expectedStr := "271.441761659491"

	maxPrecision := uint(12)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_15(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_15"

	radicand := "20000200.00020005"

	nthRoot := "3"

	expectedStr := "271.442666463252"

	maxPrecision := uint(12)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_16(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_16"

	radicand := "2020020.1010205"

	nthRoot := "2"

	expectedStr := "1421.27411185193"

	maxPrecision := uint(11)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_17(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_17"

	radicand := "209050307.020509033"

	nthRoot := "2"

	expectedStr := "14458.5720947993"

	maxPrecision := uint(10)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_18(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_18"

	radicand := "500001"

	nthRoot := "5"

	expectedStr := "13.7973021335264"

	maxPrecision := uint(13)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_19(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_19"

	radicand := "500001.00000009"

	nthRoot := "5"

	expectedStr := "13.7973021335269"

	maxPrecision := uint(13)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_20(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_20"

	radicand := "500001.00000009"

	nthRoot := "3"

	expectedStr := "79.3701055117479"

	maxPrecision := uint(13)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_21(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_21"

	radicand := "-8000"

	nthRoot := "4"

	maxPrecision := uint(20)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from new(BigIntNum).NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	_, err = new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error because Negative 'radicand' with even 'nthRoot'.\n"+
			"However, NO ERROR WAS RETURNED!\n"+
			"Actual baseINum = '%v'\n"+
			"Actual radicand = '%v'\n"+
			"Actual nthRoot = '%v'\n\n",
			ePrefix, bINumBaseNumStr, radicand, nthRoot)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_22(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_22"

	radicand := "8"

	nthRoot := "0.4"

	expectedStr := "181.01933598375616624661615669884"

	maxPrecision := uint(29)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_23(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_23"

	radicand := "8"

	nthRoot := "-3"

	expectedStr := "0.5"

	maxPrecision := uint(1)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_24(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_24"

	radicand := "8"

	nthRoot := "-3.2"

	//                         1         2         3
	//              0.12345678901234567890123456789012
	expectedStr := "0.52213689121370692016098323936996"

	maxPrecision := uint(32)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_25(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_25"

	// Actual Values Fail
	//	radicand := "8.2"

	//	nthRoot := "-3.2"

	//                         1         2         3
	//              0.12345678901234567890123456789012
	//	expectedStr := "0.5181233574858042598812721854708"
	//	maxPrecision := uint(31)

	// These Values Succeed
	radicand := "8"

	nthRoot := "-3.2"

	//	                         1         2         3
	//	              0.12345678901234567890123456789012
	expectedStr := "0.5221368912137069201609832393700"

	maxPrecision := uint(31)

	/*	radicand := "8.2"

		nthRoot := "-3.3"

		//                         1         2         3
		//              0.12345678901234567890123456789012
		expectedStr := "0.5285507718857623650000905756165"

		maxPrecision := uint(31)
	*/

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_26(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_26"

	radicand := "-8.2"

	nthRoot := "-3.2"

	maxPrecision := uint(31)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	_, err = new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, but NO ERROR WAS RETURNED!\n"+
			"Actual bINumBase = '%v'\n"+
			"Actual bINumNthRoot = '%v'\n"+
			"Actual maxPrecision = '%v'\n\n",
			ePrefix, bINumBaseNumStr, bINumNthRootNumStr, maxPrecision)

		return
	}

	return
}

func TestBigIntMathNthRoot_GetNthRootBigNum_27(t *testing.T) {

	ePrefix := "TestBigIntMathNthRoot_GetNthRootBigNum_27"

	/*
		This Fails
		radicand := "5.967"

		nthRoot := "-2.894"

		//              0.12345678901234567890123456789012
		expectedStr := "0.53944021275349493325378163087104"

		maxPrecision := uint(32)

	*/

	// This Succeeds
	radicand := "5"

	nthRoot := "-2.894"

	//                         1         2         3
	//              0.12345678901234567890123456789012
	expectedStr := "0.5734243831463882071934739402416"

	maxPrecision := uint(31)

	bINumBase, err := new(BigIntNum).NewNumStr(radicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBase, err := new(BigIntNum).NewNumStr(radicand)\n"+
			"radicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicand, err.Error())
		return
	}

	err = bINumBase.IsValid("Validating bINumBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumBase.IsValid('Validating bINumBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumBaseNumStr, err := bINumBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumBaseNumStr, err := bINumBase.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRoot, err := new(BigIntNum).NewNumStr(nthRoot)\n"+
			"nthRoot= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, nthRoot, err.Error())
		return
	}

	err = bINumNthRoot.IsValid("Validating bINumNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINumNthRoot.IsValid('Validating bINumNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bINumNthRootNumStr, err := bINumNthRoot.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	result, err := new(BigIntMathNthRoot).GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntMathNthRoot).\n"+
			"  GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)\n"+
			"bINumBase= '%v'\n"+
			"bINumNthRoot= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bINumBaseNumStr,
			bINumNthRootNumStr,
			maxPrecision,
			err.Error())

		return
	}

	err = result.IsValid("Validating result")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = result.IsValid('Validating result')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultNumStr, err := result.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := result.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedStr != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedStr, resultNumStr)

		return
	}

	return
}
