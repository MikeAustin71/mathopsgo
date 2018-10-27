package mathops

import (
	"math/big"
	"testing"
)


func TestBigIntMathPower_BigIntToPositiveIntegerPower_01(t *testing.T) {

	base := big.NewInt(-525)
	basePrecision := big.NewInt(2)
	exponent := big.NewInt(7)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(14)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_02(t *testing.T) {

	base := big.NewInt(525)
	basePrecision := big.NewInt(2)
	exponent := big.NewInt(8)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(16)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_03(t *testing.T) {

	base := big.NewInt(1123456)
	basePrecision := big.NewInt(6)
	exponent := big.NewInt(51)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(28)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_04(t *testing.T) {

	base := big.NewInt(-1123456)
	basePrecision := big.NewInt(6)
	exponent := big.NewInt(51)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(28)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_05(t *testing.T) {

	base := big.NewInt(2)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(100)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_06(t *testing.T) {

	base := big.NewInt(2345)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(25)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(47)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_07(t *testing.T) {

	base := big.NewInt(-2345)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(25)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(47)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_08(t *testing.T) {

	base := big.NewInt(0)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(5)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_09(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(0)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

	binResult,err := BigIntNum{}.NewBigIntPrecision(result, resultPrecision)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_10(t *testing.T) {

	base := big.NewInt(-5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(0)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_11(t *testing.T) {

	base := big.NewInt(91)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(1)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_12(t *testing.T) {

	base := big.NewInt(-91)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(1)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(0)

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

func TestBigIntMathPower_BigIntToPositiveIntegerPower_13(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(22)
	exponentPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(5)

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
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(-2)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(5)

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
	basePrecision := big.NewInt(2)
	exponent := big.NewInt(-7)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(37)

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
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_02(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_03(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_04(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_05(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_06(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_07(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_08(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_09(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_10(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_11(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_12(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeIntegerPower() " +
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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_13(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(-12)
	exponentPrecision := big.NewInt(1)
	maxPrecision := big.NewInt(32)

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
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(2)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

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
