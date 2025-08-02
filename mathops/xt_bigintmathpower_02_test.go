package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathPower_BigIntToNegativeIntegerPower_01(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_01"

	base := big.NewInt(525)

	basePrecision := big.NewInt(2)

	exponent := big.NewInt(-7)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(37)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.0000090967210256655561054952045247619"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_02(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_02"

	base := big.NewInt(18)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-2)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.00308641975308641975308641975309"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_03(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_03"

	base := big.NewInt(-18)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-2)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.00308641975308641975308641975309"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_04(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_04"

	base := big.NewInt(1231234)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-5)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(42)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.000000000035342478361550254485244873919253"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_05(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_05"

	base := big.NewInt(-1231234)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-5)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(42)

	//                             1         2         3         4
	//                    123456789012345678901234567890123456789012
	expectedResult := "-0.000000000035342478361550254485244873919253"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_06(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_06"

	base := big.NewInt(10052)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-91)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.62376977529181936206917481802668"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_07(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_07"

	base := big.NewInt(5)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-5)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(0)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "32000000000000000"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_08(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_08"

	base := big.NewInt(-5)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-5)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(0)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "-32000000000000000"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_09(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_09"

	base := big.NewInt(0)

	basePrecision := big.NewInt(3)

	exponent := big.NewInt(-5)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(12)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_10(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_10"

	base := big.NewInt(92)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(0)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(12)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "1"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_11(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_11"

	base := big.NewInt(92)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-1)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.01086956521739130434782608695652"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_12(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_12"

	base := big.NewInt(-92)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-1)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	//                             1         2         3         4
	//                    123456789012345678901234567890123456789012
	expectedResult := "-0.01086956521739130434782608695652"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeIntegerPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_13(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_13"

	base := big.NewInt(5)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-12)

	exponentPrecision := big.NewInt(1)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error to be returned due to fractional exponent.\n"+
			"However, NO ERROR WAS RETURNED!\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_14(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_14"

	base := big.NewInt(5)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(2)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error to be returned due to positive exponent.\n"+
			"However, NO ERROR WAS RETURNED!\n", ePrefix)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_15(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeIntegerPower_15"

	base := big.NewInt(5)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-2)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(-1)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error to be returned due to negative 'maxPrecision' value.\n"+
			"However, NO ERROR WAS RETURNED!", ePrefix)
	}

}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_01(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_01"

	base := big.NewInt(8)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-666)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.25034681392783363227080619360469"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_02(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_02"

	base := big.NewInt(37)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-325)

	exponentPrecision := big.NewInt(2)

	maxPrecision := big.NewInt(37)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.0000080046877744411952288377104402677"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_03(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_03"

	base := big.NewInt(32)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-36)

	exponentPrecision := big.NewInt(1)

	maxPrecision := big.NewInt(18)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.000003814697265625"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_04(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_04"

	base := big.NewInt(-3289)

	basePrecision := big.NewInt(2)

	exponent := big.NewInt(-36)

	exponentPrecision := big.NewInt(1)

	maxPrecision := big.NewInt(37)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.0000034559707288372642783133414450658"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_05(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_05"

	base := big.NewInt(19)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-23)

	exponentPrecision := big.NewInt(1)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.00114516144480839927662319331457"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_06(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_06"

	base := big.NewInt(190)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-234)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.29293526501657707028369875121837"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_07(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_07"

	base := big.NewInt(191)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "3.7657702658735250133238104119483"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_08(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_08"

	base := big.NewInt(0)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_09(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_09"

	base := big.NewInt(1)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1"

	result,
		resultPrecision,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"result, resultPrecision, err := new(BigIntMathPower).\n"+
			"  BigIntToNegativeFractionalPower(base, basePrecision, exponent,\n"+
			"    exponentPrecision, maxPrecision)\n"+
			"base= '%v'\n"+
			"basePrecision= '%v'\n"+
			"exponent= '%v'\n"+
			"exponentPrecision= '%v'\n"+
			"maxPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			base.Text(10),
			basePrecision.Text(10),
			exponent.Text(10),
			exponentPrecision.Text(10),
			maxPrecision.Text(10),
			err.Error())

		return
	}

	binResult, err := new(BigIntNum).NewBigIntBigPrecision(
		result, resultPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResult, err := new(BigIntNum).\n"+
			"  NewBigIntBigPrecision(result, resultPrecision)\n"+
			"result= '%v'\n"+
			"resultPrecision= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix,
			result.Text(10),
			resultPrecision.Text(10),
			err.Error())

		return
	}

	err = binResult.IsValid("Validating binResult")

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = binResult.IsValid('Validating binResult')\n"+
			"Error= '%v'\n\n", ePrefix, err.Error())
		return
	}

	binResultNumStr, err := binResult.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"binResultNumStr, err := binResult.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedResult != binResultNumStr {
		t.Errorf("%v\n"+
			"Error: Number String Values Not Equal\n"+
			"Because expectedResult != binResultNumStr \n"+
			"Expected binResultNumStr = '%v'\n"+
			"  Actual binResultNumStr = '%v'\n\n",
			ePrefix, expectedResult, binResultNumStr)

		return
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_10(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_09"

	base := big.NewInt(-1)

	basePrecision := big.NewInt(0)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from base= -1.\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_11(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_11"

	base := big.NewInt(191)

	basePrecision := big.NewInt(-1)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from basePrecision= -1.\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_12(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_12"

	base := big.NewInt(191)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(-1)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from exponentPrecision= -1.\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_13(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_13"

	base := big.NewInt(191)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(-335)

	exponentPrecision := big.NewInt(0)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from exponentPrecision= 0\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_14(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_14"

	base := big.NewInt(191)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(335)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from exponent= positive value\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_15(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_15"

	base := big.NewInt(191)

	basePrecision := big.NewInt(4)

	exponent := big.NewInt(1)

	exponentPrecision := big.NewInt(3)

	maxPrecision := big.NewInt(32)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from exponent= Zero\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_16(t *testing.T) {

	ePrefix := "TestBigIntMathPower_BigIntToNegativeFractionalPower_16"

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(-1)

	_,
		_,
		err := new(BigIntMathPower).BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected error return from maxPrecision= negative value\n"+
			"However, no error was returned.\n", ePrefix)
	}

	return
}
