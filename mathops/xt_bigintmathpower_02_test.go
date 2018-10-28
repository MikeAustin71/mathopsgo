package mathops

import (
	"math/big"
	"testing"
)


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

func TestBigIntMathPower_BigIntToNegativeIntegerPower_15(t *testing.T) {

	base := big.NewInt(5)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(-2)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(-1)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeIntegerPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error to be returned due to negative 'maxPrecision' value. " +
			"NO ERROR RETURNED!")
	}

}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_01(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_02(t *testing.T) {
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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_03(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_04(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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


func TestBigIntMathPower_BigIntToNegativeFractionalPower_05(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_06(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_07(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_08(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_09(t *testing.T) {

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
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.BigIntToNegativeFractionalPower() " +
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

func TestBigIntMathPower_BigIntToNegativeFractionalPower_10(t *testing.T) {

	base := big.NewInt(-1)
	basePrecision := big.NewInt(0)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from base==-1. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_11(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(-1)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from basePrecision==-1. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_12(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(-1)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from exponentPrecision==-1. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_13(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(0)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from exponentPrecision==0. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_14(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(335)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from exponent==positive value. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_15(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(1)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(32)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from exponent==Zero. No Error returned! ")
	}
}

func TestBigIntMathPower_BigIntToNegativeFractionalPower_16(t *testing.T) {

	base := big.NewInt(191)
	basePrecision := big.NewInt(4)
	exponent := big.NewInt(-335)
	exponentPrecision := big.NewInt(3)
	maxPrecision := big.NewInt(-1)

	_,
	_,
	err := BigIntMathPower{}.BigIntToNegativeFractionalPower(
		base,
		basePrecision,
		exponent,
		exponentPrecision,
		maxPrecision)

	if err == nil {
		t.Error("Error: Expected error return from maxPrecision== negative value. No Error returned! ")
	}
}

