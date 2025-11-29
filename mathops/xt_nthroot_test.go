package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

/*
	These tests are designed to test library methods found in
	source file nthroot.go .

	This test file is located in source code repository:

       https://github.com/MikeAustin71/mathhlpr.git

*/

func TestNthRootOp_GetNthRootFloat32_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootFloat32_01"

	originalFloat32 := float32(125.0)

	originalFloat32PrecisionInt := 1

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nRt := new(NthRootOp)

	intAryResult, err := nRt.GetNthRootFloat32(originalFloat32, 0, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := nRt.GetNthRootFloat32(\n"+
			"  originalFloat32, 0, nthRootInt, maxPrecisionInt)\n"+
			"originalFloat32= '%v'\n"+
			"nthRootInt= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			strconv.FormatFloat(float64(originalFloat32), 'f', originalFloat32PrecisionInt, 32),
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootFloat64_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootFloat64_01"

	originalFloat64 := 125.0

	originalFloat64PrecisionInt := 1

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	nRt := new(NthRootOp)

	intAryResult, err := nRt.GetNthRootFloat64(originalFloat64, 0, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := nRt.GetNthRootFloat64(\n"+
			"  originalFloat64, 0, nthRootInt, maxPrecisionInt)\n"+
			"originalFloat64= '%v'\n"+
			"nthRootInt= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			strconv.FormatFloat(originalFloat64, 'f', originalFloat64PrecisionInt, 64),
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootBigFloat_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootBigFloat_01"

	originalBigFloat := big.NewFloat(125.0)

	originalBigFloatPrecisionInt := 1

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryResult, err := new(NthRootOp).GetNthRootBigFloat(originalBigFloat, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(NthRootOp).GetNthRootBigFloat(\n"+
			"  originalBigFloat, nthRootInt, maxPrecisionInt)\n"+
			"originalBigFloat= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBigFloat.Text('f', originalBigFloatPrecisionInt),
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootInt32_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootInt32_01"

	originalNumInt32 := 125

	originalNumPrecisionInt := 0

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryResult, err := new(NthRootOp).GetNthRootInt(originalNumInt32, originalNumPrecisionInt, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(NthRootOp).GetNthRootInt(\n"+
			"  originalNumInt32, originalNumPrecisionInt,\n"+
			"  nthRootInt, maxPrecisionInt)\n"+
			"originalNumInt32= '%v'\n"+
			"originalNumPrecisionInt= '%v'\n"+
			"nthRootInt= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt32,
			originalNumPrecisionInt,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootInt64_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootInt64_01"

	originalNumInt64 := int64(125)

	originalNumPrecisionInt := 0

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryResult, err := new(NthRootOp).GetNthRootInt64(originalNumInt64, originalNumPrecisionInt, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(NthRootOp).GetNthRootInt64(\n"+
			"  originalNumInt64, originalNumPrecisionInt,\n"+
			"  nthRootInt, maxPrecisionInt)\n"+
			"originalNumInt64= '%v'\n"+
			"originalNumPrecisionInt= '%v'\n"+
			"nthRootInt= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumInt64,
			originalNumPrecisionInt,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootBigInt_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootBigInt_01"

	originalNumBigInt := big.NewInt(125)

	originalNumPrecisionInt := 0

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryResult, err := new(NthRootOp).GetNthRootBigInt(originalNumBigInt, originalNumPrecisionInt, nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := new(NthRootOp).GetNthRootBigInt(\n"+
			"  originalNumBigInt, originalNumPrecisionInt,\n"+
			"  nthRootInt, maxPrecisionInt)\n"+
			"originalNumBigInt= '%v'\n"+
			"originalNumPrecisionInt= '%v'\n"+
			"nthRootInt= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalNumBigInt,
			originalNumPrecisionInt,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootIntAry_01(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootIntAry_01"

	originalBaseNumberStr := "125"

	nthRootOriginalInt := 5

	nthRootOriginalPrecisionUint := uint(0)

	nthRootOriginalNumberStr := "5"

	calcMaxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionInt := 14

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalBaseNumberStr, err.Error())
		return
	}

	err = intAryBase.IsValid("Validating intAryBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryBase.IsValid('Validating intAryBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryBaseNumberStr, err := intAryBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
			"intAryBase set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalBaseNumberStr != intAryBaseNumberStr {
		t.Errorf("%v\n"+
			"Error: intAryBase Number String is INVALID!\n"+
			"Because originalBaseNumberStr != intAryBaseNumberStr\n"+
			"Expected intAryBaseNumberStr = '%v'\n"+
			"  Actual intAryBaseNumberStr = '%v'\n\n",
			ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

		return
	}

	intAryNthRoot, err := new(IntAry).NewInt(nthRootOriginalInt, nthRootOriginalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNthRoot, err := new(IntAry).NewInt(\n"+
			"  nthRootOriginalInt, nthRootOriginalPrecisionUint)\n"+
			"nthRootOriginalInt= '%v'\n"+
			"nthRootOriginalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nthRootOriginalInt,
			nthRootOriginalPrecisionUint,
			err.Error())

		return
	}

	err = intAryNthRoot.IsValid("Validating intAryNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryNthRoot.IsValid('Validating intAryNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNthRootNumberStr, err := intAryNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNthRootNumberStr, err := intAryNthRoot.GetNumStr()\n"+
			"intAryNthRoot set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if nthRootOriginalNumberStr != intAryNthRootNumberStr {
		t.Errorf("%v\n"+
			"Error: intAryNthRoot Number String is INVALID!\n"+
			"Because nthRootOriginalNumberStr != intAryNthRootNumberStr\n"+
			"Expected intAryNthRootNumberStr = '%v'\n"+
			"  Actual intAryNthRootNumberStr = '%v'\n\n",
			ePrefix, nthRootOriginalNumberStr, intAryNthRootNumberStr)

		return
	}

	intAryResult, err := new(NthRootOp).GetNthRootIntAry(&intAryBase, &intAryNthRoot, calcMaxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := nRt.GetNthRootIntAry(\n"+
			"  &intAryBase, &intAryNthRoot, calcMaxPrecisionInt)\n"+
			"intAryBase= '%v'\n"+
			"intAryNthRoot= '%v'\n"+
			"calcMaxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryBaseNumberStr,
			intAryNthRootNumberStr,
			calcMaxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootIntAry_02(t *testing.T) {

	ePrefix := "TestNthRootOp_GetNthRootIntAry_01"

	originalBaseNumberStr := "5604423"

	nthRootOriginalInt := 6

	nthRootOriginalPrecisionUint := uint(0)

	nthRootOriginalNumberStr := "6"

	calcMaxPrecisionInt := 13

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "13.3276982415963"

	expectedPrecisionInt := 13

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBase, err := new(IntAry).NewNumStr(originalBaseNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalBaseNumberStr, err.Error())
		return
	}

	err = intAryBase.IsValid("Validating intAryBase")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryBase.IsValid('Validating intAryBase')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryBaseNumberStr, err := intAryBase.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBaseNumberStr, err := intAryBase.GetNumStr()\n"+
			"intAryBase set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalBaseNumberStr != intAryBaseNumberStr {
		t.Errorf("%v\n"+
			"Error: intAryBase Number String is INVALID!\n"+
			"Because originalBaseNumberStr != intAryBaseNumberStr\n"+
			"Expected intAryBaseNumberStr = '%v'\n"+
			"  Actual intAryBaseNumberStr = '%v'\n\n",
			ePrefix, originalBaseNumberStr, intAryBaseNumberStr)

		return
	}

	intAryNthRoot, err := new(IntAry).NewInt(nthRootOriginalInt, nthRootOriginalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNthRoot, err := new(IntAry).NewInt(\n"+
			"  nthRootOriginalInt, nthRootOriginalPrecisionUint)\n"+
			"nthRootOriginalInt= '%v'\n"+
			"nthRootOriginalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nthRootOriginalInt,
			nthRootOriginalPrecisionUint,
			err.Error())

		return
	}

	err = intAryNthRoot.IsValid("Validating intAryNthRoot")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryNthRoot.IsValid('Validating intAryNthRoot')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNthRootNumberStr, err := intAryNthRoot.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNthRootNumberStr, err := intAryNthRoot.GetNumStr()\n"+
			"intAryNthRoot set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if nthRootOriginalNumberStr != intAryNthRootNumberStr {
		t.Errorf("%v\n"+
			"Error: intAryNthRoot Number String is INVALID!\n"+
			"Because nthRootOriginalNumberStr != intAryNthRootNumberStr\n"+
			"Expected intAryNthRootNumberStr = '%v'\n"+
			"  Actual intAryNthRootNumberStr = '%v'\n\n",
			ePrefix, nthRootOriginalNumberStr, intAryNthRootNumberStr)

		return
	}

	intAryResult, err := new(NthRootOp).GetNthRootIntAry(&intAryBase, &intAryNthRoot, calcMaxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := nRt.GetNthRootIntAry(\n"+
			"  &intAryBase, &intAryNthRoot, calcMaxPrecisionInt)\n"+
			"intAryBase= '%v'\n"+
			"intAryNthRoot= '%v'\n"+
			"calcMaxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryBaseNumberStr,
			intAryNthRootNumberStr,
			calcMaxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
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
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultSignValue, err := intAryResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultSignValue, err := intAryResult.GetSign()\n"+
			"intAryResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumSeps, err := intAryResult.GetNumericSeparatorsDto()\n"+
			"intAryResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionInt != intAryResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult3 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAryResultPrecisionInt\n"+
			"Expected intAryResultPrecisionInt = '%v'\n"+
			"  Actual intAryResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAryResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & IntAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignVal != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestNthRootOp_GetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"

	nthRoot := IntAry{}.NewInt(6, 0)

	maxPrecision := 13
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_04(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"

	nthRoot := IntAry{}.NewInt(3, 0)

	maxPrecision := 2
	expected := "-3.00"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_05(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"

	nthRoot := IntAry{}.NewInt(4, 0)

	maxPrecision := 2
	originalNum.SetIntAryWithNumStr(numStr1)
	_, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for negative number with even nthRoot. No Error triggered")
	}

}

func TestNthRootOp_GetNthRootIntAry_06(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "-5604423.924"

	nthRoot := IntAry{}.NewInt(5, 0)

	maxPrecision := 13

	expected := "-22.3720713464898"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_07(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "5604423.924"

	nthRoot := IntAry{}.NewZero(0)

	maxPrecision := 1
	expected := "1.0"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_08(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "27"

	nthRoot := IntAry{}.NewOne(0)

	maxPrecision := 2
	originalNum.SetIntAryWithNumStr(numStr1)

	_, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for nthRoot == 1. No Error triggered")
	}

}

func TestNthRootOp_GetNthRootIntAry_09(t *testing.T) {
	nRt := NthRootOp{}
	numStr1 := "0.027"

	nthRoot := IntAry{}.NewInt(3, 0)

	maxPrecision := 6

	expected := "0.300000"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_10(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "0.0005"

	nthRoot := IntAry{}.NewInt(9, 0)

	expected := "0.429752972587713"
	maxPrecision := 15

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_11(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "200000.000005"

	nthRoot := IntAry{}.NewTwo(0)

	expected := "447.213595505548"
	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_12(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "200001.100005"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "447.214825341245"
	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_13(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "2000000.0000005"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "1414.21356237327"
	maxPrecision := 11

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
			expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_14(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "20000000.00000005"

	nthRoot := IntAry{}.NewInt(3, 0)

	expected := "271.441761659491"

	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_15(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "20000200.00020005"

	nthRoot := IntAry{}.NewInt(3, 0)

	expected := "271.442666463252"

	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_16(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "2020020.1010205"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "1421.27411185193"
	maxPrecision := 11

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_17(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "209050307.020509033"

	nthRoot := IntAry{}.NewTwo(0)

	expected := "14458.5720947993"

	maxPrecision := 10

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
			expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_18(t *testing.T) {

	nRt := NthRootOp{}
	radicandStr := "500001"

	nthRoot := IntAry{}.NewInt(5, 0)

	expected := "13.7973021335264"

	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
			expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_19(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "500001.00000009"

	nthRoot := IntAry{}.NewInt(5, 0)

	expected := "13.7973021335269"
	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_20(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "500001.00000009"

	nthRoot := IntAry{}.NewInt(3, 0)

	expected := "79.3701055117479"
	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_21(t *testing.T) {

	radicandStr := "-8000"
	maxPrecision := 20
	nthRoot := IntAry{}.NewInt(4, 0)

	origRadicand, _ := IntAry{}.NewNumStr(radicandStr)

	nthRt := NthRootOp{}
	_, err := nthRt.GetNthRootIntAry(&origRadicand, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected an Error. Negative OriginalRadicand with even NthRootInt. " +
			"Instead, NO ERROR WAS RETURNED.")
	}

}

func TestNthRootOp_GetNthRootIntAry_22(t *testing.T) {
	radicandStr := "8"
	nthRootStr := "0.4"
	expected := "181.01933598375616624661615669884"
	maxPrecision := 29

	nthRoot, err := IntAry{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRootStr, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicandStr)

	nthRt := NthRootOp{}

	ai, err := nthRt.GetNthRootIntAry(&origRadicand, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.GetNthRootIntAry(...) - %v",
			err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_23(t *testing.T) {
	radicand := "8"
	nthRoot := "-3"
	expected := "0.5"
	maxPrecision := 1

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
			err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_24(t *testing.T) {
	radicand := "8"
	nthRoot := "-3.2"
	expectedStr := "0.52213689121370692016098323936996"
	maxPrecision := 32

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
			err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_25(t *testing.T) {

	radicand := "8.2"
	nthRoot := "-3.2"
	//              0.12345678901234567890123456789012
	expectedStr := "0.5181233574858042598812721854708"
	maxPrecision := 31

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
			err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_26(t *testing.T) {

	radicand := "-8.2"
	nthRoot := "-3.2"
	maxPrecision := 31

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	_, err = NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected a valid error object to be returned. Error: err==nil!")
	}

}

func TestNthRootOp_GetNthRootIntAry_27(t *testing.T) {

	radicand := "5.967"
	nthRoot := "-2.894"
	//              0.12345678901234567890123456789012
	expectedStr := "0.53944021275349493325378163087104"
	maxPrecision := 32

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
			err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootFloat32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float32(2686.5)

	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootFloat64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float64(2686.5)

	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigFloat_01(t *testing.T) {

	nRt := NthRootOp{}
	num, ok := big.NewFloat(0).SetString("2686.5")

	if !ok {
		t.Error("Conversion failed big.NewFloat(0).SetString(\"2686.5\")")
	}

	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigFloat(num, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigFloat() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int32(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int64(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := big.NewInt(int64(26865))
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "2686.5"
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetSquareRootIntAry(&originalNum, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "125"

	nthRoot := IntAry{}.NewInt(5, 0)

	maxPrecision := 14
	expected := "2.62652780440377"
	originalNum.SetIntAryWithNumStr(numStr1)

	err := nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_02(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423"

	nthRoot := IntAry{}.NewInt(6, 0)

	maxPrecision := 13
	expected := "13.3276982415963"
	originalNum.SetIntAryWithNumStr(numStr1)

	err := nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"

	nthRoot := IntAry{}.NewInt(6, 0)

	maxPrecision := 13
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	err := nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}
