package mathops

import (
	"math/big"
	"testing"
)

func TestNFactorial_CalcFactorialValueBigInt_01(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_02(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_03(t *testing.T) {

	n := big.NewInt(-5)

	lowerLimit := big.NewInt(0)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=-1")

	}

}

func TestNFactorial_CalcFactorialValueBigInt_04(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(-2)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueBigInt_05(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_06(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_07(t *testing.T) {

	n := big.NewInt(23)

	lowerLimit := big.NewInt(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}
