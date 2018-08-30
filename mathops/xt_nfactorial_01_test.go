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
