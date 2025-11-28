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

	expectedBigInt := big.NewInt(int64(25852016738884976640000))

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

	n := BigIntNum{}.NewIntExponent(5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(1, 0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_02(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_03(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(-5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	_, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigIntNum(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_04(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(-2, 0)

	_, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigIntNum(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_05(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(2, 0)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_06(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(5, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(4, 0)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigIntNum_07(t *testing.T) {

	n := BigIntNum{}.NewIntExponent(23, 0)

	lowerLimit := BigIntNum{}.NewIntExponent(0, 0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigIntNum(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_01(t *testing.T) {

	n := 5

	lowerLimit := 1

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_02(t *testing.T) {

	n := 5

	lowerLimit := 0

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_03(t *testing.T) {

	n := -5

	lowerLimit := 0

	_, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt_04(t *testing.T) {

	n := 5

	lowerLimit := -2

	_, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt_05(t *testing.T) {

	n := 5

	lowerLimit := 2

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_06(t *testing.T) {

	n := 5

	lowerLimit := 4

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_07(t *testing.T) {

	n := 23

	lowerLimit := 0

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_01(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_02(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_03(t *testing.T) {

	n := int32(-5)

	lowerLimit := int32(0)

	_, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt32(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt32_04(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(-2)

	_, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt32(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt32_05(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_06(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_07(t *testing.T) {

	n := int32(23)

	lowerLimit := int32(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_01(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_02(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_03(t *testing.T) {

	n := int64(-5)

	lowerLimit := int64(0)

	_, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt64(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt64_04(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(-2)

	_, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt64(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt64_05(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_06(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_07(t *testing.T) {

	n := int64(23)

	lowerLimit := int64(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_01(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_02(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_03(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(94)

	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Error("Expected error return from CalcNFactorialValue(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcNFactorialValue_04(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(0)

	nFac.LowerLimit = uint64(2)

	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Error("Expected error return from CalcNFactorialValue(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcNFactorialValue_05(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(nFac). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_06(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(nFac). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_07(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(23)

	nFac.LowerLimit = uint64(1)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_01(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_02(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_03(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(9)

	_, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint_04(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(32)

	_, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint_05(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_06(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_07(t *testing.T) {

	n := uint(23)

	lowerLimit := uint(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_01(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_02(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_03(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(9)

	_, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint32(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint32_04(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(32)

	_, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint32(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint32_05(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_06(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_07(t *testing.T) {

	n := uint32(23)

	lowerLimit := uint32(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_01(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_02(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_03(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(9)

	_, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint64(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint64_04(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(32)

	_, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint64(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint64_05(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_06(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_07(t *testing.T) {

	n := uint64(23)

	lowerLimit := uint64(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). "+
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}
