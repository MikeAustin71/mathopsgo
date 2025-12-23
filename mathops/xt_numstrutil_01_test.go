package mathops

import (
	"fmt"
	"math/big"

	"testing"
)

/*
	These tests are associated with numstrutility.go.

	The source code repository for numstrutility.go is located at :
					https://github.com/MikeAustin71/numstrutility.git

*/

func TestNumStrUtility_DlimDecStr(t *testing.T) {

	ePrefix := "TestNumStrUtility_DlimDecStr"

	ns := new(NumStrUtility)

	inputNumstr := "1234567890"

	expectedNumStr := "1,234,567,890"

	resultNumStr := ns.DlimDecCurrStr(inputNumstr, ',', '.', '$')

	if resultNumStr != expectedNumStr {
		t.Errorf("%v\n"+
			"Expected resultNumStr = %v;\n"+
			"instead resultNumStr= %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_DlimDecStr_With_DecimalCurrency(t *testing.T) {

	ePrefix := "TestNumStrUtility_DlimDecStr_With_DecimalCurrency"

	numStrUtil := new(NumStrUtility)

	inputNumStr := "$1234567890.25"

	expectedNumStr := "$1,234,567,890.25"

	resultNumStr := numStrUtil.DlimDecCurrStr(inputNumStr, ',', '.', '$')

	if resultNumStr != expectedNumStr {

		t.Errorf("%v\n"+
			"Expected resultNumStr = %v\n"+
			"instead resultNumStr = %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_DnumStr(t *testing.T) {

	ePrefix := "TestNumStrUtility_DnumStr"

	numStrUtility := new(NumStrUtility)

	inputNumStr := "1234567890"

	expectedNumStr := "1,234,567,890"

	resultNumStr := numStrUtility.DnumStr(inputNumStr, ',')

	if resultNumStr != expectedNumStr {
		t.Errorf("%v\n"+
			"Expected resultNumStr = %v"+
			"instead resultNumStr = %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_DNumI64(t *testing.T) {

	ePrefix := "TestNumStrUtility_DNumI64"

	numStrUtility := new(NumStrUtility)

	inputNumInt64 := int64(1234567890)

	expectedNumStr := "1,234,567,890"

	resultNumStr := numStrUtility.DLimI64(inputNumInt64, ',')

	if resultNumStr != expectedNumStr {

		t.Errorf("%v\n"+
			"Expected resultNumStr = %v\n"+
			"instead resultNumStr = %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_DNumI64_EvenThousands(t *testing.T) {
	ePrefix := "TestNumStrUtility_DNumI64_EvenThousands"

	numStrUtility := new(NumStrUtility)

	inputNumInt64 := int64(123456)

	expectedNumStr := "123,456"

	resultNumStr := numStrUtility.DLimI64(inputNumInt64, ',')

	if resultNumStr != expectedNumStr {

		t.Errorf("%v\n"+
			"Expected resultNumStr = %v\n"+
			"instead resultNumStr = %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_DLimInt(t *testing.T) {

	ePrefix := "TestNumStrUtility_DLimInt"

	numStrUtility := new(NumStrUtility)

	inputNumInt := 1234567

	expectedNumStr := "1,234,567"

	resultNumStr := numStrUtility.DLimInt(inputNumInt, ',')

	if resultNumStr != expectedNumStr {

		t.Errorf("%v\n"+
			"Expected resultNumStr = %v\n"+
			"instead resultNumStr = %v",
			ePrefix, expectedNumStr, resultNumStr)
	}

	return
}

func TestNumStrUtility_ConvertStrToInt64_01(t *testing.T) {

	ePrefix := "TestNumStrUtility_ConvertStrToInt64_01"

	expectedNumStr := "-12314617914"

	numStrUtility := new(NumStrUtility)

	initialResultNumStr, err := numStrUtility.ConvertNumStrToInt64(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"initialResultNumStr, err := numStrUtility.\n"+
			"  ConvertNumStrToInt64(inputNumStr)\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())

		return
	}

	finalResultNumStr, err := numStrUtility.ConvertInt64ToStr(initialResultNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"finalResultNumStr, err := numStrUtility.\n"+
			"  ConvertInt64ToStr(initialResultNumStr)\n"+
			"initialResultNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, initialResultNumStr, err.Error())

		return
	}

	if expectedNumStr != finalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Results DON'T MATCH!\n"+
			"Because expectedNumStr != finalResultNumStr\n"+
			"Expected finalResultNumStr = '%v'\n"+
			"  Actual finalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, finalResultNumStr)

		return
	}

	return
}

func TestNumStrUtility_ConvertStrToInt64_02(t *testing.T) {

	ePrefix := "TestNumStrUtility_ConvertStrToInt64_02"

	expectedNumStr := "+12314617914"

	numStrUtility := new(NumStrUtility)

	resultInt64, err := numStrUtility.ConvertNumStrToInt64(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultInt64, err := numStrUtility.\n"+
			"  ConvertNumStrToInt64(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	resultNumStr, err := numStrUtility.ConvertInt64ToStr(resultInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := numStrUtility.\n"+
			"  ConvertInt64ToStr(resultInt64)\n"+
			"resultInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, resultInt64, err.Error())
		return
	}

	resultNumStr = "+" + resultNumStr

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	return
}

func TestNumStrUtility_ConvertStrToInt64_03(t *testing.T) {

	ePrefix := "TestNumStrUtility_ConvertStrToInt64_03"

	expectedNumStr := "12314617914"

	numStrUtility := new(NumStrUtility)

	initialResultInt64, err := numStrUtility.ConvertNumStrToInt64(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"initialResultInt64, err := numStrUtility.\n"+
			"  ConvertNumStrToInt64(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	finalResultNumStr, err := numStrUtility.ConvertInt64ToStr(initialResultInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"finalResultNumStr, err := numStrUtility.\n"+
			"  ConvertInt64ToStr(initialResultInt64)\n"+
			"initialResultInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, initialResultInt64, err.Error())
		return
	}

	if expectedNumStr != finalResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != finalResultNumStr\n"+
			"Expected finalResultNumStr = '%v'\n"+
			"  Actual finalResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, finalResultNumStr)

		return
	}

	return
}

func TestNumStrUtility_ConvertStrToIntNumStr(t *testing.T) {

	ePrefix := "TestNumStrUtility_ConvertStrToIntNumStr"

	inputNumStr := "-12,314,617,914"

	expectedNumStr := "-12314617914"

	numStrUtility := NumStrUtility{}

	numStrUtility.ThousandsSeparator = ','

	resultNumStr, err := numStrUtility.ConvertStrToIntNumStr(inputNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumStr, err := numStrUtility.\n"+
			"  ConvertStrToIntNumStr(inputNumStr)\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr, err.Error())
		return
	}

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	return
}

func TestNumStrUtility_SetCountryAndCurrency(t *testing.T) {

	ePrefix := "TestNumStrUtility_SetCountryAndCurrency"

	numStrUtility := new(NumStrUtility)

	inputCountryStr := "United States"

	expectedCurrencySymbolRune := '$'

	expectedNationStr := inputCountryStr

	expectedThousandsSeparatorRune := ','

	expectedDecimalSeparatorRune := '.'

	err := numStrUtility.SetCountryAndCurrency(inputCountryStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrUtility.SetCountryAndCurrency(\n"+
			"  inputCountryStr)\n"+
			"inputCountryStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputCountryStr, err.Error())
		return
	}

	if expectedCurrencySymbolRune != numStrUtility.CurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Currency Symbols DON'T MATCH!\n"+
			"Because expectedCurrencySymbolRune != numStrUtility.CurrencySymbol\n"+
			"Expected numStrUtility.CurrencySymbol = '%v'\n"+
			"  Actual numStrUtility.CurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbolRune, numStrUtility.CurrencySymbol)

		return
	}

	if expectedNationStr != numStrUtility.Nation {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Nation Strings DON'T MATCH!\n"+
			"Because expectedNationStr != numStrUtility.Nation\n"+
			"Expected numStrUtility.Nation = '%v'\n"+
			"  Actual numStrUtility.Nation = '%v'\n\n",
			ePrefix, expectedNationStr, numStrUtility.Nation)

		return
	}

	if expectedThousandsSeparatorRune != numStrUtility.ThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Thousands Separators DON'T MATCH!\n"+
			"Because expectedThousandsSeparatorRune != numStrUtility.ThousandsSeparator\n"+
			"Expected numStrUtility.ThousandsSeparator = '%v'\n"+
			"  Actual numStrUtility.ThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparatorRune, numStrUtility.ThousandsSeparator)

		return
	}

	if expectedDecimalSeparatorRune != numStrUtility.DecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrUtility.DecimalSeparator = '%v'\n"+
			"  Actual numStrUtility.DecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparatorRune, numStrUtility.DecimalSeparator)

		return
	}

	return
}

func TestNumStrUtility_ConvertInt64ToFloat64Value(t *testing.T) {

	ePrefix := "TestNumStrUtility_ConvertInt64ToFloat64Value"

	expectedNumStr := "0.123456"

	numStrUtility := NumStrUtility{}

	inputInt64 := int64(123456)

	resultFloat64, err := numStrUtility.ConvertInt64ToFloat64Value(inputInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultFloat64, err := numStrUtility.\n"+
			"  ConvertInt64ToFloat64Value(inputInt64)\n"+
			"inputInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputInt64, err.Error())
		return
	}

	resultNumStr := fmt.Sprintf("%v", resultFloat64)

	if expectedNumStr != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != resultNumStr\n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, resultNumStr)

		return
	}

	return
}

func TestNumStrUtility_ParseNumString_01(t *testing.T) {

	ePrefix := "TestNumStrUtility_ParseNumString_01"

	inputNumStr := "123456.654321"

	numStrUtility := NumStrUtility{}

	expectedNumStr := "123456.654321"

	expectedBigInt := big.NewInt(123456654321)

	expectedPrecisionInt := 6

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := "654321"

	expectedAbsAllRunesNumStr := "123456654321"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDtoResult, err := numStrUtility.ParseNumString(inputNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrUtility.\n"+
			" ParseNumString(inputNumStr)\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrUtility_ParseNumString_02(t *testing.T) {

	ePrefix := "TestNumStrUtility_ParseNumString_02"

	inputNumStr := "-123456.654321"

	numStrUtility := new(NumStrUtility)

	expectedNumStr := "-123456.654321"

	expectedBigInt := big.NewInt(-123456654321)

	expectedPrecisionInt := 6

	expectedPrecisionUint := uint(expectedPrecisionInt)

	big10 := big.NewInt(10)

	baseExp := big.NewInt(int64(expectedPrecisionInt))

	expectedScaleFactorBigInt := big.NewInt(0).Exp(big10, baseExp, nil)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	var expectedNumHasNumericDigits, expectedNumIsFractionalValue bool

	expectedNumHasNumericDigits = true

	expectedNumIsFractionalValue = true

	expectedNumAbsIntStr := "123456"

	expectedNumAbsFracStr := "654321"

	expectedAbsAllRunesNumStr := "123456654321"

	expectedNumStrDto, err := new(NumStrDto).ParseNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).\n"+
			"  ParseNumStr(expectedNumStr)\n"+
			"expectedNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDtoNumStr, err := expectedNumStrDto.GetNumStr()\n"+
			"expectedNumStrDto set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumStr != expectedNumStrDtoNumStr {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != expectedNumStrDtoNumStr\n"+
			"Expected expectedNumStrDtoNumStr = '%v'\n"+
			"  Actual expectedNumStrDtoNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, expectedNumStrDtoNumStr)

		return
	}

	numStrDtoResult, err := numStrUtility.ParseNumString(inputNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrUtility.\n"+
			" ParseNumString(inputNumStr)\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
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

	numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultAbsAllRunes, err := numStrDtoResult.GetAbsAllNumRunes()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, err.Error())
		return
	}

	numStrDtoResultAbsAllRunesNumStr := string(numStrDtoResultAbsAllRunes)

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

	expectedAndResultNumStrDtosAreEqual, err := expectedNumStrDto.EqualTo(numStrDtoResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedAndResultNumStrDtosAreEqual, err :=\n"+
			"  expectedNumStrDto.EqualTo(numStrDtoResult)\n"+
			"expectedNumStrDto= '%v'\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumStrDtoNumStr, numStrDtoResultNumStr, err.Error())
		return
	}

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and numStrDtoResult Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if expectedAndResultNumStrDtosAreEqual == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultNumStr and numStrDtoResultNumStr ARE NOT EQUAL!\n"+
			"Because expectedAndResultNumStrDtosAreEqual == false\n"+
			"Expected numStrDtoResult = '%v'\n"+
			"  Actual numStrDtoResult = '%v'\n\n",
			ePrefix, numStrDtoResultNumStr, numStrDtoResultNumStr)

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

	if expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits {
		t.Errorf("%v\n"+
			"Error: Numeric Digits Flag is Invalid!\n"+
			"Because expectedNumHasNumericDigits != numStrDtoResultHasNumericDigits\n"+
			"Expected numStrDtoResultHasNumericDigits = '%v'\n"+
			"  Actual numStrDtoResultHasNumericDigits = '%v'\n\n",
			ePrefix, expectedNumHasNumericDigits, numStrDtoResultHasNumericDigits)

		return
	}

	if expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue {
		t.Errorf("%v\n"+
			"Error: IsFractionalValue Flag Invalid!\n"+
			"Because expectedNumIsFractionalValue != numStrDtoResultIsFractionalValue\n"+
			"Expected numStrDtoResultIsFractionalValue = '%v'\n"+
			"  Actual numStrDtoResultIsFractionalValue = '%v'\n\n",
			ePrefix, expectedNumIsFractionalValue, numStrDtoResultIsFractionalValue)

		return
	}

	if expectedNumAbsIntStr != numStrDtoResultAbsIntStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Integer Strings ARE NOT EQUAL!\n"+
			"Because expectedNumAbsIntStr != numStrDtoResultAbsIntStr\n"+
			"Expected numStrDtoResultAbsIntStr = '%v'\n"+
			"  Actual numStrDtoResultAbsIntStr = '%v'\n\n",
			ePrefix, expectedNumAbsIntStr, numStrDtoResultAbsIntStr)

		return
	}

	if expectedNumAbsFracStr != numStrDtoResultAbsFracStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because!!!!\n"+
			"Expected numStrDtoResultAbsFracStr = '%v'\n"+
			"  Actual numStrDtoResultAbsFracStr = '%v'\n\n",
			ePrefix, expectedNumAbsFracStr, numStrDtoResultAbsFracStr)

		return
	}

	if expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Absolute Runes Num Strings ARE NOT EQUAL!\n"+
			"Because expectedAbsAllRunesNumStr != numStrDtoResultAbsAllRunesNumStr\n"+
			"Expected numStrDtoResultAbsAllRunesNumStr = '%v'\n"+
			"  Actual numStrDtoResultAbsAllRunesNumStr = '%v'\n\n",
			ePrefix, expectedAbsAllRunesNumStr, numStrDtoResultAbsAllRunesNumStr)

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

func TestNumStrUtility_ParseNumString_03(t *testing.T) {

	ePrefix := "TestNumStrUtility_ParseNumString_03"

	inputNumStr := "Nothing"

	//numStrUtility := NumStrUtility{}

	expectedNumStr := "0"

	numStrUtility := new(NumStrUtility)

	numStrDtoResult, err := numStrUtility.ParseNumString(inputNumStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := numStrUtility.\n"+
			"  ParseNumString(inputNumStr)\n\n"+
			"inputNumStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumStr, err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
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

	numStrDtoResultHasNumericDigits := numStrDtoResult.HasNumericDigits()

	if expectedNumStr != numStrDtoResultNumStr {
		t.Errorf("%v\n"+
			"Error: Expected vs Actual Number Strings DON'T MATCH!\n"+
			"Because expectedNumStr != numStrDtoResultNumStr\n"+
			"Expected numStrDtoResultNumStr = '%v'\n"+
			"  Actual numStrDtoResultNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, numStrDtoResultNumStr)

		return
	}

	if numStrDtoResultHasNumericDigits == false {
		t.Errorf("%v\n"+
			"Error: numStrDtoResultHasNumericDigits HAS NO NUMERIC DIGITS!`\n"+
			"Because numStrDtoResultHasNumericDigits == false\n"+
			"Expected numStrDtoResultHasNumericDigits = 'true'\n"+
			"  Actual numStrDtoResultHasNumericDigits = 'false'\n\n",
			ePrefix)

		return
	}

	return
}
