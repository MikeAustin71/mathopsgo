package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathPower_BigIntPwrIteration_01(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_01"

	base := big.NewInt(-2123456789012)

	basePrecision := uint(12)

	exponent := uint(9)

	internalMaxPrecision := uint(14)

	outputMaxPrecision := uint(5)

	expectedResult := "-877.79045"

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if outputMaxPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, outputMaxPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_02(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_02"

	base := big.NewInt(2123456789012)

	basePrecision := uint(12)

	exponent := uint(9)

	internalMaxPrecision := uint(14)

	outputMaxPrecision := uint(5)

	expectedResult := "877.79045"

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if outputMaxPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, outputMaxPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_03(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_03"

	base := big.NewInt(5)

	basePrecision := uint(0)

	exponent := uint(9)

	internalMaxPrecision := uint(14)

	outputMaxPrecision := uint(5)

	expectedResult := "1953125"

	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_04(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_04"

	base := big.NewInt(123456789)

	basePrecision := uint(9)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(40)

	expectedResult := "0.0000000066624627597199420074400375313628"

	expectedPrecision := uint(40)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_05(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_05"

	base := big.NewInt(-123456789)

	basePrecision := uint(9)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(40)

	expectedResult := "-0.0000000066624627597199420074400375313628"

	expectedPrecision := uint(40)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_06(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_06"

	base := big.NewInt(10)

	basePrecision := uint(0)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(40)

	expectedResult := "1000000000"

	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_07(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_07"

	base := big.NewInt(-10)

	basePrecision := uint(0)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(40)

	expectedResult := "-1000000000"

	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_08(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_08"

	base := big.NewInt(91)

	basePrecision := uint(1)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(9)

	expectedResult := "427929800.129788411"

	expectedPrecision := uint(9)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_09(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_09"

	base := big.NewInt(-91)

	basePrecision := uint(1)

	exponent := uint(9)

	internalMaxPrecision := uint(80)

	outputMaxPrecision := uint(9)

	expectedResult := "-427929800.129788411"

	expectedPrecision := uint(9)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntPwrIteration_10(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntPwrIteration_10"

	eNum := eulersNumber1k

	base := eNum.GetInteger()

	basePrecision := eNum.GetPrecisionUint()

	exponent := uint(9)

	internalMaxPrecision := (9 * basePrecision) + 10

	outputMaxPrecision := uint(28)

	expectedResult := "8103.0839275753840077099966894328"

	expectedPrecision := outputMaxPrecision

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntPwrIteration(\n"+
			" base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"internalMaxPrecision= '%v'\n"+
			"outputMaxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision,
			exponent,
			internalMaxPrecision,
			outputMaxPrecision,
			err.Error())

		return
	}

	result, err := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, err := new(BigIntNum).NewBigInt(\n"+
			"  baseToPwr, baseToPwrPrecision)\n"+
			"baseToPwr= '%v'\n"+
			"baseToPwrPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			baseToPwr.Text(10),
			baseToPwrPrecision,
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

	resultPrecisionUint, err := result.GetPrecisionUint()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultPrecisionUint, err := result.GetPrecisionUint()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != resultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != resultNumStr \n"+
			"Expected resultNumStr = '%v'\n"+
			"  Actual resultNumStr = '%v'\n\n",
			ePrefix, expectedResult, resultNumStr)

		return
	}

	if expectedPrecision != resultPrecisionUint {
		t.Errorf("%v\n"+
			"Error: Result Precision Invalid!\n"+
			"Because outputMaxPrecision != resultPrecisionUint \n"+
			"Expected resultPrecisionUint = '%v'\n"+
			"  Actual resultPrecisionUint = '%v'\n\n",
			ePrefix, expectedPrecision, resultPrecisionUint)

		return
	}

	return
}

func TestBigIntMathPower_BigIntegerPwrIteration_01(t *testing.T) {

	base := big.NewInt(-2123456789012)
	basePrecision := big.NewInt(12)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(14)
	outputMaxPrecision := big.NewInt(5)
	expectedResult := "-877.79045"

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if outputMaxPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			outputMaxPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_02(t *testing.T) {

	base := big.NewInt(2123456789012)
	basePrecision := big.NewInt(12)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(14)
	outputMaxPrecision := big.NewInt(5)
	expectedResult := "877.79045"
	expectedPrecision := big.NewInt(5)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_03(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(14)
	outputMaxPrecision := big.NewInt(5)
	expectedResult := "1953125"
	expectedPrecision := big.NewInt(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_04(t *testing.T) {

	base := big.NewInt(123456789)
	basePrecision := big.NewInt(9)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(40)
	expectedResult := "0.0000000066624627597199420074400375313628"
	expectedPrecision := big.NewInt(40)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_05(t *testing.T) {

	base := big.NewInt(-123456789)
	basePrecision := big.NewInt(9)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(40)
	expectedResult := "-0.0000000066624627597199420074400375313628"
	expectedPrecision := big.NewInt(40)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_06(t *testing.T) {

	base := big.NewInt(10)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(40)
	expectedResult := "1000000000"
	expectedPrecision := big.NewInt(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_07(t *testing.T) {

	base := big.NewInt(-10)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(40)
	expectedResult := "-1000000000"
	expectedPrecision := big.NewInt(0)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_08(t *testing.T) {

	base := big.NewInt(91)
	basePrecision := big.NewInt(1)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(9)
	expectedResult := "427929800.129788411"
	expectedPrecision := big.NewInt(9)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_09(t *testing.T) {

	base := big.NewInt(-91)
	basePrecision := big.NewInt(1)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(80)
	outputMaxPrecision := big.NewInt(9)
	expectedResult := "-427929800.129788411"
	expectedPrecision := big.NewInt(9)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}
}

func TestBigIntMathPower_BigIntegerPwrIteration_10(t *testing.T) {

	eNum := eulersNumber1k

	base := eNum.GetInteger()

	basePrecision := eNum.GetPrecisionBigInt()
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(0).Mul(big.NewInt(9), basePrecision)
	internalMaxPrecision.Add(internalMaxPrecision, big.NewInt(10))
	outputMaxPrecision := big.NewInt(28)
	expectedResult := "8103.0839275753840077099966894328"
	expectedPrecision := big.NewInt(0).Set(outputMaxPrecision)

	baseToPwr, baseToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntMathPower).BigIntegerPwrIteration(...) "+
			"Error='%v' ", err.Error())
	}

	result, err := new(BigIntNum).NewBigIntBigPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr)
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint())
	}

}
