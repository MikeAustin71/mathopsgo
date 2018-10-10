package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntMath_GetMagnitude_01(t *testing.T) {

	target := big.NewInt(98327123)
	expectedMagnitude := big.NewInt(7)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_02(t *testing.T) {

	target := big.NewInt(0)
	expectedMagnitude := big.NewInt(0)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_03(t *testing.T) {

	target := big.NewInt(82)
	expectedMagnitude := big.NewInt(1)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_04(t *testing.T) {

	target := big.NewInt(-5)
	expectedMagnitude := big.NewInt(0)
	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}
}

func TestBigIntMath_GetMagnitude_05(t *testing.T) {

	target := big.NewInt(2)
	expectedMagnitude := big.NewInt(0)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_06(t *testing.T) {

	target := big.NewInt(85652647928)
	expectedMagnitude := big.NewInt(10)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_07(t *testing.T) {


	numStr := "8565264792812345678901234567890"
	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())
	expectedMagnitude := big.NewInt(30)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_08(t *testing.T) {

	numStr := "-8565264792812345678901234567890"
	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())
	expectedMagnitude := big.NewInt(30)

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_09(t *testing.T) {

	numStr := "131072"
	expectedMagnitude := big.NewInt(5)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_10(t *testing.T) {

	numStr := "131071"
	expectedMagnitude := big.NewInt(5)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_11(t *testing.T) {

	numStr := "999999999999"
	expectedMagnitude := big.NewInt(11)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_12(t *testing.T) {

	numStr := "1999999999999"
	expectedMagnitude := big.NewInt(12)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_13(t *testing.T) {

	numStr := "9999999999999"
	expectedMagnitude := big.NewInt(12)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_14(t *testing.T) {

	numStr := "7"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_15(t *testing.T) {

	numStr := "8"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_16(t *testing.T) {

	numStr := "9"
	expectedMagnitude := big.NewInt(0)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}

func TestBigIntMath_GetMagnitude_17(t *testing.T) {

	numStr := "10"
	expectedMagnitude := big.NewInt(1)

	fixDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	target := big.NewInt(0).Set(fixDec.GetInteger())

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		t.Errorf("Error returned by BigIntMath{}.GetMagnitudeDigits(target). "+
			"target='%v' Error='%v' ",
			target.Text(10), err.Error())
	}

	if expectedMagnitude.Cmp(magnitude) != 0 {
		t.Errorf("Error: Expected Magnitude='%v'. Instead, Magnitude='%v'.",
			expectedMagnitude.Text(10), magnitude.Text(10))
	}

}
