package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathPower_BigIntPwr_01(t *testing.T) {

	base := big.NewInt(-2123456789012)
	basePrecision := uint(12)
	exponent := uint(9)
	internalMaxPrecision := uint(14)
	outputMaxPrecision := uint(5)
	expectedResult := "-877.79045"

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if outputMaxPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			outputMaxPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_02(t *testing.T) {

	base := big.NewInt(2123456789012)
	basePrecision := uint(12)
	exponent := uint(9)
	internalMaxPrecision := uint(14)
	outputMaxPrecision := uint(5)
	expectedResult := "877.79045"

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if outputMaxPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			outputMaxPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_03(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := uint(9)
	internalMaxPrecision := uint(14)
	outputMaxPrecision := uint(5)
	expectedResult := "1953125"
	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_04(t *testing.T) {

	base := big.NewInt(123456789)
	basePrecision := uint(9)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(40)
	expectedResult := "0.0000000066624627597199420074400375313628"
	expectedPrecision := uint(40)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_05(t *testing.T) {

	base := big.NewInt(-123456789)
	basePrecision := uint(9)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(40)
	expectedResult := "-0.0000000066624627597199420074400375313628"
	expectedPrecision := uint(40)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_06(t *testing.T) {

	base := big.NewInt(10)
	basePrecision := uint(0)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(40)
	expectedResult := "1000000000"
	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_07(t *testing.T) {

	base := big.NewInt(-10)
	basePrecision := uint(0)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(40)
	expectedResult := "-1000000000"
	expectedPrecision := uint(0)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_08(t *testing.T) {

	base := big.NewInt(91)
	basePrecision := uint(1)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(9)
	expectedResult := "427929800.129788411"
	expectedPrecision := uint(9)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_09(t *testing.T) {

	base := big.NewInt(-91)
	basePrecision := uint(1)
	exponent := uint(9)
	internalMaxPrecision := uint(80)
	outputMaxPrecision := uint(9)
	expectedResult := "-427929800.129788411"
	expectedPrecision := uint(9)

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntPwr_10(t *testing.T) {

	eNum := eulersNumber1050

	base := eNum.GetInteger()

	basePrecision := eNum.GetPrecision()
	exponent := uint(9)
	internalMaxPrecision := uint((9*basePrecision) + 10)
	outputMaxPrecision := uint(28)
	expectedResult := "8103.0839275753840077099966894328"
	expectedPrecision := outputMaxPrecision

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	result := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision != result.GetPrecisionUint() {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision, result.GetPrecisionUint() )
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_01(t *testing.T) {

	base := big.NewInt(82)
	basePrecision := big.NewInt(2)
	exponent := big.NewInt(25)
	exponentPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.60888409668835989397082286114801"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveFractionalPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_02(t *testing.T) {

	base := big.NewInt(987654)
	basePrecision := big.NewInt(5)
	exponent := big.NewInt(25)
	exponentPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1.7727651549444647752885992938899"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_03(t *testing.T) {

	base := big.NewInt(987654)
	basePrecision := big.NewInt(2)
	exponent := big.NewInt(333)
	exponentPrecision := big.NewInt(2)
	maxPrecision := big.NewInt(18)

	//                                         1         2         3
	//                                1234567890123456789012345678901234567
	expectedResult := "20046293000573.168546517448618919"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_04(t *testing.T) {

	base := big.NewInt(54)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(333)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "3.7747406845455202662527585226808"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveFractionalPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_05(t *testing.T) {

	base := big.NewInt(1)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(333)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveFractionalPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_06(t *testing.T) {

	base := big.NewInt(-1)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(333)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return with radicand==-1. NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_07(t *testing.T) {

	base := big.NewInt(0)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(333)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveFractionalPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_08(t *testing.T) {

	base := big.NewInt(256)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(0)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "1"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveFractionalPower(...) " +
			"Error='%v' ", err.Error())
	}

	binResult, err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_09(t *testing.T) {

	base := big.NewInt(256)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(234)
	exponentPrecision := big.NewInt(-3)
	maxPrecision := big.NewInt(31)

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return with exponentPrecision==-3. NO ERROR RETURNED!")	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_10(t *testing.T) {

	base := big.NewInt(256)
	basePrecision := big.NewInt(-1)
	exponent := big.NewInt(222)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(31)

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return with basePrecision==-1. NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_11(t *testing.T) {

	base := big.NewInt(256)
	basePrecision := big.NewInt(1)
	exponent := big.NewInt(222)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(31)

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return with exponentPrecision==0. NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToPositiveFractionalPower_12(t *testing.T) {

	base := big.NewInt(256)
	basePrecision := big.NewInt(1)
	exponent := big.NewInt(-222)
	exponentPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(31)

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return with exponent==-22.2. NO ERROR RETURNED!")
	}

}
