package mathops

import (
	"math/big"
	"strconv"
	"testing"
)

func TestIntAry_GetAbsoluteValue_01(t *testing.T) {

	ePrefix := "TestIntAry_GetAbsoluteValue_01"

	originalNumberStr := "-927.351"

	//                                 1         2         3
	//                      0.123456789012345678901234567890
	expectedNumberStr := "927.351"

	expectedPrecisionUint := uint(3)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1, err := new(IntAry).NewNumStr(\n"+
			"  originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetAbsoluteValue()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetAbsoluteValue()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionUint, err := intAry2.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2PrecisionUint, err :=\n"+
			"  intAry2.GetPrecisionUint()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionUint != intAry2PrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAry2PrecisionUint\n"+
			"Expected intAry2PrecisionUint = '%v'\n"+
			"  Actual intAry2PrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAry2PrecisionUint)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	return
}

func TestIntAry_GetBigIntNum_01(t *testing.T) {

	ePrefix := "TestIntAry_GetBigIntNum_01"

	originalBigInt := big.NewInt(int64(123456123456))

	originalPrecisionUint := uint(6)

	originalPrecisionInt := 6

	//                                    1         2         3
	//                         0.123456789012345678901234567890
	expectedNumberStr := "123456.123456"

	expectedPrecisionUint := uint(6)

	expectedScaleFactor := big.NewInt(1000000)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigInt(originalBigInt, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" expectedBigIntNum, err := new(BigIntNum).NewBigInt(\n"+
			"  originalBigInt, originalPrecisionUint)\n"+
			"originalBigInt= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBigInt.Text(10),
			originalPrecisionUint,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = expectedBigIntNum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {

		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())

		return
	}

	if expectedNumberStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & expected bInt Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != expectedBigIntNumNumberStr\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, expectedBigIntNumNumberStr)

		return
	}

	intAry, err := new(IntAry).NewBigInt(originalBigInt, originalPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewBigInt(\n"+
			"  originalBigInt, originalPrecisionInt)\n"+
			"multiplierStr= '%v'\n"+
			"multiplicandStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalBigInt.Text(10),
			originalPrecisionInt,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigINum, err := intAry.GetBigIntNum()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINum, err := intAry.GetBigIntNum()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = bigINum.IsValid("Validating bigINum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bigINum.IsValid('Validating bigINum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigINumNumberStr, err := bigINum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumNumberStr, err := bigINum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigINumPrecisionUint, err := bigINum.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumPrecisionUint, err := bigINum.GetPrecisionUint()\n"+
			"bigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigINumNumberStr, err.Error())
		return
	}

	bigINumSignValue, err := bigINum.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumSignValue, err := bigINum.GetSign()\n"+
			"bigINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bigINumNumberStr, err.Error())
		return
	}

	bigINumScaleFactor, err := bigINum.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumScaleFactor, err := bigINum.GetScaleFactor()\n"+
			"bigINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bigINumNumberStr, err.Error())
		return
	}

	bigINumBigInt, err := bigINum.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumBigInt, err := bigINum.GetBigInt()\n"+
			"bigINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bigINumNumberStr, err.Error())
		return
	}

	bigINumSeps, err := bigINum.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigINumSeps, err := bigINum.GetNumericSeparatorsDto()\n"+
			"bigINum= '%v\n"+
			"Error= '%v'\n\n", ePrefix, bigINumNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigINumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & bINum Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigINumNumberStr\n"+
			"Expected bigINumNumberStr = '%v'\n"+
			"  Actual bigINumNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigINumNumberStr)

		return
	}

	expectedBigIntNumEqualsBigINum, err := expectedBigIntNum.Equal(bigINum)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumEqualsBigINum, err :=\n"+
			"  expectedBigIntNum.Equal(bigINum)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigINum= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigINumNumberStr,
			err.Error())

		return
	}

	if !expectedBigIntNumEqualsBigINum {
		t.Errorf("%v\n"+
			"Error: expectedBigIntNumEqualsBigINum == 'false'!\n"+
			"Because expectedBigIntNum != bigINum\n"+
			"Expected bigINum = '%v'\n"+
			"  Actual bigINum = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigINumNumberStr)

		return
	}

	if originalBigInt.Cmp(bigINumBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: expected/bigINum Big Int Values ARE NOT EQUAL!\n"+
			"Because originalBigInt.Cmp(bigINumBigInt) != 0\n"+
			"Expected bigINumBigInt = '%v'\n"+
			"  Actual bigINumBigInt = '%v'\n\n",
			ePrefix, originalBigInt.Text(10), bigINumBigInt.Text(10))

		return
	}

	if expectedPrecisionUint != bigINumPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != bigINumPrecisionUint\n"+
			"Expected bigINumPrecisionUint = '%v'\n"+
			"  Actual bigINumPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigINumPrecisionUint)

		return
	}

	if expectedScaleFactor.Cmp(bigINumScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
			"Because expectedScaleFactor.Cmp(bigINumScaleFactor) != 0\n"+
			"Expected bigINumScaleFactor = '%v'\n"+
			"  Actual bigINumScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), bigINumScaleFactor.Text(10))

		return
	}

	if expectedSignValue != bigINumSignValue {
		t.Errorf("%v\n"+
			"Error: expected & bigINum Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != bigINumSignValue\n"+
			"Expected bigINumSignValue = '%v'\n"+
			"  Actual bigINumSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, bigINumSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigINumSeps) {
		t.Errorf("%v\n"+
			"Error: Number Sign Values NOT Equal\n"+
			"Because expectedNumSeps != bigINumSeps \n"+
			"Expected bigINumSeps = '%v'\n"+
			"  Actual bigINumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigINumSeps.String())

		return
	}

	return
}

func TestIntAry_GetCurrencySymbol_01(t *testing.T) {

	ePrefix := "TestIntAry_GetCurrencySymbol_01"

	expectedCurrencyRune := '$'

	intAry := new(IntAry).New()

	curSymbol := intAry.GetCurrencySymbol()

	if expectedCurrencyRune != curSymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency SymbolS ARE NOT EQUAL!\n"+
			"Because expectedCurrencyRune != curSymbol\n"+
			"Expected curSymbol = '%v'\n"+
			"  Actual curSymbol = '%v'\n\n",
			ePrefix, expectedCurrencyRune, curSymbol)

		return
	}

	return
}

func TestIntAry_GetCurrencySymbol_03(t *testing.T) {

	ePrefix := "TestIntAry_GetCurrencySymbol_03"

	expectedCurrencyRune := '$'

	originalNumberStr := "50.37"

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	curSymbol := intAry.GetCurrencySymbol()

	if expectedCurrencyRune != curSymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency SymbolS ARE NOT EQUAL!\n"+
			"Because expectedCurrencyRune != curSymbol\n"+
			"Expected curSymbol = '%v'\n"+
			"  Actual curSymbol = '%v'\n\n",
			ePrefix, expectedCurrencyRune, curSymbol)

		return
	}

	return
}

func TestIntAry_GetCurrencySymbol_04(t *testing.T) {

	ePrefix := "TestIntAry_GetCurrencySymbol_04"

	expectedNumberStr := "50.34"

	expectedCurrencyRune := '$'

	originalInt64 := int64(5034)

	originalPrecisionUint := uint(2)

	intAry, err := new(IntAry).NewInt64(originalInt64, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewInt64(\n"+
			"  originalInt64, originalPrecisionUint)\n"+
			"originalInt64= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalInt64,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	curSymbol := intAry.GetCurrencySymbol()

	if expectedCurrencyRune != curSymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency SymbolS ARE NOT EQUAL!\n"+
			"Because expectedCurrencyRune != curSymbol\n"+
			"Expected curSymbol = '%v'\n"+
			"  Actual curSymbol = '%v'\n\n",
			ePrefix, expectedCurrencyRune, curSymbol)

		return
	}

	return
}

func TestIntAry_GetCurrencySymbol_05(t *testing.T) {

	ePrefix := "TestIntAry_GetCurrencySymbol_04"

	expectedNumberStr := "50.37"

	var expectedCurrencyRune rune

	expectedCurrencyRune = '\U000000a3'

	intAry, err := new(IntAry).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	err = intAry.SetCurrencySymbol(expectedCurrencyRune)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetCurrencySymbol(expectedCurrencyRune)\n"+
			"intAry= '%v'\n"+
			"expectedCurrencyRune= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			expectedCurrencyRune,
			err.Error())

		return
	}

	curSymbol := intAry.GetCurrencySymbol()

	if expectedCurrencyRune != curSymbol {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Currency SymbolS ARE NOT EQUAL!\n"+
			"Because expectedCurrencyRune != curSymbol\n"+
			"Expected curSymbol = '%v'\n"+
			"  Actual curSymbol = '%v'\n\n",
			ePrefix, expectedCurrencyRune, curSymbol)

		return
	}

	return
}

func TestIntAry_GetDecimal_01(t *testing.T) {

	ePrefix := "TestIntAry_GetDecimal_01"

	expectedNumberStr := "198649257.12345678"

	expectedPrecisionUint := uint(8)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	controlDecimal, err := new(Decimal).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"controlDecimal, err := new(Decimal).NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
		return
	}

	err = controlDecimal.IsValid("Validating controlDecimal")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = controlDecimal.IsValid('Validating controlDecimal')\n"+
			"controlDecimal set to expectedNumberStr\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	controlDecimalNumberStr, err := controlDecimal.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"controlDecimalNumberStr, err := controlDecimal.GetNumStr()\n"+
			"controlDecimal set to expectedNumberStr\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != controlDecimalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Control Decimal Number String Values ARE NOT Equal!\n"+
			"Because expectedNumberStr != controlDecimalNumberStr \n"+
			"Expected controlDecimalNumberStr = '%v'\n"+
			"  Actual controlDecimalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, controlDecimalNumberStr)

		return
	}

	intAry, err := new(IntAry).NewNumStr(expectedNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(expectedNumberStr)\n"+
			"expectedNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, expectedNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to expectedNumberStr value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	decActual, err := intAry.GetDecimal()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decActual, err := intAry.GetDecimal()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = decActual.IsValid("Validating decActual-nDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = decActual.IsValid('Validating decActual-nDto')\n"+
			"decActual set to nDto\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decActualNumberStr, err := decActual.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decActualNumberStr, err := decActual.GetNumStr()\n"+
			"decActual set to nDto\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	decActualPrecisionUint, err := decActual.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decActualPrecisionUint, err :=\n"+
			"  decActual.GetPrecisionUint()\n"+
			"decActual= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decActualNumberStr, err.Error())
		return
	}

	decActualSignValue, err := decActual.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decActualSignValue, err := decActual.GetSign()\n"+
			"decActual= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, decActualNumberStr, err.Error())
		return
	}

	decActualNumSeps, err := decActual.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"decActualNumSeps, err := decActual.GetNumericSeparatorsDto()\n"+
			"decActual= '%v\n"+
			"Error= '%v'\n\n", ePrefix, decActualNumberStr, err.Error())
		return
	}

	if expectedNumberStr != decActualNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Decimal Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != decActualNumberStr\n"+
			"Expected decActualNumberStr = '%v'\n"+
			"  Actual decActualNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, decActualNumberStr)

		return
	}

	if expectedPrecisionUint != decActualPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & decActual Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != decActualPrecisionUint\n"+
			"Expected decActualPrecisionUint = '%v'\n"+
			"  Actual decActualPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, decActualPrecisionUint)

		return
	}

	if expectedSignVal != decActualSignValue {
		t.Errorf("%v\n"+
			"Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != decActualSignValue\n"+
			"Expected decActualSignValue = '%v'\n"+
			"  Actual decActualSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, decActualSignValue)

		return
	}

	if !expectedNumSeps.Equal(decActualNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != decActualNumSeps \n"+
			"Expected decActualNumSeps = '%v'\n"+
			"  Actual decActualNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), decActualNumSeps.String())

		return
	}

	controlDecimalEqualsDecimalActual, err := controlDecimal.Equal(decActual)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"controlDecimalEqualsDecimalActual, err := \n"+
			"  controlDecimal.Equal(decActual)\n"+
			"controlDecimal= '%v'\n"+
			"decActual= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			controlDecimalNumberStr,
			decActualNumberStr,
			err.Error())

		return
	}

	if !controlDecimalEqualsDecimalActual {
		t.Errorf("%v\n"+
			"Error: Control Decimal and Actual Decimal Values ARE NOT EQUAL!\n"+
			"Because controlDecimalEqualsDecimalActual == 'false'\n"+
			"Expected controlDecimalEqualsDecimalActual = 'true'\n"+
			"  Actual controlDecimalEqualsDecimalActual = 'false\n"+
			"Expected decActual = '%v'\n"+
			"  Actual decActual = '%v'\n\n",
			ePrefix, controlDecimalNumberStr, decActualNumberStr)

		return
	}

	return
}

func TestIntAry_GetDecimalSeparator_01(t *testing.T) {

	ePrefix := "TestIntAry_GetDecimalSeparator_01"

	expectedDecimalSeparatorRune := '.'

	intAry := new(IntAry).New()

	actualDecimalSeparatorRune := intAry.GetDecimalSeparator()

	if expectedDecimalSeparatorRune != actualDecimalSeparatorRune {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparatorRune != actualDecimalSeparatorRune\n"+
			"Expected actualDecimalSeparatorRune = '%v'\n"+
			"  Actual actualDecimalSeparatorRune = '%v'\n\n",
			ePrefix, expectedDecimalSeparatorRune, actualDecimalSeparatorRune)

		return
	}

	return
}

func TestIntAry_GetDecimalSeparator_02(t *testing.T) {

	ePrefix := "TestIntAry_GetDecimalSeparator_02"

	originalNumberStr := "50.37"

	expectedDecimalSeparatorRune := '.'

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualDecimalSeparatorRune := intAry.GetDecimalSeparator()

	if expectedDecimalSeparatorRune != actualDecimalSeparatorRune {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparatorRune != actualDecimalSeparatorRune\n"+
			"Expected actualDecimalSeparatorRune = '%v'\n"+
			"  Actual actualDecimalSeparatorRune = '%v'\n\n",
			ePrefix, expectedDecimalSeparatorRune, actualDecimalSeparatorRune)

		return
	}

	return
}

func TestIntAry_GetDecimalSeparator_03(t *testing.T) {

	ePrefix := "TestIntAry_GetDecimalSeparator_03"

	originalInt := 5064

	originalPrecisionUint := uint(2)

	originalNumberStr := "50.64"

	expectedDecimalSeparatorRune := '.'

	intAry, err := new(IntAry).NewInt(originalInt, originalPrecisionUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewInt(originalInt, originalPrecisionUint)\n"+
			"originalInt= '%v'\n"+
			"originalPrecisionUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalInt,
			originalPrecisionUint,
			err.Error())

		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualDecimalSeparatorRune := intAry.GetDecimalSeparator()

	if expectedDecimalSeparatorRune != actualDecimalSeparatorRune {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparatorRune != actualDecimalSeparatorRune\n"+
			"Expected actualDecimalSeparatorRune = '%v'\n"+
			"  Actual actualDecimalSeparatorRune = '%v'\n\n",
			ePrefix, expectedDecimalSeparatorRune, actualDecimalSeparatorRune)

		return
	}

	return
}

func TestIntAry_GetDecimalSeparator_04(t *testing.T) {

	ePrefix := "TestIntAry_GetDecimalSeparator_04"

	originalNumberStr := "450 123 647,1234"

	var originalFrenchDecSeparator rune

	originalFrenchDecSeparator = ','

	var originalFrenchThousandsSeparator rune

	originalFrenchThousandsSeparator = ' '

	expectedNumberStr := "450123647,1234"

	expectedPrecisionUint := uint(4)

	expectedSignValue := 1

	expectedDecimalSeparatorRune := ','

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetDecimalSeparator(originalFrenchDecSeparator)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetDecimalSeparator(originalFrenchDecSeparator)\n"+
			"originalFrenchDecSeparator= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalFrenchDecSeparator, err.Error())
		return
	}

	err = intAry.SetThousandsSeparator(originalFrenchThousandsSeparator)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetThousandsSeparator(\n"+
			"  originalFrenchThousandsSeparator)\n"+
			"originalFrenchThousandsSeparator= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalFrenchThousandsSeparator, err.Error())
		return
	}

	err = intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	actualDecimalSeparatorRune := intAry.GetDecimalSeparator()

	if expectedDecimalSeparatorRune != actualDecimalSeparatorRune {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparatorRune != actualDecimalSeparatorRune\n"+
			"Expected actualDecimalSeparatorRune = '%v'\n"+
			"  Actual actualDecimalSeparatorRune = '%v'\n\n",
			ePrefix, expectedDecimalSeparatorRune, actualDecimalSeparatorRune)

		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetInt_01(t *testing.T) {

	ePrefix := "TestIntAry_GetInt_01"

	originalNumberStr := "50"

	expectedNumberStr := "50"

	expectedNumInt := 50

	expectedPrecisionUint := uint(0)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultInt, err := intAry.GetInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultInt, err := intAry.GetInt()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumInt != resultInt {
		t.Errorf("%v\n"+
			"Error: Integer Value is INVALID!\n"+
			"Because expectedNumInt != resultInt\n"+
			"Expected resultInt = '%v'\n"+
			"  Actual resultInt = '%v'\n\n",
			ePrefix, expectedNumInt, resultInt)

		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetInt_02(t *testing.T) {

	ePrefix := "TestIntAry_GetInt_02"

	originalNumberStr := "-50"

	expectedNumberStr := "-50"

	expectedNumInt := -50

	expectedPrecisionUint := uint(0)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultInt, err := intAry.GetInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultInt, err := intAry.GetInt()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumInt != resultInt {
		t.Errorf("%v\n"+
			"Error: Integer Value is INVALID!\n"+
			"Because expectedNumInt != resultInt\n"+
			"Expected resultInt = '%v'\n"+
			"  Actual resultInt = '%v'\n\n",
			ePrefix, expectedNumInt, resultInt)

		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetInt_03(t *testing.T) {

	ePrefix := "TestIntAry_GetInt_03"

	originalNumberStr := "2147483647"

	expectedNumberStr := "2147483647"

	expectedNumInt := 2147483647

	expectedPrecisionUint := uint(0)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	resultInt, err := intAry.GetInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultInt, err := intAry.GetInt()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryPrecisionUint, err := intAry.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryPrecisionUint, err :=\n"+
			"  intAry.GetPrecisionUint()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intArySignValue, err := intAry.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArySignValue, err := intAry.GetSign()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
			"intAry= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumInt != resultInt {
		t.Errorf("%v\n"+
			"Error: Integer Value is INVALID!\n"+
			"Because expectedNumInt != resultInt\n"+
			"Expected resultInt = '%v'\n"+
			"  Actual resultInt = '%v'\n\n",
			ePrefix, expectedNumInt, resultInt)

		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	if expectedPrecisionUint != intAryPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryPrecisionUint\n"+
			"Expected intAryPrecisionUint = '%v'\n"+
			"  Actual intAryPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryPrecisionUint)

		return
	}

	if expectedSignValue != intArySignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intArySignValue\n"+
			"Expected intArySignValue = '%v'\n"+
			"  Actual intArySignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intArySignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryNumSeps \n"+
			"Expected intAryNumSeps = '%v'\n"+
			"  Actual intAryNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetInt_04(t *testing.T) {

	ePrefix := "TestIntAry_GetInt_04"

	originalNumberStr := "2147483648"

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetInt()

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call: _, err = intAry.GetInt()\n"+
			"originalNumberStr= '%v'\n"+
			"A numeric value greater than maximum 'int' should produce an error.\n\n",
			ePrefix, originalNumberStr)
		return
	}

	return
}

func TestIntAry_GetInt_05(t *testing.T) {

	ePrefix := "TestIntAry_GetInt_05"

	originalNumberStr := "-2147483649"

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetInt()

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call: _, err = intAry.GetInt()\n"+
			"originalNumberStr= '%v'\n"+
			"A numeric value less than minimum 'int' should produce an error.\n\n",
			ePrefix, originalNumberStr)
		return
	}

	return
}

func TestIntAry_GetIntAry_01(t *testing.T) {

	ePrefix := "TestIntAry_GetIntAry_01"

	originalNumberStr := "12345694829"

	expectedNumberStr := "12345694829"

	expectedAry := []uint8{1, 2, 3, 4, 5, 6, 9, 4, 8, 2, 9}

	expectedArrayLen := len(expectedAry) // 11

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedArrayLen != intArrayElementsLen {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because expectedArrayLen != intArrayElementsLen\n"+
			"Expected intArrayElementsLen = '%v'\n"+
			"  Actual intArrayElementsLen = '%v'\n\n",
			ePrefix, expectedArrayLen, intArrayElementsLen)

		return
	}

	for i := 0; i < intArrayElementsLen; i++ {

		if expectedAry[i] != intArrayElements[i] {
			t.Errorf("%v\n"+
				"Error: Expected and Actual Array Elements Don't Match!\n"+
				"Because expectedAry[%v] != intArrayElements[%v]\n"+
				"Expected intArrayElements[%v] = '%v'\n"+
				"  Actual intArrayElements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"IntAry Number String   = '%v'\n",
				ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
				expectedNumberStr, intAryNumberStr)

			return
		}

	}

	return
}

func TestIntAry_GetIntAry_02(t *testing.T) {

	ePrefix := "TestIntAry_GetIntAry_02"

	originalNumberStr := "987654321888629"

	expectedNumberStr := "987654321888629"

	expectedAry := []uint8{9, 8, 7, 6, 5, 4, 3, 2, 1, 8, 8, 8, 6, 2, 9}

	expectedArrayLen := len(expectedAry) // 15

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedArrayLen != intArrayElementsLen {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because expectedArrayLen != intArrayElementsLen\n"+
			"Expected intArrayElementsLen = '%v'\n"+
			"  Actual intArrayElementsLen = '%v'\n\n",
			ePrefix, expectedArrayLen, intArrayElementsLen)

		return
	}

	for i := 0; i < intArrayElementsLen; i++ {

		if expectedAry[i] != intArrayElements[i] {
			t.Errorf("%v\n"+
				"Error: Expected and Actual Array Elements Don't Match!\n"+
				"Because expectedAry[%v] != intArrayElements[%v]\n"+
				"Expected intArrayElements[%v] = '%v'\n"+
				"  Actual intArrayElements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"IntAry Number String   = '%v'\n",
				ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
				expectedNumberStr, intAryNumberStr)

			return
		}

	}

	return

}

func TestIntAry_GetIntAryDeepCopy(t *testing.T) {

	ePrefix := "TestIntAry_GetIntAryDeepCopy"

	originalNumberStr := "12345694829"

	expectedNumberStr := "12345694829"

	expectedAry := []uint8{1, 2, 3, 4, 5, 6, 9, 4, 8, 2, 9}

	expectedArrayLen := len(expectedAry) // 11

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	intArrayElements, intArrayElementsLen, err := intAry.GetIntAryDeepCopy()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArrayElements, intArrayElementsLen, err :=\n"+
			"  intAry.GetIntAryDeepCopy()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedArrayLen != intArrayElementsLen {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because expectedArrayLen != intArrayElementsLen\n"+
			"Expected intArrayElementsLen = '%v'\n"+
			"  Actual intArrayElementsLen = '%v'\n\n",
			ePrefix, expectedArrayLen, intArrayElementsLen)

		return
	}

	for i := 0; i < intArrayElementsLen; i++ {

		if expectedAry[i] != intArrayElements[i] {
			t.Errorf("%v\n"+
				"Error: Expected and Actual Array Elements Don't Match!\n"+
				"Because expectedAry[%v] != intArrayElements[%v]\n"+
				"Expected intArrayElements[%v] = '%v'\n"+
				"  Actual intArrayElements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"IntAry Number String   = '%v'\n",
				ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
				expectedNumberStr, intAryNumberStr)

			return
		}

	}

	return
}

func TestIntAry_GetIntAryDeepCopy_02(t *testing.T) {

	ePrefix := "TestIntAry_GetIntAryDeepCopy"

	originalNumberStr := "12345690888632"

	expectedNumberStr := "12345690888632"

	expectedAry := []uint8{1, 2, 3, 4, 5, 6, 9, 0, 0, 8, 8, 8, 6, 3, 2}

	expectedArrayLen := len(expectedAry) // 15

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	intArrayElements, intArrayElementsLen, err := intAry.GetIntAryDeepCopy()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intArrayElements, intArrayElementsLen, err :=\n"+
			"  intAry.GetIntAryDeepCopy()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedArrayLen != intArrayElementsLen {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
			"Because expectedArrayLen != intArrayElementsLen\n"+
			"Expected intArrayElementsLen = '%v'\n"+
			"  Actual intArrayElementsLen = '%v'\n\n",
			ePrefix, expectedArrayLen, intArrayElementsLen)

		return
	}

	for i := 0; i < intArrayElementsLen; i++ {

		if expectedAry[i] != intArrayElements[i] {
			t.Errorf("%v\n"+
				"Error: Expected and Actual Array Elements Don't Match!\n"+
				"Because expectedAry[%v] != intArrayElements[%v]\n"+
				"Expected intArrayElements[%v] = '%v'\n"+
				"  Actual intArrayElements[%v] = '%v'\n"+
				"Expected Number String = '%v'\n"+
				"IntAry Number String   = '%v'\n",
				ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
				expectedNumberStr, intAryNumberStr)

			return
		}

	}

	return
}

func TestIntAry_GetInt64_01(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_01"

	originalNumberStr := "50"

	expectedNumberInt64 := int64(50)

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	resultNumberInt64, err := intAry.GetInt64()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumberInt64, err := intAry.GetInt64()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberInt64 != resultNumberInt64 {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Int64 Values ARE NOT EQUAL!\n"+
			"Because expectedNumberInt64 != resultNumberInt64\n"+
			"Expected resultNumberInt64 = '%v'\n"+
			"  Actual resultNumberInt64 = '%v'\n\n",
			ePrefix, expectedNumberInt64, resultNumberInt64)

		return
	}

	return
}

func TestIntAry_GetInt64_02(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_02"

	originalNumberStr := "-50"

	expectedNumberInt64 := int64(-50)

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	resultNumberInt64, err := intAry.GetInt64()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumberInt64, err := intAry.GetInt64()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberInt64 != resultNumberInt64 {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Int64 Values ARE NOT EQUAL!\n"+
			"Because expectedNumberInt64 != resultNumberInt64\n"+
			"Expected resultNumberInt64 = '%v'\n"+
			"  Actual resultNumberInt64 = '%v'\n\n",
			ePrefix, expectedNumberInt64, resultNumberInt64)

		return
	}

	return
}

func TestIntAry_GetInt64_03(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_03"

	originalNumberStr := "9223372036854775807"

	expectedNumberInt64 := int64(9223372036854775807)

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	resultNumberInt64, err := intAry.GetInt64()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumberInt64, err := intAry.GetInt64()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberInt64 != resultNumberInt64 {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Int64 Values ARE NOT EQUAL!\n"+
			"Because expectedNumberInt64 != resultNumberInt64\n"+
			"Expected resultNumberInt64 = '%v'\n"+
			"  Actual resultNumberInt64 = '%v'\n\n",
			ePrefix, expectedNumberInt64, resultNumberInt64)

		return
	}

	return
}

func TestIntAry_GetInt64_04(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_04"

	originalNumberStr := "9223372036854775808"

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetInt64()

	if err == nil {
		t.Errorf("%v\n"+
			"Function Call: _, err = intAry.GetInt64()\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"originalNumberStr= %v\n"+
			"originalNumberStr exceeds Int64 max value and should produce an error.\n\n",
			ePrefix, originalNumberStr)
		return
	}

	return
}

func TestIntAry_GetInt64_05(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_05"

	originalNumberStr := "-9223372036854775809"

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetInt64()

	if err == nil {
		t.Errorf("%v\n"+
			"Function Call: _, err = intAry.GetInt64()\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"originalNumberStr= %v\n"+
			"originalNumberStr is less than Int64 minimum value and should produce an error.\n\n",
			ePrefix, originalNumberStr)
		return
	}

	return
}

func TestIntAry_GetInt64_06(t *testing.T) {

	ePrefix := "TestIntAry_GetInt64_06"

	originalNumberStr := "-9223372036854775808"

	expectedNumberInt64 := int64(-9223372036854775808)

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	resultNumberInt64, err := intAry.GetInt64()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultNumberInt64, err := intAry.GetInt64()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedNumberInt64 != resultNumberInt64 {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Int64 Values ARE NOT EQUAL!\n"+
			"Because expectedNumberInt64 != resultNumberInt64\n"+
			"Expected resultNumberInt64 = '%v'\n"+
			"  Actual resultNumberInt64 = '%v'\n\n",
			ePrefix, expectedNumberInt64, resultNumberInt64)

		return
	}

	return
}

func TestIntAry_GetFractionalDigits_01(t *testing.T) {

	ePrefix := "TestIntAry_GetFractionalDigits_01"

	originalNumberStr := "2.7894"

	expectedNumberStr := "0.7894"

	expectedPrecisionInt := 4

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if false != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'true' and invalid!\n"+
			"Because false != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, false, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetFractionalDigits_02(t *testing.T) {

	ePrefix := "TestIntAry_GetFractionalDigits_02"

	originalNumberStr := "2"

	expectedNumberStr := "0"

	expectedPrecisionInt := 0

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if true != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'false' and invalid!\n"+
			"Because true != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, true, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetFractionalDigits_03(t *testing.T) {

	ePrefix := "TestIntAry_GetFractionalDigits_03"

	originalNumberStr := "2.00"

	expectedNumberStr := "0.00"

	expectedPrecisionInt := 2

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if true != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'false' and invalid!\n"+
			"Because true != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, true, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetFractionalDigits_04(t *testing.T) {

	ePrefix := "TestIntAry_GetFractionalDigits_04"

	originalNumberStr := "-2.978562154907"

	expectedNumberStr := "0.978562154907"

	expectedPrecisionInt := 12

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if false != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'true' and invalid!\n"+
			"Because false != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, false, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetIntegerDigits_01(t *testing.T) {

	ePrefix := "TestIntAry_GetIntegerDigits_01"

	originalNumberStr := "997562.4692"

	expectedNumberStr := "997562"

	expectedPrecisionInt := 0

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetIntegerDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetIntegerDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if false != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'true' and invalid!\n"+
			"Because false != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, false, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetIntegerDigits_02(t *testing.T) {

	ePrefix := "TestIntAry_GetFractionalDigits_01"

	originalNumberStr := "0.4692"

	expectedNumberStr := "0"

	expectedPrecisionInt := 0

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if true != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'false' and invalid!\n"+
			"Because true != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, true, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetIntegerDigits_03(t *testing.T) {

	ePrefix := "TestIntAry_GetIntegerDigits_03"

	originalNumberStr := "-987.4692"

	expectedNumberStr := "-987"

	expectedPrecisionInt := 0

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if false != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'true' and invalid!\n"+
			"Because true != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, false, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetIntegerDigits_04(t *testing.T) {

	ePrefix := "TestIntAry_GetIntegerDigits_04"

	originalNumberStr := "-0.4692"

	expectedNumberStr := "0"

	expectedPrecisionInt := 0

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry1 := new(IntAry).New()

	err := intAry1.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := ia.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry1.IsValid("Validating initial intAry1")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry1.IsValid('Validating initial intAry1')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry1NumberStr, err := intAry1.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry1NumberStr, err := intAry1.GetNumStr()\n"+
			"intAry1 set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAry1NumberStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Because originalNumberStr != intAry1NumberStr\n"+
			"Expected intAry1NumberStr = '%v'\n"+
			"  Actual intAry1NumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAry1NumberStr)

		return
	}

	intAry2, err := intAry1.GetFractionalDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2, err := intAry1.GetFractionalDigits()\n"+
			"intAry1= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry1NumberStr, err.Error())
		return
	}

	err = intAry2.IsValid("Validating intAry2")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry2.IsValid('Validating intAry2')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2NumberStr, err := intAry2.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumberStr, err := intAry2.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAry2PrecisionInt := intAry2.GetPrecision()

	intAry2SignValue, err := intAry2.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2SignValue, err := intAry2.GetSign()\n"+
			"intAry2= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAry2NumberStr, err.Error())
		return
	}

	intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry2NumSeps, err := intAry2.GetNumericSeparatorsDto()\n"+
			"intAry2= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAry2NumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAry2NumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Decimal Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAry2NumberStr \n"+
			"Expected intAry2NumberStr = '%v'\n"+
			"  Actual intAry2NumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAry2NumberStr)

		return
	}

	if expectedPrecisionInt != intAry2PrecisionInt {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionInt != intAry2PrecisionInt\n"+
			"Expected intAry2PrecisionInt = '%v'\n"+
			"  Actual intAry2PrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, intAry2PrecisionInt)

		return
	}

	if expectedSignValue != intAry2SignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAry2 Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAry2SignValue\n"+
			"Expected intAry2SignValue = '%v'\n"+
			"  Actual intAry2SignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAry2SignValue)

		return
	}

	if !expectedNumSeps.Equal(intAry2NumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAry2NumSeps \n"+
			"Expected intAry2NumSeps = '%v'\n"+
			"  Actual intAry2NumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAry2NumSeps.String())

		return
	}

	iAry2Stats := intAry2.GetIntAryStats()

	iAry2StatsIsZeroValue := iAry2Stats.IsZeroValue

	if true != iAry2StatsIsZeroValue {
		t.Errorf("%v\n"+
			"Error: iAry2Stats.IsZeroValue is 'false' and invalid!\n"+
			"Because true != iAry2StatsIsZeroValue\n"+
			"Expected iAry2StatsIsZeroValue = '%v'\n"+
			"  Actual iAry2StatsIsZeroValue = '%v'\n\n",
			ePrefix, true, iAry2StatsIsZeroValue)

		return
	}

	return
}

func TestIntAry_GetMagnitude_01(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitude_01"

	originalNumberStr := "98327123"

	expectedMagnitudeInt := 7

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualMagnitudeInt, err := intAry.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualMagnitudeInt, err := intAry.GetMagnitude()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedMagnitudeInt != actualMagnitudeInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Magnitude values DON'T MATCH!\n"+
			"Because expectedMagnitudeInt != actualMagnitudeInt\n"+
			"Expected actualMagnitudeInt = '%v'\n"+
			"  Actual actualMagnitudeInt = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, actualMagnitudeInt)

		return
	}

	return
}

func TestIntAry_GetMagnitude_02(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitude_02"

	originalNumberStr := "98327123.1234"

	expectedMagnitudeInt := 7

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualMagnitudeInt, err := intAry.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualMagnitudeInt, err := intAry.GetMagnitude()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedMagnitudeInt != actualMagnitudeInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Magnitude values DON'T MATCH!\n"+
			"Because expectedMagnitudeInt != actualMagnitudeInt\n"+
			"Expected actualMagnitudeInt = '%v'\n"+
			"  Actual actualMagnitudeInt = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, actualMagnitudeInt)

		return
	}

	return
}

func TestIntAry_GetMagnitude_03(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitude_03"

	originalNumberStr := "-98327123"

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetMagnitude()

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call: _, err = intAry.GetMagnitude()\n"+
			"intAry= %v\n"+
			"A negative 'intAry' value should produce an error.\n\n",
			ePrefix, intAryNumberStr)
		return
	}
}

func TestIntAry_GetMagnitudeDigits_01(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitudeDigits_01"

	originalNumberStr := "98327123"

	expectedMagnitudeInt := 8

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedMagnitudeInt != actualMagnitudeDigitsInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Magnitude Digits Values DON'T MATCH!\n"+
			"Because expectedMagnitudeInt != actualMagnitudeDigitsInt\n"+
			"Expected actualMagnitudeDigitsInt = '%v'\n"+
			"  Actual actualMagnitudeDigitsInt = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, actualMagnitudeDigitsInt)

		return
	}

	return
}

func TestIntAry_GetMagnitudeDigits_02(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitudeDigits_02"

	originalNumberStr := "-98327123"

	expectedMagnitudeInt := 8

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedMagnitudeInt != actualMagnitudeDigitsInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Magnitude Digits Values DON'T MATCH!\n"+
			"Because expectedMagnitudeInt != actualMagnitudeDigitsInt\n"+
			"Expected actualMagnitudeDigitsInt = '%v'\n"+
			"  Actual actualMagnitudeDigitsInt = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, actualMagnitudeDigitsInt)

		return
	}

	return
}

func TestIntAry_GetMagnitudeDigits_03(t *testing.T) {

	ePrefix := "TestIntAry_GetMagnitudeDigits_03"

	originalNumberStr := "98327123.123"

	expectedMagnitudeInt := 8

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualMagnitudeDigitsInt, err := intAry.GetMagnitudeDigits()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedMagnitudeInt != actualMagnitudeDigitsInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Magnitude Digits Values DON'T MATCH!\n"+
			"Because expectedMagnitudeInt != actualMagnitudeDigitsInt\n"+
			"Expected actualMagnitudeDigitsInt = '%v'\n"+
			"  Actual actualMagnitudeDigitsInt = '%v'\n\n",
			ePrefix, expectedMagnitudeInt, actualMagnitudeDigitsInt)

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_01(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_01"

	radicandNumberStr := "125"

	nthRootInt := 5

	maxPrecisionInt := 14

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "2.62652780440377"

	expectedPrecisionUint := uint(14)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_02(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_02"

	radicandNumberStr := "5604423"

	nthRootInt := 6

	maxPrecisionInt := 13

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "13.3276982415963"

	expectedPrecisionUint := uint(13)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_03(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_03"

	radicandNumberStr := "5604423.924"

	nthRootInt := 6

	maxPrecisionInt := 13

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "13.3276986078187"

	expectedPrecisionUint := uint(13)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_04(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_04"

	radicandNumberStr := "-27"

	nthRootInt := 3

	maxPrecisionInt := 2

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "-3.00"

	expectedPrecisionUint := uint(2)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_05(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_05"

	radicandNumberStr := "-27"

	nthRootInt := 4

	maxPrecisionInt := 2

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"  _, err = intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"radicand= %v; nthRootInt= %v\n"+
			"A negative 'radicand' value with an even nthRoot should produce an error.\n\n",
			ePrefix, radicandNumberStr, nthRootInt)
		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_06(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_06"

	radicandNumberStr := "-5604423.924"

	nthRootInt := 5

	maxPrecisionInt := 13

	//                                 1         2         3
	//                      0.1234567890123456789012345678901234567
	expectedNumberStr := "-22.3720713464898"

	expectedPrecisionUint := uint(13)

	expectedSignValue := -1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_07(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_07"

	radicandNumberStr := "5604423.924"

	nthRootInt := 0

	maxPrecisionInt := 1

	//                               1         2         3
	//                    0.1234567890123456789012345678901234567
	expectedNumberStr := "1.0"

	expectedPrecisionUint := uint(1)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.\n"+
			"  GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"intAry= '%v'\n"+
			"nthRoot= '%v'\n"+
			"maxPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			nthRootInt,
			maxPrecisionInt,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating final intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating final intAryResult')\n"+
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

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetNthRootOfThis_08(t *testing.T) {

	ePrefix := "TestIntAry_GetNthRootOfThis_08"

	radicandNumberStr := "27"

	nthRootInt := 1

	maxPrecisionInt := 2

	intAry, err := new(IntAry).NewNumStr(radicandNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(radicandNumberStr)\n"+
			"radicandNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, radicandNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if radicandNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and IntAry Number String Values ARE NOT Equal\n"+
			"Because radicandNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, radicandNumberStr, intAryNumberStr)

		return
	}

	_, err = intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			"  _, err = intAry.GetNthRootOfThis(nthRootInt, maxPrecisionInt)\n"+
			"nthRootInt= %v\n"+
			"nthRoot == 1 should produce an error.\n\n",
			ePrefix, nthRootInt)
		return
	}

	return
}

func TestIntAry_GetNumStrDto_01(t *testing.T) {

	ePrefix := "TestIntAry_GetNumStrDto_01"

	//                                    1         2         3
	//                         0.1234567890123456789012345678901234567
	originalNumberStr := "589627.123456"

	expectedNumberStr := originalNumberStr

	expectedPrecisionUint := uint(6)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedNumStrDto, err := new(NumStrDto).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedNumStrDto, err := new(NumStrDto).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualNumStrDto, err := intAry.GetNumStrDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStrDto, err := intAry.GetNumStrDto()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	err = actualNumStrDto.IsValid("Validating initial actualNumStrDto")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = actualNumStrDto.IsValid('Validating initial actualNumStrDto')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStrDtoNumberStr, err := actualNumStrDto.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStrDtoNumberStr, err := actualNumStrDto.GetNumStr()\n"+
			"actualNumStrDto set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	actualNumStrDtoPrecisionUint, err := actualNumStrDto.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStrDtoPrecisionUint, err :=\n"+
			"  actualNumStrDto.GetPrecisionUint()\n"+
			"actualNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualNumStrDtoNumberStr, err.Error())
		return
	}

	actualNumStrDtoSignValue, err := actualNumStrDto.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStrDtoSignValue, err := actualNumStrDto.GetSign()\n"+
			"actualNumStrDto= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, actualNumStrDtoNumberStr, err.Error())
		return
	}

	actualNumStrDtoNumSeps, err := actualNumStrDto.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualNumStrDtoNumSeps, err := actualNumStrDto.GetNumericSeparatorsDto()\n"+
			"actualNumStrDto= '%v\n"+
			"Error= '%v'\n\n", ePrefix, actualNumStrDtoNumberStr, err.Error())
		return
	}

	if expectedNumberStr != actualNumStrDtoNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != actualNumStrDtoNumberStr \n"+
			"Expected actualNumStrDtoNumberStr = '%v'\n"+
			"  Actual actualNumStrDtoNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, actualNumStrDtoNumberStr)

		return
	}

	if expectedPrecisionUint != actualNumStrDtoPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & Actual Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != actualNumStrDtoPrecisionUint\n"+
			"Expected actualNumStrDtoPrecisionUint = '%v'\n"+
			"  Actual actualNumStrDtoPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, actualNumStrDtoPrecisionUint)

		return
	}

	if expectedSignValue != actualNumStrDtoSignValue {
		t.Errorf("%v\n"+
			"Error: expected & actualNumStrDto Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != actualNumStrDtoSignValue\n"+
			"Expected actualNumStrDtoSignValue = '%v'\n"+
			"  Actual actualNumStrDtoSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, actualNumStrDtoSignValue)

		return
	}

	if !expectedNumSeps.Equal(actualNumStrDtoNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != actualNumStrDtoNumSeps \n"+
			"Expected actualNumStrDtoNumSeps = '%v'\n"+
			"  Actual actualNumStrDtoNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), actualNumStrDtoNumSeps.String())

		return
	}

	expectedNumStrDtoEqualsActualNumStrDto := expectedNumStrDto.Equal(actualNumStrDto)

	if false == expectedNumStrDtoEqualsActualNumStrDto {
		t.Errorf("%v\n"+
			"Error: expectedNumStrDto NOT EQUAL TO actualNumStrDto\n"+
			"Because false == expectedNumStrDtoEqualsActualNumStrDto\n"+
			"Expected expectedNumStrDtoEqualsActualNumStrDto = '%v'\n"+
			"  Actual expectedNumStrDtoEqualsActualNumStrDto = '%v'\n\n",
			ePrefix, true, expectedNumStrDtoEqualsActualNumStrDto)

		return
	}

	return
}

func TestIntAry_GetScaleFactor_01(t *testing.T) {

	ePrefix := "TestIntAry_GetScaleFactor_01"

	//                                  1         2         3
	//                       0.1234567890123456789012345678901234567
	originalNumberStr := "2686.12345"

	expectedScaleFactor := big.NewInt(100000)

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryScaleFactor, err := intAry.GetScaleFactor()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryScaleFactor, err := intAry.GetScaleFactor()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryNumberStr, err.Error())
		return
	}

	if expectedScaleFactor.Cmp(intAryScaleFactor) != 0 {
		t.Errorf("%v\n"+
			"Error: expectedScaleFactor NOT EQUAL TO intAryScaleFactor!\n"+
			"Because expectedScaleFactor.Cmp(intAryScaleFactor) != 0\n"+
			"Expected intAryScaleFactor = '%v'\n"+
			"  Actual intAryScaleFactor = '%v'\n\n",
			ePrefix, expectedScaleFactor.Text(10), intAryScaleFactor.Text(10))

		return
	}

	return
}

func TestIntAry_GetSquareRootInt_01(t *testing.T) {

	ePrefix := "TestIntAry_GetSquareRootInt_01"

	originalNumberStr := "2686.5"

	maxPrecision := 30

	//                                1         2         3
	//                     0.1234567890123456789012345678901234567
	expectedNumberStr := "51.831457629512986714934518985668"

	expectedPrecisionUint := uint(30)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	intAryResult, err := intAry.GetSquareRootOfThis(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResult, err := intAry.GetSquareRootOfThis(maxPrecision)\n"+
			"intAry= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			maxPrecision,
			err.Error())

		return
	}

	err = intAryResult.IsValid("Validating initial intAryResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryResult.IsValid('Validating initial intAryResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryResultNumberStr, err := intAryResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryResultNumberStr, err := intAryResult.GetNumStr()\n"+
			"intAryResult set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

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
			"Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryResultNumberStr \n"+
			"Expected intAryResultNumberStr = '%v'\n"+
			"  Actual intAryResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryResultNumberStr)

		return
	}

	if expectedPrecisionUint != intAryResultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryResultPrecisionUint\n"+
			"Expected intAryResultPrecisionUint = '%v'\n"+
			"  Actual intAryResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryResultPrecisionUint)

		return
	}

	if expectedSignValue != intAryResultSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryResult Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryResultSignValue\n"+
			"Expected intAryResultSignValue = '%v'\n"+
			"  Actual intAryResultSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryResultNumSeps \n"+
			"Expected intAryResultNumSeps = '%v'\n"+
			"  Actual intAryResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryResultNumSeps.String())

		return
	}

	return
}

func TestIntAry_GetThousandsSeparator_01(t *testing.T) {

	ePrefix := "TestIntAry_GetThousandsSeparator_01"

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	intAry := new(IntAry).New()

	actualDecimalSeparator := intAry.GetDecimalSeparator()

	if expectedDecimalSeparator != actualDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != actualDecimalSeparator\n"+
			"Expected actualDecimalSeparator = '%v'\n"+
			"  Actual actualDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, actualDecimalSeparator)

		return
	}

	return
}

func TestIntAry_GetThousandsSeparator_02(t *testing.T) {

	ePrefix := "TestIntAry_GetThousandsSeparator_01"

	originalNumberStr := "50.47"

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualDecimalSeparator := intAry.GetDecimalSeparator()

	if expectedDecimalSeparator != actualDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != actualDecimalSeparator\n"+
			"Expected actualDecimalSeparator = '%v'\n"+
			"  Actual actualDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, actualDecimalSeparator)

		return
	}

	return
}

func TestIntAry_GetThousandsSeparator_03(t *testing.T) {

	ePrefix := "TestIntAry_GetThousandsSeparator_03"

	originalFloat64 := 50.47

	originalPrecisionInt := 2

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	intAry, err := new(IntAry).NewFloat64(originalFloat64, originalPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewFloat64(originalFloat64, originalPrecisionInt)\n"+
			"originalFloat64= '%v'\n"+
			"originalPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			strconv.FormatFloat(originalFloat64, 'f', originalPrecisionInt, 64),
			originalPrecisionInt,
			err.Error())

		return
	}

	actualDecimalSeparator := intAry.GetDecimalSeparator()

	if expectedDecimalSeparator != actualDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != actualDecimalSeparator\n"+
			"Expected actualDecimalSeparator = '%v'\n"+
			"  Actual actualDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, actualDecimalSeparator)

		return
	}

	return
}

func TestIntAry_GetThousandsSeparator_04(t *testing.T) {

	ePrefix := "TestIntAry_GetThousandsSeparator_04"

	originalNumberStr := "450 123 647,1234"

	var frenchDecimalSeparator rune

	frenchDecimalSeparator = ','

	expectedDecimalSeparator := frenchDecimalSeparator

	var frenchThousandsSeparator rune

	frenchThousandsSeparator = ' '

	expectedNumberStr := "450123647,1234"

	intAry := new(IntAry).New()

	err := intAry.SetDecimalSeparator(frenchDecimalSeparator)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetDecimalSeparator(frenchDecimalSeparator)\n"+
			"frenchDecimalSeparator= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, frenchDecimalSeparator, err.Error())

		return
	}

	err = intAry.SetThousandsSeparator(frenchThousandsSeparator)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetThousandsSeparator(frenchThousandsSeparator)n"+
			"frenchThousandsSeparator= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, frenchThousandsSeparator, err.Error())
		return
	}

	err = intAry.SetIntAryWithNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating final intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating final intAry')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	actualDecimalSeparator := intAry.GetDecimalSeparator()

	if expectedDecimalSeparator != actualDecimalSeparator {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Decimal Separators ARE NOT EQUAL!\n"+
			"Because expectedDecimalSeparator != actualDecimalSeparator\n"+
			"Expected actualDecimalSeparator = '%v'\n"+
			"  Actual actualDecimalSeparator = '%v'\n\n",
			ePrefix, expectedDecimalSeparator, actualDecimalSeparator)

		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAryNumberStr set to final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and IntAry Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != intAryNumberStr\n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryNumberStr)

		return
	}

	return
}

func TestIntAry_IncrementIntegerOne_01(t *testing.T) {

	ePrefix := "TestIntAry_IncrementIntegerOne_01"

	originalNumerStr := "-100.123"

	expectedNumberStr := "100.123"

	incrementCycles := 200

	intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumerStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := intAry.SetIntAryWithNumStr(originalNumerStr)\n"+
			"originalNumerStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumerStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	for i := 0; i < incrementCycles; i++ {

		err = intAry.IncrementIntegerOne()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = intAry.IncrementIntegerOne()\n"+
				"intAry= '%v'\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		intAryNumberStr, err = intAry.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"intAryNumberStr, err = intAry.GetNumStr()\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

	}

	intAryFinalNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & Final number strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != intAryFinalNumberStr\n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	return
}

func TestIntAry_IncrementIntegerOne_02(t *testing.T) {

	ePrefix := "TestIntAry_IncrementIntegerOne_02"

	originalNumerStr := "-2000"

	expectedNumberStr := "2000"

	incrementCycles := 4000

	intAry, err := new(IntAry).NewNumStr(originalNumerStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumerStr)\n"+
			"originalNumerStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumerStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	for i := 0; i < incrementCycles; i++ {

		err = intAry.IncrementIntegerOne()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = intAry.IncrementIntegerOne()\n"+
				"intAry= '%v'\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		intAryNumberStr, err = intAry.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"intAryNumberStr, err = intAry.GetNumStr()\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

	}

	intAryFinalNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & Final number strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != intAryFinalNumberStr\n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	return
}

func TestIntAry_IncrementIntegerOne_03(t *testing.T) {

	ePrefix := "TestIntAry_IncrementIntegerOne_03"

	originalNumerStr := "-2000.123"

	expectedNumberStr := "2000.123"

	incrementCycles := 4000

	intAry, err := new(IntAry).NewNumStr(originalNumerStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumerStr)\n"+
			"originalNumerStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumerStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	for i := 0; i < incrementCycles; i++ {

		err = intAry.IncrementIntegerOne()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = intAry.IncrementIntegerOne()\n"+
				"intAry= '%v'\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		intAryNumberStr, err = intAry.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"intAryNumberStr, err = intAry.GetNumStr()\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

	}

	intAryFinalNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & Final number strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != intAryFinalNumberStr\n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	return
}

func TestIntAry_IncrementIntegerOne_04(t *testing.T) {

	ePrefix := "TestIntAry_IncrementIntegerOne_04"

	originalNumerStr := "0"

	expectedNumberStr := "40"

	incrementCycles := 40

	intAry, err := new(IntAry).NewNumStr(originalNumerStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumerStr)\n"+
			"originalNumerStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumerStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	incrementCycles = 40

	for i := 0; i < incrementCycles; i++ {

		err = intAry.IncrementIntegerOne()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"err = intAry.IncrementIntegerOne()\n"+
				"intAry= '%v'\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix,
				intAryNumberStr,
				i,
				err.Error())

			return
		}

		intAryNumberStr, err = intAry.GetNumStr()

		if err != nil {
			t.Errorf("%v\n"+
				"Error returned by:\n"+
				"intAryNumberStr, err = intAry.GetNumStr()\n"+
				"Increment Cycle No = '%v'\n"+
				"Error= '%v'\n\n",
				ePrefix, i, err.Error())
			return
		}

	}

	intAryFinalNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryFinalNumberStr, err = intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedNumberStr != intAryFinalNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected & Final number strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != intAryFinalNumberStr\n"+
			"Expected intAryFinalNumberStr = '%v'\n"+
			"  Actual intAryFinalNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryFinalNumberStr)

		return
	}

	return
}

func TestIntAry_Inverse_01(t *testing.T) {

	ePrefix := "TestIntAry_Inverse_01"

	originalNumberStr := "25"

	maxPrecision := 2

	expectedNumberStr := "0.04"

	expectedPrecisionUint := uint(2)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryBaseNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Base Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryBaseNumberStr \n"+
			"Expected intAryBaseNumberStr = '%v'\n"+
			"  Actual intAryBaseNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryBaseNumberStr)

		return
	}

	intAryInverse, err := intAryBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverse, err := intAryBase.Inverse(maxPrecision)\n"+
			"intAryBase= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryBaseNumberStr,
			maxPrecision,
			err.Error())

		return
	}

	err = intAryInverse.IsValid("Validating intAryInverse")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryInverse.IsValid('Validating intAryInverse')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInverseNumberStr, err := intAryInverse.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumberStr, err := intAryInverse.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInversePrecisionUint, err := intAryInverse.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInversePrecisionUint, err :=\n"+
			"  intAryInverse.GetPrecisionUint()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseSignValue, err := intAryInverse.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseSignValue, err := intAryInverse.GetSign()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()\n"+
			"intAryInverse= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryInverseNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryInverse Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryInverseNumberStr \n"+
			"Expected intAryInverseNumberStr = '%v'\n"+
			"  Actual intAryInverseNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryInverseNumberStr)

		return
	}

	if expectedPrecisionUint != intAryInversePrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryInversePrecisionUint\n"+
			"Expected intAryInversePrecisionUint = '%v'\n"+
			"  Actual intAryInversePrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryInversePrecisionUint)

		return
	}

	if expectedSignValue != intAryInverseSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryInverseSignValue\n"+
			"Expected intAryInverseSignValue = '%v'\n"+
			"  Actual intAryInverseSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryInverseSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryInverseNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryInverseNumSeps \n"+
			"Expected intAryInverseNumSeps = '%v'\n"+
			"  Actual intAryInverseNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryInverseNumSeps.String())

		return
	}

	return
}

func TestIntAry_Inverse_02(t *testing.T) {

	ePrefix := "TestIntAry_Inverse_02"

	originalNumberStr := "30517578125"

	maxPrecision := 15

	//                               1         2         3
	//                    0.123456789012345678901234567890
	expectedNumberStr := "0.000000000032768"

	expectedPrecisionUint := uint(15)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryBaseNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Base Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryBaseNumberStr \n"+
			"Expected intAryBaseNumberStr = '%v'\n"+
			"  Actual intAryBaseNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryBaseNumberStr)

		return
	}

	intAryInverse, err := intAryBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverse, err := intAryBase.Inverse(maxPrecision)\n"+
			"intAryBase= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryBaseNumberStr,
			maxPrecision,
			err.Error())

		return
	}

	err = intAryInverse.IsValid("Validating intAryInverse")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryInverse.IsValid('Validating intAryInverse')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInverseNumberStr, err := intAryInverse.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumberStr, err := intAryInverse.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInversePrecisionUint, err := intAryInverse.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInversePrecisionUint, err :=\n"+
			"  intAryInverse.GetPrecisionUint()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseSignValue, err := intAryInverse.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseSignValue, err := intAryInverse.GetSign()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()\n"+
			"intAryInverse= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryInverseNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryInverse Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryInverseNumberStr \n"+
			"Expected intAryInverseNumberStr = '%v'\n"+
			"  Actual intAryInverseNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryInverseNumberStr)

		return
	}

	if expectedPrecisionUint != intAryInversePrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryInversePrecisionUint\n"+
			"Expected intAryInversePrecisionUint = '%v'\n"+
			"  Actual intAryInversePrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryInversePrecisionUint)

		return
	}

	if expectedSignValue != intAryInverseSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryInverseSignValue\n"+
			"Expected intAryInverseSignValue = '%v'\n"+
			"  Actual intAryInverseSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryInverseSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryInverseNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryInverseNumSeps \n"+
			"Expected intAryInverseNumSeps = '%v'\n"+
			"  Actual intAryInverseNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryInverseNumSeps.String())

		return
	}

	return
}

func TestIntAry_Inverse_03(t *testing.T) {

	ePrefix := "TestIntAry_Inverse_03"

	originalNumberStr := "25"

	maxPrecision := 2

	expectedNumberStr := "0.04"

	expectedPrecisionUint := uint(2)

	expectedSignValue := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryBase, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
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
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryBaseNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and Base Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryBaseNumberStr \n"+
			"Expected intAryBaseNumberStr = '%v'\n"+
			"  Actual intAryBaseNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryBaseNumberStr)

		return
	}

	intAryInverse, err := intAryBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverse, err := intAryBase.Inverse(maxPrecision)\n"+
			"intAryBase= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryBaseNumberStr,
			maxPrecision,
			err.Error())

		return
	}

	err = intAryInverse.IsValid("Validating intAryInverse")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAryInverse.IsValid('Validating intAryInverse')\n"+
			"Validating Final IntAry value.\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInverseNumberStr, err := intAryInverse.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumberStr, err := intAryInverse.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryInversePrecisionUint, err := intAryInverse.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInversePrecisionUint, err :=\n"+
			"  intAryInverse.GetPrecisionUint()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseSignValue, err := intAryInverse.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseSignValue, err := intAryInverse.GetSign()\n"+
			"intAryInverse= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryInverseNumSeps, err := intAryInverse.GetNumericSeparatorsDto()\n"+
			"intAryInverse= '%v\n"+
			"Error= '%v'\n\n", ePrefix, intAryInverseNumberStr, err.Error())
		return
	}

	if expectedNumberStr != intAryInverseNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and intAryInverse Number String Values ARE NOT Equal\n"+
			"Because expectedNumberStr != intAryInverseNumberStr \n"+
			"Expected intAryInverseNumberStr = '%v'\n"+
			"  Actual intAryInverseNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, intAryInverseNumberStr)

		return
	}

	if expectedPrecisionUint != intAryInversePrecisionUint {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Precision Values ARE NOT EQUAL!\n"+
			"Because expectedPrecisionUint != intAryInversePrecisionUint\n"+
			"Expected intAryInversePrecisionUint = '%v'\n"+
			"  Actual intAryInversePrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, intAryInversePrecisionUint)

		return
	}

	if expectedSignValue != intAryInverseSignValue {
		t.Errorf("%v\n"+
			"Error: expected & intAryInverse Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignValue != intAryInverseSignValue\n"+
			"Expected intAryInverseSignValue = '%v'\n"+
			"  Actual intAryInverseSignValue = '%v'\n\n",
			ePrefix, expectedSignValue, intAryInverseSignValue)

		return
	}

	if !expectedNumSeps.Equal(intAryInverseNumSeps) {
		t.Errorf("%v\n"+
			"Error: Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != intAryInverseNumSeps \n"+
			"Expected intAryInverseNumSeps = '%v'\n"+
			"  Actual intAryInverseNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), intAryInverseNumSeps.String())

		return
	}

	return
}

func TestIntAry_IsEvenNumber_01(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_01"

	originalNumberStr := "24"

	var expectedIsEven bool

	expectedIsEven = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_02(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_02"

	originalNumberStr := "25"

	var expectedIsEven bool

	expectedIsEven = false

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_03(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_03"

	originalNumberStr := "4.44"

	var expectedIsEven bool

	expectedIsEven = false

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_04(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_04"

	originalNumberStr := "0"

	var expectedIsEven bool

	expectedIsEven = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_05(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_05"

	originalNumberStr := "-24"

	var expectedIsEven bool

	expectedIsEven = true

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_06(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_06"

	originalNumberStr := "-25"

	var expectedIsEven bool

	expectedIsEven = false

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}

func TestIntAry_IsEvenNumber_07(t *testing.T) {

	ePrefix := "TestIntAry_IsEvenNumber_07"

	originalNumberStr := "-4.44"

	var expectedIsEven bool

	expectedIsEven = false

	intAry, err := new(IntAry).NewNumStr(originalNumberStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
			"originalNumberStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, originalNumberStr, err.Error())
		return
	}

	err = intAry.IsValid("Validating initial intAry")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = intAry.IsValid('Validating initial intAry')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	intAryNumberStr, err := intAry.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intAryNumberStr, err := intAry.GetNumStr()\n"+
			"intAry set to initial value\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if originalNumberStr != intAryNumberStr {
		t.Errorf("%v\n"+
			"Error: Original and intAry Number String Values ARE NOT Equal\n"+
			"Because originalNumberStr != intAryNumberStr \n"+
			"Expected intAryNumberStr = '%v'\n"+
			"  Actual intAryNumberStr = '%v'\n\n",
			ePrefix, originalNumberStr, intAryNumberStr)

		return
	}

	actualIsEven, err := intAry.IsEvenNumber()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"actualIsEven, err := intAry.IsEvenNumber()\n"+
			"intAry= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			intAryNumberStr,
			err.Error())

		return
	}

	if expectedIsEven != actualIsEven {
		t.Errorf("%v\n"+
			"Error: actualIsEven Result Is INVALID!\n"+
			"Because expectedIsEven != actualIsEven\n"+
			"Expected actualIsEven = '%v'\n"+
			"  Actual actualIsEven = '%v'\n\n",
			ePrefix, expectedIsEven, actualIsEven)

		return
	}

	return
}
