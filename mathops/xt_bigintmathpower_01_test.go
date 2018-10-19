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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_01(t *testing.T) {

	base := big.NewInt(-525)
	basePrecision := uint(2)
	exponent := big.NewInt(7)
	exponentPrecision := uint(0)
	maxPrecision := uint(14)

	//                                  1
	//                         1234567890123456
	expectedResult := "-109929.72052001953125"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_02(t *testing.T) {

	base := big.NewInt(-525)
	basePrecision := uint(2)
	exponent := big.NewInt(8)
	exponentPrecision := uint(0)
	maxPrecision := uint(16)

	//                                 1
	//                        1234567890123456
	expectedResult := "577131.0327301025390625"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_03(t *testing.T) {

	base := big.NewInt(1123456)
	basePrecision := uint(6)
	exponent := big.NewInt(51)
	exponentPrecision := uint(0)
	maxPrecision := uint(28)

	//                              1         2
	//                     1234567890123456789012345678
	expectedResult := "378.7559536547494902948952952204"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_04(t *testing.T) {

	base := big.NewInt(-1123456)
	basePrecision := uint(6)
	exponent := big.NewInt(51)
	exponentPrecision := uint(0)
	maxPrecision := uint(28)

	//                               1         2
	//                      1234567890123456789012345678
	expectedResult := "-378.7559536547494902948952952204"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_05(t *testing.T) {

	base := big.NewInt(2)
	basePrecision := uint(0)
	exponent := big.NewInt(100)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                               1         2
	//                      1234567890123456789012345678
	// xx
	expectedResult := "1267650600228229401496703205376"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_06(t *testing.T) {

	base := big.NewInt(2345)
	basePrecision := uint(4)
	exponent := big.NewInt(25)
	exponentPrecision := uint(0)
	maxPrecision := uint(47)

	//                            1         2         3         4
	//                   12345678901234567890123456789012345678901234567
	expectedResult := "0.00000000000000017929623758795375247136216776668"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_07(t *testing.T) {

	base := big.NewInt(-2345)
	basePrecision := uint(4)
	exponent := big.NewInt(25)
	exponentPrecision := uint(0)
	maxPrecision := uint(47)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "-0.00000000000000017929623758795375247136216776668"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_08(t *testing.T) {

	base := big.NewInt(0)
	basePrecision := uint(4)
	exponent := big.NewInt(5)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "0"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_09(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := big.NewInt(0)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "1"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_10(t *testing.T) {

	base := big.NewInt(-5)
	basePrecision := uint(0)
	exponent := big.NewInt(0)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "1"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_11(t *testing.T) {

	base := big.NewInt(91)
	basePrecision := uint(0)
	exponent := big.NewInt(1)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "91"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_12(t *testing.T) {

	base := big.NewInt(-91)
	basePrecision := uint(0)
	exponent := big.NewInt(1)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567
	expectedResult := "-91"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_13(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := big.NewInt(22)
	exponentPrecision := uint(1)
	maxPrecision := uint(5)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error to be returned due to fractional exponent " +
			"NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToPositiveIntegerPower_14(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := big.NewInt(-2)
	exponentPrecision := uint(0)
	maxPrecision := uint(5)

	//                             1         2         3         4
	//                    12345678901234567890123456789012345678901234567

	_,
	_,
	err := BigIntMathPower{}.BigIntToPositiveIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error to be returned due to negative exponent " +
			"NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_01(t *testing.T) {

	base := big.NewInt(525)
	basePrecision := uint(2)
	exponent := big.NewInt(-7)
	exponentPrecision := uint(0)
	maxPrecision := uint(37)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.0000090967210256655561054952045247619"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_02(t *testing.T) {

	base := big.NewInt(18)
	basePrecision := uint(0)
	exponent := big.NewInt(-2)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.00308641975308641975308641975309"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_03(t *testing.T) {

	base := big.NewInt(-18)
	basePrecision := uint(0)
	exponent := big.NewInt(-2)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	//                            1         2         3
	//                   1234567890123456789012345678901234567
	expectedResult := "0.00308641975308641975308641975309"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_04(t *testing.T) {

	base := big.NewInt(1231234)
	basePrecision := uint(4)
	exponent := big.NewInt(-5)
	exponentPrecision := uint(0)
	maxPrecision := uint(42)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.000000000035342478361550254485244873919253"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_05(t *testing.T) {

	base := big.NewInt(-1231234)
	basePrecision := uint(4)
	exponent := big.NewInt(-5)
	exponentPrecision := uint(0)
	maxPrecision := uint(42)

	//                             1         2         3         4
	//                    123456789012345678901234567890123456789012
	expectedResult := "-0.000000000035342478361550254485244873919253"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_06(t *testing.T) {

	base := big.NewInt(10052)
	basePrecision := uint(4)
	exponent := big.NewInt(-91)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.62376977529181936206917481802668"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_07(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(4)
	exponent := big.NewInt(-5)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "32000000000000000"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_08(t *testing.T) {

	base := big.NewInt(-5)
	basePrecision := uint(4)
	exponent := big.NewInt(-5)
	exponentPrecision := uint(0)
	maxPrecision := uint(0)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "-32000000000000000"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_09(t *testing.T) {

	base := big.NewInt(0)
	basePrecision := uint(3)
	exponent := big.NewInt(-5)
	exponentPrecision := uint(0)
	maxPrecision := uint(12)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_10(t *testing.T) {

	base := big.NewInt(92)
	basePrecision := uint(0)
	exponent := big.NewInt(0)
	exponentPrecision := uint(0)
	maxPrecision := uint(12)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "1"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_11(t *testing.T) {

	base := big.NewInt(92)
	basePrecision := uint(0)
	exponent := big.NewInt(-1)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	//                            1         2         3         4
	//                   123456789012345678901234567890123456789012
	expectedResult := "0.01086956521739130434782608695652"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_12(t *testing.T) {

	base := big.NewInt(-92)
	basePrecision := uint(0)
	exponent := big.NewInt(-1)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	//                             1         2         3         4
	//                    123456789012345678901234567890123456789012
	expectedResult := "-0.01086956521739130434782608695652"

	result,
	resultPrecision,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToPositiveIntegerPower() " +
			"Error='%v' ", err.Error())
	}

	binResult := BigIntNum{}.NewBigInt(result, resultPrecision)

	actualNumStr := binResult.GetNumStr()

	if expectedResult != actualNumStr {
		t.Errorf("Expected result='%v'.  Instead, result='%v'. ",
			expectedResult, actualNumStr)
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_13(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := big.NewInt(-12)
	exponentPrecision := uint(1)
	maxPrecision := uint(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error to be returned due to fractional exponent " +
			"NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToNegativeIntegerPower_14(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := uint(0)
	exponent := big.NewInt(2)
	exponentPrecision := uint(0)
	maxPrecision := uint(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error to be returned due to positive exponent " +
			"NO ERROR RETURNED!")
	}

}
