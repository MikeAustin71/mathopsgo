package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_GetSignedBigInt_01(t *testing.T) {

	ePrefix := "TestNumStrDto_GetSignedBigInt_01"

	inputNumberStr := "-123.456"

	originalBigIntNumStr := "-123456"

	expectedBigIntNum, isOk := big.NewInt(0).SetString(originalBigIntNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"isOk='false' returned by:\n"+
			"expectedBigIntNum, isOk :=\n"+
			"  big.NewInt(0).SetString(originalBigIntNumStr, 10)\n"+
			"originalBigIntNumStr= '%v'\n\n",
			ePrefix, originalBigIntNumStr)
		return
	}

	numStrDtoResult, err := new(NumStrDto).NewPtr().ParseNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err :=\n"+
			"  new(NumStrDto).NewPtr().ParseNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
			"numStrDtoResult set to final value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultBigIntNum, err := numStrDtoResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultBigIntNum, err := numStrDtoResult.GetBigInt()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedBigIntNum.Cmp(numStrDtoResultBigIntNum) != 0 {
		t.Errorf("%v\n"+
			"Error: expectedBigIntNum And numStrDtoResultBigIntNum ARE NOT EQUAL!\n"+
			"Because expectedBigIntNum.Cmp(numStrDtoResultBigIntNum) != 0\n"+
			"Expected numStrDtoResultBigIntNum = '%v'\n"+
			"  Actual numStrDtoResultBigIntNum = '%v'\n\n",
			ePrefix, expectedBigIntNum.Text(10), numStrDtoResultBigIntNum.Text(10))

		return
	}

	return
}

func TestNumStrDto_GetThouStr_01(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_01"

	inputNumberStr := "123456.97"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "123456.97"

	expectedThousandsStr := "123,456.97"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_02(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_02"

	inputNumberStr := "123.45"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "123.45"

	expectedThousandsStr := "123.45"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_03(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_03"

	inputNumberStr := "12345.29"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345.29"

	expectedThousandsStr := "12,345.29"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_04(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_04"

	inputNumberStr := "12345"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345"

	expectedThousandsStr := "12,345"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_05(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_05"

	inputNumberStr := "12345.1234"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345.1234"

	expectedThousandsStr := "12,345.1234"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_06(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_06"

	inputNumberStr := "-12345.29"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-12345.29"

	expectedThousandsStr := "-12,345.29"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_07(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_07"

	inputNumberStr := "-12345"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-12345"

	expectedThousandsStr := "-12,345"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_08(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_08"

	inputNumberStr := "-123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-123"

	expectedThousandsStr := "-123"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_09(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_09"

	inputNumberStr := "-0.123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-0.123"

	expectedThousandsStr := "-0.123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_10"

	inputNumberStr := "-1234567890.123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-1234567890.123"

	expectedThousandsStr := "-1,234,567,890.123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_11(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_11"

	inputNumberStr := "-1234567890.123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-1234567890.123"

	expectedThousandsStr := "-1,234,567,890.123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_12(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_12"

	inputNumberStr := "-1234567890.123"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "-1234567890,123"

	expectedThousandsStr := "-1 234 567 890,123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_13(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_13"

	inputNumberStr := "1234567890.123"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "1234567890,123"

	expectedThousandsStr := "1 234 567 890,123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_14(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_14"

	inputNumberStr := "1234567890"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "1234567890"

	expectedThousandsStr := "1 234 567 890"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouStr_15(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouStr_15"

	inputNumberStr := "-1234567890"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "-1234567890"

	expectedThousandsStr := "-1 234 567 890"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
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

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsNumStr := numStrDtoResult.GetThouStr()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandsStr != numStrDtoResultThousandsNumStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Strings ARE NOT EQUAL!"+
			"Because expectedThousandsStr != numStrDtoResultThousandsNumStr\n"+
			"Expected numStrDtoResultThousandsNumStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsNumStr = '%v'\n\n",
			ePrefix, expectedThousandsStr, numStrDtoResultThousandsNumStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_01(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_01"

	inputNumberStr := "123456.97"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "123456.97"

	expectedThousandParenStr := "123,456.97"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_02(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_02"

	inputNumberStr := "123.45"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "123.45"

	expectedThousandParenStr := "123.45"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_03(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_03"

	inputNumberStr := "12345.29"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345.29"

	expectedThousandParenStr := "12,345.29"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_04(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_04"

	inputNumberStr := "12345"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345"

	expectedThousandParenStr := "12,345"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_05(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_05"

	inputNumberStr := "12345.1234"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "12345.1234"

	expectedThousandParenStr := "12,345.1234"

	expectedPrecisionInt := 4

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_06(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_06"

	inputNumberStr := "1234567890.25"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "1234567890.25"

	expectedThousandParenStr := "1,234,567,890.25"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_07(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_07"

	inputNumberStr := "-12345.29"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-12345.29"

	expectedThousandParenStr := "(12,345.29)"

	expectedPrecisionInt := 2

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_08(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_08"

	inputNumberStr := "-12345"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-12345"

	expectedThousandParenStr := "(12,345)"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_09(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_09"

	inputNumberStr := "-123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-123"

	expectedThousandParenStr := "(123)"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_10(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_10"

	inputNumberStr := "-0.123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-0.123"

	expectedThousandParenStr := "(0.123)"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_11(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_11"

	inputNumberStr := "-1234567890.123"

	expectedCurrencySymbol := '$'

	expectedDecimalSeparator := '.'

	expectedThousandsSeparator := ','

	expectedNumberStr := "-1234567890.123"

	expectedThousandParenStr := "(1,234,567,890.123)"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).NewNumStr(inputNumberStr)\n"+
			"inputNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, err.Error())
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

	err = numStrDtoResult.SetNumericSeparators(expectedDecimalSeparator, expectedThousandsSeparator, expectedCurrencySymbol)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.SetNumericSeparators(\n"+
			"  decimalSeparator, thousandsSeparator, currencySymbol)\n"+
			"decimalSeparator= '%v'\n"+
			"thousandsSeparator= '%v'\n"+
			"currencySymbol= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			string(expectedDecimalSeparator),
			string(expectedThousandsSeparator),
			string(expectedCurrencySymbol),
			err.Error())

		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_12(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_12"

	inputNumberStr := "-1234567890.123"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "-1234567890,123"

	expectedThousandParenStr := "(1 234 567 890,123)"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_13(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_13"

	inputNumberStr := "1234567890.123"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "1234567890,123"

	expectedThousandParenStr := "1 234 567 890,123"

	expectedPrecisionInt := 3

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_14(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_14"

	inputNumberStr := "1234567890"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "1234567890"

	expectedThousandParenStr := "1 234 567 890"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}

func TestNumStrDto_GetThouParen_15(t *testing.T) {

	ePrefix := "TestNumStrDto_GetThouParen_14"

	inputNumberStr := "-1234567890"

	//frenchDecSeparator := ','
	//frenchThousandsSeparator := ' '
	// '\U000020ac'
	//frenchCurrencySymbol := '€'

	expectedCurrencySymbol := '\U000020ac'

	expectedDecimalSeparator := ','

	expectedThousandsSeparator := ' '

	expectedNumberStr := "-1234567890"

	expectedThousandParenStr := "(1 234 567 890)"

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignValue := -1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = expectedDecimalSeparator

	expectedNumSeps.ThousandsSeparator = expectedThousandsSeparator

	expectedNumSeps.CurrencySymbol = expectedCurrencySymbol

	numStrDtoResult, err := new(NumStrDto).NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResult, err := new(NumStrDto).\n"+
			"  NewNumStrWithNumSeps(inputNumberStr, &expectedNumSeps)n"+
			"inputNumberStr= '%v'\n"+
			"expectedNumSeps= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, inputNumberStr, expectedNumSeps.String(), err.Error())
		return
	}

	err = numStrDtoResult.IsValid("Validating final numStrDtoResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = numStrDtoResult.IsValid('Validating final numStrDtoResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
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
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultSignValue, err := numStrDtoResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultSignValue, err := numStrDtoResult.GetSign()\n"+
			"numStrDtoResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	numStrDtoResultCurrencySymbol := numStrDtoResult.GetCurrencySymbol()

	numStrDtoResultDecimalSeparator := numStrDtoResult.GetDecimalSeparator()

	numStrDtoResultThousandsSeparator := numStrDtoResult.GetThousandsSeparator()

	numStrDtoResultThousandsParenStr := numStrDtoResult.GetThouParen()

	numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"numStrDtoResultNumSeps, err := numStrDtoResult.GetNumericSeparatorsDto()\n"+
			"numStrDtoResult= '%v\n"+
			"Error= '%v'\n\n", ePrefix, numStrDtoResultNumberStr, err.Error())
		return
	}

	if expectedCurrencySymbol != numStrDtoResultCurrencySymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency Symbols ARE NOT EQUAL!\n"+
			"Because expectedCurrencySymbol != numStrDtoResultCurrencySymbol\n"+
			"Expected numStrDtoResultCurrencySymbol = '%v'\n"+
			"  Actual numStrDtoResultCurrencySymbol = '%v'\n\n",
			ePrefix, expectedCurrencySymbol, numStrDtoResultCurrencySymbol)

		return
	}

	if expectedDecimalSeparator != numStrDtoResultDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != numStrDtoResultDecimalSeparator\n"+
			"Expected numStrDtoResultDecimalSeparator = '%v'\n"+
			"  Actual numStrDtoResultDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, numStrDtoResultDecimalSeparator)

		return
	}

	if expectedThousandsSeparator != numStrDtoResultThousandsSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Separators ARE NOT EQUAL!\n"+
			"Because expectedThousandsSeparator != numStrDtoResultThousandsSeparator\n"+
			"Expected numStrDtoResultThousandsSeparator = '%v'\n"+
			"  Actual numStrDtoResultThousandsSeparator = '%v'\n\n",
			ePrefix, expectedThousandsSeparator, numStrDtoResultThousandsSeparator)

		return
	}

	if expectedThousandParenStr != numStrDtoResultThousandsParenStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Thousands Paren Strings ARE NOT EQUAL!"+
			"Because expectedThousandParenStr != numStrDtoResultThousandsParenStr\n"+
			"Expected numStrDtoResultThousandsParenStr = '%v'\n"+
			"  Actual numStrDtoResultThousandsParenStr = '%v'\n\n",
			ePrefix, expectedThousandParenStr, numStrDtoResultThousandsParenStr)

		return
	}

	if expectedNumberStr != numStrDtoResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual number strings DO NOT MATCH!\n"+
			"Because expectedNumberStr != numStrDtoResultNumberStr\n"+
			"Expected numStrDtoResultNumberStr = '%v'\n"+
			"  Actual numStrDtoResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, numStrDtoResultNumberStr)

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

	return
}
