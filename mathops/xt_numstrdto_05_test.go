package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_IsNumStrDtoValid_01(t *testing.T) {

	ePrefix := "TestNumStrDto_IsNumStrDtoValid_01"

	nDto := NumStrDto{}

	absAllNumRunes, err := nDto.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"absAllNumRunes, err := nDto.GetAbsAllNumRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenAbsAllNumRunes := len(absAllNumRunes)

	if lenAbsAllNumRunes != 0 {
		t.Errorf("%v\n"+
			"Error: Invalid Result!\n"+
			"Because lenAbsAllNumRunes != 0\n"+
			"Expected lenAbsAllNumRunes = 0'\n"+
			"  Actual lenAbsAllNumRunes = '%v'\n\n",
			ePrefix, lenAbsAllNumRunes)

		return
	}

	err = nDto.IsValid("Validating nDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = nDto.IsValid('Validating nDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	absAllNumRunes, err = nDto.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"absAllNumRunes, err = nDto.GetAbsAllNumRunes()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	lenAbsAllNumRunes = len(absAllNumRunes)

	if lenAbsAllNumRunes != 0 {
		t.Errorf("%v\n"+
			"lenAbsAllNumRunes Error #2!\n"+
			"Because lenAbsAllNumRunes != 0\n"+
			"Expected lenAbsAllNumRunes = '0'\n"+
			"  Actual fixedDecNumStr2 = '%v'\n\n",
			ePrefix, lenAbsAllNumRunes)

		return
	}

	return
}

func TestNumStrDto_Multiply_01(t *testing.T) {

	ePrefix := "TestNumStrDto_Multiply_01"

	multiplierNumStr := "35.123456"

	multiplicandNumStr := "47.9876514"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "1685.4921624912384"

	expectedPrecisionInt := 13

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(10000000000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "1685"

	expectedProductAbsFracStr := "4921624912384"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating initial numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating initial numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	err = numStrDtoMultiplier.Multiply(numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.Multiply(numStrDtoMultiplicand)\n"+
			"multiplierStr= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err = numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResult, err := numStrDtoMultiplier.CopyOut()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrDtoMultiplier.CopyOut()\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoMultiplierNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_01(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_01"

	multiplierNumStr := "35.123456"

	multiplicandNumStr := "47.9876514"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "1685.4921624912384"

	expectedBigInt := big.NewInt(16854921624912384)

	expectedPrecisionInt := 13

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(10000000000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "1685"

	expectedProductAbsFracStr := "4921624912384"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_02(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_02"

	multiplierNumStr := "35.123456"

	multiplicandNumStr := "-47.9876514"

	//                                       1         2         3
	//                            0.1234567890123456789012345678901234567
	expectedProductNumStr := "-1685.4921624912384"

	expectedBigInt := big.NewInt(-16854921624912384)

	expectedPrecisionInt := 13

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(10000000000000)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "1685"

	expectedProductAbsFracStr := "4921624912384"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_03(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_03"

	multiplierNumStr := "-35.123456"

	multiplicandNumStr := "-47.9876514"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "1685.4921624912384"

	expectedBigInt := big.NewInt(16854921624912384)

	expectedPrecisionInt := 13

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(10000000000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "1685"

	expectedProductAbsFracStr := "4921624912384"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_04(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_04"

	multiplierNumStr := "57"

	multiplicandNumStr := "123"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "7011"

	expectedBigInt := big.NewInt(7011)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "7011"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_05(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_05"

	multiplierNumStr := "57"

	multiplicandNumStr := "-123"

	//                                       1         2         3
	//                            0.1234567890123456789012345678901234567
	expectedProductNumStr := "-7011"

	expectedBigInt := big.NewInt(-7011)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "7011"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_06(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_06"

	multiplierNumStr := "-57"

	multiplicandNumStr := "123"

	//                                       1         2         3
	//                            0.1234567890123456789012345678901234567
	expectedProductNumStr := "-7011"

	expectedBigInt := big.NewInt(-7011)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "7011"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_07(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_07"

	multiplierNumStr := "-57"

	multiplicandNumStr := "-123"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "7011"

	expectedBigInt := big.NewInt(7011)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "7011"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_08(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_08"

	multiplierNumStr := "0"

	multiplicandNumStr := "123"

	//                                   1         2         3
	//                        0.1234567890123456789012345678901234567
	expectedProductNumStr := "0"

	expectedBigInt := big.NewInt(0)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "0"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_09(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_09"

	multiplierNumStr := "57"

	multiplicandNumStr := "0"

	//                                   1         2         3
	//                        0.1234567890123456789012345678901234567
	expectedProductNumStr := "0"

	expectedBigInt := big.NewInt(0)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "0"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_10(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_10"

	multiplierNumStr := "-57"

	multiplicandNumStr := "0"

	//                                   1         2         3
	//                        0.1234567890123456789012345678901234567
	expectedProductNumStr := "0"

	expectedBigInt := big.NewInt(0)

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = false

	expectedProductAbsIntStr := "0"

	expectedProductAbsFracStr := ""

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_11(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_11"

	multiplierNumStr := "57"

	multiplicandNumStr := "0.123"

	//                                   1         2         3
	//                        0.1234567890123456789012345678901234567
	expectedProductNumStr := "7.011"

	expectedBigInt := big.NewInt(7011)

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedScaleFactorBigInt := big.NewInt(1000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "7"

	expectedProductAbsFracStr := "011"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_12(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_12"

	multiplierNumStr := "62.1234567890123"

	multiplicandNumStr := "3.12345678901234"

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedProductNumStr := "194.039932864555212496281111782"

	expectedBigInt := big.NewInt(194039932864555212496281111782)

	expectedPrecisionInt := 27

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(1000000000000000000000000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "194"

	expectedProductAbsFracStr := "039932864555212496281111782"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_13(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_13"

	multiplierNumStr := "-62.1234567890123"

	multiplicandNumStr := "3.12345678901234"

	//                                      1         2         3
	//                           0.1234567890123456789012345678901234567
	expectedProductNumStr := "-194.039932864555212496281111782"

	expectedBigInt := big.NewInt(-194039932864555212496281111782)

	expectedPrecisionInt := 27

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(1000000000000000000000000000)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "194"

	expectedProductAbsFracStr := "039932864555212496281111782"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}

func TestNumStrDto_MultiplyNumStrs_14(t *testing.T) {

	ePrefix := "TestNumStrDto_MultiplyNumStrs_14"

	multiplierNumStr := "-62.1234567890123"

	multiplicandNumStr := "-3.12345678901234"

	//                                     1         2         3
	//                          0.1234567890123456789012345678901234567
	expectedProductNumStr := "194.039932864555212496281111782"

	expectedBigInt := big.NewInt(194039932864555212496281111782)

	expectedPrecisionInt := 27

	expectedPrecisionUint := uint(expectedPrecisionInt)

	//                                                   1         2         3
	//                                         01234567890123456789012345678901234567
	expectedScaleFactorBigInt := big.NewInt(1000000000000000000000000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedProductHasNumericDigits, expectedProductIsFractionalValue bool

	expectedProductHasNumericDigits = true

	expectedProductIsFractionalValue = true

	expectedProductAbsIntStr := "194"

	expectedProductAbsFracStr := "039932864555212496281111782"

	nDto := new(NumStrDto).New()

	numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplier, err := nDto.ParseNumStr(multiplierNumStr)\n"+
			"multiplierNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplierNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplier.IsValid("Validating final numStrDtoMultiplier")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplier.IsValid('Validating final numStrDtoMultiplier')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplierNumStr, err := numStrDtoMultiplier.GetNumStr()\n"+
			"numStrDtoMultiplier set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplierNumStr != numStrDtoMultiplierNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplierNumStr Number String is INVALID!\n"+
			"Because multiplierNumStr != numStrDtoMultiplierNumStr\n"+
			"Expected numStrDtoMultiplierNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplierNumStr = '%v'\n\n",
			ePrefix, multiplierNumStr, numStrDtoMultiplierNumStr)

		return
	}

	numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicand, err := nDto.ParseNumStr(multiplicandNumStr)\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, multiplicandNumStr, err.Error())
		return
	}

	err = numStrDtoMultiplicand.IsValid("Validating numStrDtoMultiplicand")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoMultiplicand.IsValid('Validating numStrDtoMultiplicand')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoMultiplicandNumStr, err := numStrDtoMultiplicand.GetNumStr()\n"+
			"numStrDtoMultiplicand set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if multiplicandNumStr != numStrDtoMultiplicandNumStr {
		t.Errorf("%v\n"+
			"Error: numStrDtoMultiplicandNumStr Number String is INVALID!\n"+
			"Because multiplicandNumStr != numStrDtoMultiplicandNumStr\n"+
			"Expected numStrDtoMultiplicandNumStr = '%v'\n"+
			"  Actual numStrDtoMultiplicandNumStr = '%v'\n\n",
			ePrefix, multiplicandNumStr, numStrDtoMultiplicandNumStr)

		return
	}

	numStrDtoResult, err := nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  nDto.MultiplyNumStrs(numStrDtoMultiplier, numStrDtoMultiplicand)\n"+
			"numStrDtoMultiplier= '%v'\n"+
			"numStrDtoMultiplicand= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			numStrDtoMultiplierNumStr,
			numStrDtoMultiplicandNumStr,
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultPrecisionInt := numStrDtoResult.GetPrecision()

	numStrDtoResultPrecisionUint, err := numStrDtoResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultPrecisionUint, err :=\n"+
			"  numStrDtoResult.GetPrecisionUint()\n"+
			"numStrDtoResultResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultScaleFactorBigInt, err := numStrDtoResult.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultScaleFactorBigInt, err :=\n"+
			"  numStrDtoResult.GetScaleFactor()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	numStrDtoResultIsFractionalValue := numStrDtoResult.IsFractionalValue()

	numStrDtoResultAbsIntRunes, err := numStrDtoResult.GetAbsIntRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsIntRunes, err := \n"+
			"  numStrDtoResult.GetAbsIntRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsIntStr := string(numStrDtoResultAbsIntRunes)

	numStrDtoResultAbsFracRunes, err := numStrDtoResult.GetAbsFracRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsFracRunes, err :=\n"+
			"  numStrDtoResult.GetAbsFracRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsFracStr := string(numStrDtoResultAbsFracRunes)

	numStrDtoBigInt, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoBigInt, err :=\n"+
			"  numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedProductNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedProductNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedProductNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedPrecisionInt != numStrDtoResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != numStrDtoResultPrecisionInt\n"+
			"Expected numStrDtoResultPrecisionInt = '%v'\n"+
			"  Actual numStrDtoResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, numStrDtoResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != numStrDtoResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & numStrDtoResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != numStrDtoResultPrecisionUint\n"+
			"Expected numStrDtoResultPrecisionUint = '%v'\n"+
			"  Actual numStrDtoResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, numStrDtoResultPrecisionUint)

		return
	}

	if expectedSignValue != numStrDtoResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & numStrDtoResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != numStrDtoResultSignValue\n"+
			"Expected numStrDtoResultSignValue = '%v'\n"+
			"  Actual numStrDtoResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, numStrDtoResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(numStrDtoResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != numStrDtoResultNumSeps \n"+
			"Expected numStrDtoResultNumSeps = '%v'\n"+
			"  Actual numStrDtoResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), numStrDtoResultNumSeps.String())

		return
	}

	if expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedProductHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedProductHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedProductIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedProductIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedProductAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedProductAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedProductAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedProductAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedProductAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedScaleFactorBigInt.Cmp(numStrDtoResultScaleFactorBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Scale Factor INVALID!\n"+
			"Because expectedScaleFactorBigInt!=numStrDtoResultScaleFactorBigInt\n"+
			"Expected numStrDtoResultScaleFactorBigInt = '%v'\n"+
			"  Actual numStrDtoResultScaleFactorBigInt = '%v'\n\n",
			ePrefix,
			expectedScaleFactorBigInt.Text(10),
			numStrDtoResultScaleFactorBigInt.Text(10))

		return
	}

	if expectedBigInt.Cmp(numStrDtoBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: BigInt Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt.Cmp(numStrDtoBigInt) != 0\n"+
			"Expected numStrDtoBigInt = '%v'\n"+
			"  Actual numStrDtoBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), numStrDtoBigInt.Text(10))

		return
	}

	return
}
