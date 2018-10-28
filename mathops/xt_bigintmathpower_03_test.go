package mathops

import (
	"math/big"
	"testing"
)


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


func TestBigIntMathPower_MinimumRequiredPrecision_01(t *testing.T) {

	base := BigIntNum{}.NewInt(312,2)
	exponent := BigIntNum{}.NewInt(4, 0)
	expectedResult := uint(8)


	result, err := BigIntMathPower{}.MinimumRequiredPrecision(base, exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}." +
			"MinimumRequiredPrecision(base, exponent)" +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestBigIntMathPower_MinimumRequiredPrecision_02(t *testing.T) {

	base := BigIntNum{}.NewInt(312345,5)
	exponent := BigIntNum{}.NewInt(18, 0)
	expectedResult := uint(90)


	result, err := BigIntMathPower{}.MinimumRequiredPrecision(base, exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}." +
			"MinimumRequiredPrecision(base, exponent)" +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestBigIntMathPower_MinimumRequiredPrecision_03(t *testing.T) {

	base := BigIntNum{}.NewInt(-312345,5)
	exponent := BigIntNum{}.NewInt(18, 0)
	expectedResult := uint(90)


	result, err := BigIntMathPower{}.MinimumRequiredPrecision(base, exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}." +
			"MinimumRequiredPrecision(base, exponent)" +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestBigIntMathPower_MinimumRequiredPrecision_04(t *testing.T) {

	base := BigIntNum{}.NewInt(312345,5)
	exponent := BigIntNum{}.NewInt(-18, 0)
	expectedResult := uint(90)


	result, err := BigIntMathPower{}.MinimumRequiredPrecision(base, exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}." +
			"MinimumRequiredPrecision(base, exponent)" +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestBigIntMathPower_MinimumRequiredPrecision_05(t *testing.T) {

	base := BigIntNum{}.NewInt(312345,5)
	exponent := BigIntNum{}.NewUint64(12345678901234567890, 0)

	result, err := BigIntMathPower{}.MinimumRequiredPrecision(base, exponent)

	if err == nil {
		t.Error("Error: Expected error be returned. NO ERROR RETURNED!")
	}

	if result != uint(4294967295) {
		t.Errorf("Error: Expected result='4294967295'. Instead, result='%v'",
			result)
	}

}

func TestBigIntMathPower_Pwr_01(t *testing.T) {
	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds
	baseStr := "2"
	exponentStr := "4"
	expectedNumStr := "16"
	maxPrecision := uint(17)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_02(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds
	baseStr := "2"
	exponentStr := "-4"
	expectedNumStr := "0.0625"
	maxPrecision := uint(17)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_03(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "37.241"
	exponentStr := "8"
	expectedNumStr := "3699735472699.4101912057680101525"
	maxPrecision := uint(19)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_04(t *testing.T) {

	// Time:
	// 0-Milliseconds 999-Microseconds 600-Nanoseconds

	baseStr := "37"
	exponentStr := "3.25"
	expectedNumStr := "124926.79641959048051506133768818"
	maxPrecision := uint(26)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_05(t *testing.T) {

	// Time:
	// 1-Milliseconds 999-Microseconds 400-Nanoseconds

	baseStr := "37"
	exponentStr := "-3.25"
	expectedNumStr := "0.0000080046877744411952288377104402677"
	maxPrecision := uint(37)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_06(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-37"
	exponentStr := "-3"
	expectedNumStr := "-0.000019742167295125658894833474818866"
	maxPrecision := uint(36)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_07(t *testing.T) {

	// Time:
	// 0-Milliseconds 999-Microseconds 300-Nanoseconds

	baseStr := "32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := uint(18)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_08(t *testing.T) {

	// Time:
	// 0-Milliseconds 999-Microseconds 300-Nanoseconds

	baseStr := "-32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := uint(18)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_09(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "5"
	exponentStr := "-3"
	expectedNumStr := "0.008"
	maxPrecision := uint(3)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_10(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-5"
	exponentStr := "4"
	expectedNumStr := "625"
	maxPrecision := uint(0)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_11(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-5"
	exponentStr := "5"
	expectedNumStr := "-3125"
	maxPrecision := uint(0)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_12(t *testing.T) {

	// Time:
	// 0-Milliseconds 999-Microseconds 300-Nanoseconds

	baseStr := "4"
	exponentStr := "0.25"
	expectedNumStr := "1.4142135623730950488016887242097"
	maxPrecision := uint(31)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_13(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45"
	exponentStr := "120"
	expectedNumStr := "2429414689006507011047680668198610544614376056243160093821961151708217872022541081600097552192834563575596196518385368260399467853246465311637303360756939677768395657864175518625415861606597900390625"
	maxPrecision := uint(200)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_14(t *testing.T) {

	// Time:
	// 1-Milliseconds 0-Microseconds 400-Nanoseconds

	baseStr := "19"
	exponentStr := "2.3"
	//                     12345648901234567890123456789
	expectedNumStr := "873.23931881701910176203214553167"
	maxPrecision := uint(29)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_15(t *testing.T) {

	// Time:
	// 1-Milliseconds 998-Microseconds 900-Nanoseconds

	baseStr := "19"
	exponentStr := "-2.3"
	//                   12345648901234567890123456789012
	expectedNumStr := "0.00114516144480839927662319331457"
	maxPrecision := uint(32)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_16(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45"
	exponentStr := "-2"
	//                      12345678901234567890123456789012345
	expectedNumStr := "0.00049382716049382716049382716049383"
	maxPrecision := uint(35)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_17(t *testing.T) {

	baseStr := "0"
	exponentStr := "5"
	expectedNumStr := "0"
	maxPrecision := uint(35)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)"+
			"Error='%v' \n", err.Error())
	}

	actualNumStr := result.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntMathPower_Pwr_18(t *testing.T) {

	// Time:
	//  0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45.1"
	exponentStr := "0"
	//                      12345678901234567890123456789012345
	expectedNumStr := "1"
	maxPrecision := uint(5)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_19(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45.1"
	exponentStr := "1"
	//                      12345678901234567890123456789012345
	expectedNumStr := "45.1"
	maxPrecision := uint(5)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_20(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-45.1"
	exponentStr := "1"
	//                      12345678901234567890123456789012345
	expectedNumStr := "-45.1"
	maxPrecision := uint(5)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_21(t *testing.T) {

	// Time:
	// 1-Milliseconds 998-Microseconds 800-Nanoseconds

	baseStr := "-45.6"
	exponentStr := "-3.2"
	//                      12345678901234567890123456789012345
	//                    0.00000491261243811417457984700270545
	// 4.91261243811417457984700270545e-6

	expectedNumStr := "0.00000491261243811417457984700270545"
	maxPrecision := uint(35)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_22(t *testing.T) {

	baseStr := "-45.632"
	exponentStr := "-1.01579"
	maxPrecision := uint(15)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	_, err = BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err == nil {
		t.Error("Expected Error to be returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"NO ERROR WAS RETURNED! \n")
	}

}

func TestBigIntMathPower_Pwr_23(t *testing.T) {
	// Time:
	// 1-Milliseconds 998-Microseconds 800-Nanoseconds

	baseStr := "2.125"
	exponentStr := "-5"
	//                   12345678901234567890123456789012
	expectedNumStr := "0.02307838042845159759046157465153"

	maxPrecision := uint(32)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). "+
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}
