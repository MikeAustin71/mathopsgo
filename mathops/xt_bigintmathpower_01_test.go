package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMathPower_BigIntPwrIteration_01(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_02(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_03(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_04(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_05(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_06(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_07(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_08(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_09(t *testing.T) {

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

func TestBigIntMathPower_BigIntPwrIteration_10(t *testing.T) {

	eNum := eulersNumber1k

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


func TestBigIntMathPower_BigIntegerPwrIteration_01(t *testing.T) {

	base := big.NewInt(-2123456789012)
	basePrecision := big.NewInt(12)
	exponent := big.NewInt(9)
	internalMaxPrecision := big.NewInt(14)
	outputMaxPrecision := big.NewInt(5)
	expectedResult := "-877.79045"

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if outputMaxPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			outputMaxPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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


	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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


	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
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

	baseToPwr, baseToPwrPrecision, err := BigIntMathPower{}.BigIntegerPwrIteration(
		base,
		basePrecision,
		exponent,
		internalMaxPrecision,
		outputMaxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntegerPwrIteration(...) " +
			"Error='%v' ", err.Error())
	}

	result, err := BigIntNum{}.NewBigIntPrecision(baseToPwr, baseToPwrPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigIntPrecision(...) " +
			"Error='%v' ", err.Error())
	}

	resultStr := result.GetNumStr()

	if expectedResult != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, resultStr )
	}

	if expectedPrecision.Cmp(result.GetPrecisionBigInt()) != 0 {
		t.Errorf("Error: Expected result precision='%v'. Instead, result precision='%v'",
			expectedPrecision.Text(10), result.GetPrecisionUint() )
	}

}
