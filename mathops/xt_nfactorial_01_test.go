package mathops

import (
	"math/big"
	"testing"
)

func TestNFactorial_CalcFactorialValueBigInt_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_01"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(1))

	expectedNumberStr := "120"

	expectedBigInt := big.NewInt(int64(120))

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_02"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(0))

	expectedNumberStr := "120"

	expectedBigInt := big.NewInt(int64(120))

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_03"

	bigIntUpperLimit := big.NewInt(int64(-5))

	bigIntLowerLimit := big.NewInt(int64(0))

	_, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"   bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"Upper limit (bigIntUpperLimit) MUST BE greater than lower limit.\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10))
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_04"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(-2))

	_, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Any lower limit value less than '+1' will produce an error.\n\n",
			ePrefix, bigIntLowerLimit.Text(10))
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_05"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(2))

	expectedNumberStr := "60"

	expectedBigInt := big.NewInt(int64(60))

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_06"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(4))

	expectedNumberStr := "5"

	expectedBigInt := big.NewInt(int64(5))

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_06"

	bigIntUpperLimit := big.NewInt(int64(23))

	bigIntLowerLimit := big.NewInt(int64(0))

	expectedNumberStr := "25852016738884976640000"

	expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumberStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
			"expectedNumStr= '%v'\n"+
			"Error: isOk == false\n\n",
			ePrefix,
			expectedNumberStr)

		return
	}

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_08(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_08"

	bigIntUpperLimit := big.NewInt(int64(5))

	bigIntLowerLimit := big.NewInt(int64(9))

	_, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'"+
			"bigIntLowerLimit= '%v'\n"+
			"Lower Limit MUST BE Less than Upper Limit.\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10))
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_09(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_09"

	bigIntUpperLimit := big.NewInt(int64(0))

	bigIntLowerLimit := big.NewInt(int64(1))

	_, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'"+
			"bigIntLowerLimit= '%v'\n"+
			"Lower Limit MUST BE Less than Upper Limit.\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10))
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigInt_10(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigInt_10"

	bigIntUpperLimit := big.NewInt(int64(23))

	bigIntLowerLimit := big.NewInt(int64(20))

	expectedNumberStr := "10626"

	expectedBigInt := big.NewInt(int64(10626))

	expectedPrecisionInt := 0

	expectedPrecisionUint := uint(expectedPrecisionInt)

	expectedSignVal := 1

	expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(bigIntUpperLimit, bigIntLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigInt(\n"+
			"  bigIntUpperLimit, bigIntLowerLimit)\n"+
			"bigIntUpperLimit= '%v'\n"+
			"bigIntLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntUpperLimit.Text(10),
			bigIntLowerLimit.Text(10),
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedNumberStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedNumberStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedNumberStr, bigIntNumResultNumberStr)

		return
	}

	if expectedBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_01"

	originalUpperLimitBaseInt := 5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "5"

	originalLowerLimitBaseInt := 1

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "1"

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_02"

	originalUpperLimitBaseInt := 5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "5"

	//
	//n := BigIntNum{}.NewIntExponent(5, 0)

	originalLowerLimitBaseInt := 0

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "0"

	//
	//lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_03"

	originalUpperLimitBaseInt := -5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "-5"

	// n := BigIntNum{}.NewIntExponent(-5, 0)

	originalLowerLimitBaseInt := 0

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "0"

	// lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	_, err = NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err = NFactorial{}.CalcFactorialValueBigIntNum(\n"+
			"  bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_04"

	originalUpperLimitBaseInt := 5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "5"

	//n := BigIntNum{}.NewIntExponent(5, 0)

	originalLowerLimitBaseInt := -2

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "-2"

	//lowerLimit := BigIntNum{}.NewIntExponent(-2, 0)

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	_, err = NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err = NFactorial{}.CalcFactorialValueBigIntNum(\n"+
			"  bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"The lower limit MUST BE Greater Than or Equal to 0\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_05"

	originalUpperLimitBaseInt := 5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "5"

	// n := BigIntNum{}.NewIntExponent(5, 0)

	originalLowerLimitBaseInt := 2

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "2"

	// lowerLimit := BigIntNum{}.NewIntExponent(2, 0)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	//expectedResultStr := "60"

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_06"

	originalUpperLimitBaseInt := 5

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "5"

	// n := BigIntNum{}.NewIntExponent(5, 0)

	originalLowerLimitBaseInt := 4

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "4"

	// lowerLimit := BigIntNum{}.NewIntExponent(4, 0)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	// expectedResultStr := "5"

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueBigIntNum_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueBigIntNum_07"

	originalUpperLimitBaseInt := 23

	originalUpperLimitExponentInt := 0

	originalUpperLimitNumStr := "23"

	// n := BigIntNum{}.NewIntExponent(23, 0)

	originalLowerLimitBaseInt := 0

	originalLowerLimitExponentInt := 0

	originalLowerLimitNumStr := "0"

	//	lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
			"expectedNumStr= '%v'\n"+
			"Error: isOk == false\n\n",
			ePrefix,
			expectedResultNumStr)

		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	// expectedResultStr := "25852016738884976640000"

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(originalUpperLimitBaseInt, originalUpperLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalUpperLimitBaseInt, originalUpperLimitExponentInt)\n"+
			"originalUpperLimitBaseInt= '%v'\n"+
			"originalUpperLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalUpperLimitBaseInt,
			originalUpperLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumUpperLimit.IsValid("Validating bigIntNumUpperLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumUpperLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumUpperLimitNumStr, err := bigIntNumUpperLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumUpperLimitNumStr != originalUpperLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumUpperLimitNumStr != originalUpperLimitNumStr\n"+
			"Expected originalUpperLimitNumStr = '%v'\n"+
			"  Actual originalUpperLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumUpperLimitNumStr, originalUpperLimitNumStr)

		return
	}

	bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(originalLowerLimitBaseInt, originalLowerLimitExponentInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimit, err := new(BigIntNum).NewIntExponent(\n"+
			"  originalLowerLimitBaseInt, originalLowerLimitExponentInt)\n"+
			"originalLowerLimitBaseInt= '%v'\n"+
			"originalLowerLimitExponentInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			originalLowerLimitBaseInt,
			originalLowerLimitExponentInt,
			err.Error())

		return
	}

	err = bigIntNumLowerLimit.IsValid("Validating bigIntNumLowerLimit")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumLowerLimit')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumLowerLimitNumStr, err := bigIntNumLowerLimit.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if bigIntNumLowerLimitNumStr != originalLowerLimitNumStr {
		t.Errorf("%v\n"+
			"Error: Number Strings DO NOT MATCH!\n"+
			"Because bigIntNumLowerLimitNumStr != originalLowerLimitNumStr\n"+
			"Expected originalLowerLimitNumStr = '%v'\n"+
			"  Actual originalLowerLimitNumStr = '%v'\n\n",
			ePrefix, bigIntNumLowerLimitNumStr, originalLowerLimitNumStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueBigIntNum(bigIntNumUpperLimit, bigIntNumLowerLimit)\n"+
			"bigIntNumUpperLimit= '%v'\n"+
			"bigIntNumLowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			bigIntNumUpperLimitNumStr,
			bigIntNumLowerLimitNumStr,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_01"

	upperLimitInt := 5

	lowerLimitInt := 1

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_02"

	upperLimitInt := 5

	lowerLimitInt := 0

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_03"

	upperLimitInt := -5

	lowerLimitInt := 0

	_, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_04"

	upperLimitInt := 5

	lowerLimitInt := -2

	_, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"The lower limit MUST BE Greater Than or Equal To 0\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_05"

	upperLimitInt := 5

	lowerLimitInt := 2

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_06"

	upperLimitInt := 5

	lowerLimitInt := 4

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt_06"

	upperLimitInt := 23

	lowerLimitInt := 0

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigInt, isOk := big.NewInt(0).SetString(expectedNumStr, 10)\n"+
			"expectedNumStr= '%v'\n"+
			"Error: isOk == false\n\n",
			ePrefix,
			expectedResultNumStr)

		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(upperLimitInt, lowerLimitInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt(\n"+
			"  upperLimitInt, lowerLimitInt)\n"+
			"upperLimitInt= '%v'\n"+
			"lowerLimitInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt,
			lowerLimitInt,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_01"

	upperLimitInt32 := int32(5)

	lowerLimitInt32 := int32(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_02"

	upperLimitInt32 := int32(5)

	lowerLimitInt32 := int32(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_03"

	upperLimitInt32 := int32(-5)

	lowerLimitInt32 := int32(0)

	_, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_04"

	upperLimitInt32 := int32(5)

	lowerLimitInt32 := int32(-2)

	_, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"The lower limit MUST BE Greater Than or Equal To Zero!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_05"

	upperLimitInt32 := int32(5)

	lowerLimitInt32 := int32(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_06"

	upperLimitInt32 := int32(5)

	lowerLimitInt32 := int32(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt32_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt32_07"

	upperLimitInt32 := int32(23)

	lowerLimitInt32 := int32(0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(upperLimitInt32, lowerLimitInt32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt32(\n"+
			"  upperLimitInt32, lowerLimitInt32)\n"+
			"upperLimitInt32= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt32,
			lowerLimitInt32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_01"

	upperLimitInt64 := int64(5)

	lowerLimitInt64 := int64(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_02"

	upperLimitInt64 := int64(5)

	lowerLimitInt64 := int64(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_03"

	upperLimitInt64 := int64(-5)

	lowerLimitInt64 := int64(0)

	_, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  bigIntUpperLimit, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt32= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_04"

	upperLimitInt64 := int64(5)

	lowerLimitInt64 := int64(-2)

	_, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"The lower limit MUST BE Greater Than or Equal To Zero!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_05"

	upperLimitInt64 := int64(5)

	lowerLimitInt64 := int64(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_06"

	upperLimitInt64 := int64(5)

	lowerLimitInt64 := int64(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueInt64_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_07"

	upperLimitInt64 := int64(23)

	lowerLimitInt64 := int64(0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(upperLimitInt64, lowerLimitInt64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueInt64(\n"+
			"  upperLimitInt64, lowerLimitInt64)\n"+
			"upperLimitInt64= '%v'\n"+
			"lowerLimitInt64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitInt64,
			lowerLimitInt64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_01"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(\n"+
			"  nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_02"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(\n"+
			"  nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_03"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(94)

	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcNFactorialValue(nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit)
		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueInt64_04"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(0)

	nFac.LowerLimit = uint64(2)

	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcNFactorialValue(nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit)
		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_05"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(\n"+
			"  nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_06"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(\n"+
			"  nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcNFactorialValue_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcNFactorialValue_07"

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(23)

	nFac.LowerLimit = uint64(1)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcNFactorialValue(\n"+
			"  nFac)\n"+
			"nFac.UpperLimit= '%v'\n"+
			"nFac.LowerLimit= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			nFac.UpperLimit,
			nFac.LowerLimit,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_01"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"  upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_02"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"  upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_03"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(9)

	_, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"    upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_04"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(32)

	_, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"    upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_05"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"  upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_06"

	upperLimitUint := uint(5)

	lowerLimitUint := uint(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"  upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint_07"

	upperLimitUint := uint(23)

	lowerLimitUint := uint(0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}
	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(upperLimitUint, lowerLimitUint)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint(\n"+
			"  upperLimitUint, lowerLimitUint)\n"+
			"upperLimitUint= '%v'\n"+
			"lowerLimitUint= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint,
			lowerLimitUint,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_01"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"  upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_02"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"  upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_03"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(9)

	_, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"    upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_04"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(32)

	_, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"    upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_05"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"  upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_06"

	upperLimitUint32 := uint32(5)

	lowerLimitUint32 := uint32(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"  upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint32_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint32_07"

	upperLimitUint32 := uint32(23)

	lowerLimitUint32 := uint32(0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(upperLimitUint32, lowerLimitUint32)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint32(\n"+
			"  upperLimitUint32, lowerLimitUint32)\n"+
			"upperLimitUint32= '%v'\n"+
			"lowerLimitUint32= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint32,
			lowerLimitUint32,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_01(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_01"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(1)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"  upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_02(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_02"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(0)

	expectedResultNumStr := "120"

	expectedResultBigInt := big.NewInt(int64(120))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"  upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_03(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_03"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(9)

	_, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"    upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_04(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_04"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(32)

	_, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
			"Function Call:\n"+
			" _, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"    upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"The lower limit MUST BE LESS Than The Upper Limit!\n"+
			"An error should have been generated!\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64)
		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_05(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_05"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(2)

	expectedResultNumStr := "60"

	expectedResultBigInt := big.NewInt(int64(60))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"  upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_06(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_06"

	upperLimitUint64 := uint64(5)

	lowerLimitUint64 := uint64(4)

	expectedResultNumStr := "5"

	expectedResultBigInt := big.NewInt(int64(5))

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"  upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}

func TestNFactorial_CalcFactorialValueUint64_07(t *testing.T) {

	ePrefix := "TestNFactorial_CalcFactorialValueUint64_07"

	upperLimitUint64 := uint64(23)

	lowerLimitUint64 := uint64(0)

	expectedResultNumStr := "25852016738884976640000"

	expectedResultBigInt, isOk := big.NewInt(0).SetString(expectedResultNumStr, 10)

	if !isOk {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedResultBigInt, isOk := big.NewInt(0).\n"+
			"SetString(expectedResultNumStr, 10)\n"+
			"expectedResultNumStr= '%v'\n"+
			"Error: isOk= 'false'\n\n",
			ePrefix, expectedResultNumStr)
		return
	}

	expectedResultPrecisionInt := 0

	expectedResultPrecisionUint := uint(expectedResultPrecisionInt)

	expectedResultSignVal := 1

	expectedResultNumSeps := new(NumericSeparatorDto).NewUSADefaults()

	expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(expectedResultBigInt, expectedResultPrecisionInt)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNum, err := new(BigIntNum).NewBigIntExponent(\n"+
			"  expectedResultBigInt, expectedResultPrecisionInt)\n"+
			"expectedResultBigInt= '%v'\n"+
			"expectedResultPrecisionInt= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedResultBigInt.Text(10),
			expectedResultPrecisionInt,
			err.Error())

		return
	}

	err = expectedBigIntNum.IsValid("Validating expectedBigIntNum")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating expectedBigIntNum')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedBigIntNumNumberStr, err := expectedBigIntNum.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResultNumStr != expectedBigIntNumNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected Big Int Number Strings DO NOT MATCH!\n"+
			"Because!!!!\n"+
			"Expected expectedBigIntNumNumberStr = '%v'\n"+
			"  Actual expectedBigIntNumNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, expectedBigIntNumNumberStr)

		return
	}

	bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(upperLimitUint64, lowerLimitUint64)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResult, err := NFactorial{}.CalcFactorialValueUint64(\n"+
			"  upperLimitUint64, lowerLimitUint64)\n"+
			"upperLimitUint64= '%v'\n"+
			"lowerLimitUint64= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			upperLimitUint64,
			lowerLimitUint64,
			err.Error())

		return
	}

	err = bigIntNumResult.IsValid("Validating bigIntNumResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = bINum.IsValid('Validating bigIntNumResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumberStr, err := bigIntNumResult.GetNumStr()\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultBigInt, err := bigIntNumResult.GetBigInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionInt, err := bigIntNumResult.GetPrecisionInt()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionInt, err := \n"+
			" bigIntNumResult.GetPrecisionInt()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultPrecisionUint, err := bigIntNumResult.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultPrecisionUint, err :=\n"+
			"  bigIntNumResult.GetPrecisionUint()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultSignValue, err := bigIntNumResult.GetSign()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			" bigIntNumResultSignValue, err := bigIntNumResult.GetSign()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	bigIntNumResultNumSeps, err := bigIntNumResult.GetNumericSeparatorsDto()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"bigIntNumResultNumSeps, err := \n"+
			" bigIntNumResult.GetNumericSeparatorsDto()\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n", ePrefix, bigIntNumResultNumberStr, err.Error())
		return
	}

	if expectedResultNumStr != bigIntNumResultNumberStr {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Number Strings ARE NOT EQUAL!\n"+
			"Because expectedResultNumStr != bigIntNumResultNumberStr\n"+
			"Expected bigIntNumResultNumberStr = '%v'\n"+
			"  Actual bigIntNumResultNumberStr = '%v'\n\n",
			ePrefix, expectedResultNumStr, bigIntNumResultNumberStr)

		return
	}

	if expectedResultBigInt.Cmp(bigIntNumResultBigInt) != 0 {
		t.Errorf("%v\n"+
			"Error: Expecgted and Actual Big Int Values ARE NOT EQUAL!\n"+
			"Because expectedResultBigInt != bigIntNumResultBigInt\n"+
			"Expected bigIntNumResultBigInt = '%v'\n"+
			"  Actual bigIntNumResultBigInt = '%v'\n\n",
			ePrefix, expectedResultBigInt.Text(10), bigIntNumResultBigInt.Text(10))

		return
	}

	if expectedResultPrecisionInt != bigIntNumResultPrecisionInt {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Integer Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionInt != bigIntNumResultPrecisionInt\n"+
			"Expected bigIntNumResultPrecisionInt = '%v'\n"+
			"  Actual bigIntNumResultPrecisionInt = '%v'\n\n",
			ePrefix, expectedResultPrecisionInt, bigIntNumResultPrecisionInt)

		return
	}

	if expectedResultPrecisionUint != bigIntNumResultPrecisionUint {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Uint Precision values DO NOT MATCH!\n"+
			"Because expectedResultPrecisionUint != bigIntNumResultPrecisionUint\n"+
			"Expected bigIntNumResultPrecisionUint = '%v'\n"+
			"  Actual bigIntNumResultPrecisionUint = '%v'\n\n",
			ePrefix, expectedResultPrecisionUint, bigIntNumResultPrecisionUint)

		return
	}

	if expectedResultSignVal != bigIntNumResultSignValue {

		t.Errorf("%v\n"+
			"Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
			"Because expectedResultSignVal != bigIntNumResultSignValue\n"+
			"Expected bigIntNumResultSignValue = '%v'\n"+
			"  Actual bigIntNumResultSignValue = '%v'\n\n",
			ePrefix, expectedResultSignVal, bigIntNumResultSignValue)

		return
	}

	if !expectedResultNumSeps.Equal(bigIntNumResultNumSeps) {
		t.Errorf("%v\n"+
			"Error: Expected and Actual Numeric Separator Values ARE NOT Equal!\n"+
			"Because expectedResultNumSeps != bigIntNumResultNumSeps \n"+
			"Expected bigIntNumResultNumSeps = '%v'\n"+
			"  Actual bigIntNumResultNumSeps = '%v'\n\n",
			ePrefix, expectedResultNumSeps.String(), bigIntNumResultNumSeps.String())

		return
	}

	expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"expectedEqualsResultBINum, err := expectedBigIntNum.Equal(bigIntNumResult)\n"+
			"expectedBigIntNum= '%v'\n"+
			"bigIntNumResult= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			expectedBigIntNumNumberStr,
			bigIntNumResultNumberStr,
			err.Error())

		return
	}

	if !expectedEqualsResultBINum {
		t.Errorf("%v\n"+
			"Error: Expected and Result Big Int Number values ARE NOT EQUAL!\n"+
			"Because expectedEqualsResultBINum = 'false' \n"+
			"Expected bigIntNumResult = '%v'\n"+
			"  Actual bigIntNumResult = '%v'\n\n",
			ePrefix, expectedBigIntNumNumberStr, bigIntNumResultNumberStr)

		return
	}

	return
}
