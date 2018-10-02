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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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

	eNum, err := BigIntMathLogarithms{}.GetEulersNumberE1050()

	if err != nil {
		t.Errorf("Error returned by " +
			"BigIntMathLogarithms{}.GetEulersNumberE1050(). Error='%v'",
			err.Error())
	}

	base, err := eNum.GetBigInt()

	if err != nil {
		t.Errorf("Error returned by " +
			"eNum.GetBigInt(). Error='%v'",
			err.Error())
	}

	basePrecision := eNum.GetPrecisionUint()
	exponent := uint(9)
	internalMaxPrecision := uint((9*basePrecision) + 10)
	outputMaxPrecision := uint(28)
	expectedResult := "8103.0839275753840077099966894328"
	expectedPrecision := outputMaxPrecision

	baseToPwr, baseToPwrPrecision := BigIntMathPower{}.BigIntPwr(
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
