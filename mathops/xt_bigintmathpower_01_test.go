package mathops

import (
	"testing"
)

func TestBigIntMathPower_Pwr_01(t *testing.T) {
	baseStr := "2"
	exponentStr := "4"
	expectedNumStr := "16"
	maxPrecision := uint(17)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}

func TestBigIntMathPower_Pwr_02(t *testing.T) {
	baseStr := "2"
	exponentStr := "-4"
	expectedNumStr := "0.0625"
	maxPrecision := uint(17)

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
	}

	bINumExponent, err := BigIntNum{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
	}

	result, err := BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}.Pwr(bINumBase, bINumExponent, maxPrecision). " +
			"bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
			bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedNumStr != result.GetNumStr() {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, result.GetNumStr())
	}

}
