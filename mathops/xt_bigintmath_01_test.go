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
	_, err := BigIntMath{}.GetMagnitude(target)

	if err == nil {
		t.Error("Expected an Error return from  BigIntMath{}.GetMagnitudeDigits(target). " +
			"NO ERROR WAS TRIGGERED! ")
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
