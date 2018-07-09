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

func TestBigIntMathPower_Pwr_03(t *testing.T) {
	baseStr := "37.241"
	exponentStr := "8"
	expectedNumStr := "3699735472699.4101912057680101525"
	maxPrecision := uint(19)

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

func TestBigIntMathPower_Pwr_04(t *testing.T) {
	baseStr := "37"
	exponentStr := "3.25"
	expectedNumStr := "124926.79641959048051506133768818"
	maxPrecision := uint(26)

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

func TestBigIntMathPower_Pwr_05(t *testing.T) {
	baseStr := "37"
	exponentStr := "-3.25"
	expectedNumStr := "0.0000080046877744411952288377104402677"
	maxPrecision := uint(37)

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

func TestBigIntMathPower_Pwr_06(t *testing.T) {
	baseStr := "-37"
	exponentStr := "-3"
	expectedNumStr := "-0.000019742167295125658894833474818866"
	maxPrecision := uint(36)

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

func TestBigIntMathPower_Pwr_07(t *testing.T) {
	baseStr := "32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := uint(18)

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

func TestBigIntMathPower_Pwr_08(t *testing.T) {
	baseStr := "-32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := uint(18)

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

func TestBigIntMathPower_Pwr_09(t *testing.T) {
	baseStr := "5"
	exponentStr := "-3"
	expectedNumStr := "0.008"
	maxPrecision := uint(3)

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

func TestBigIntMathPower_Pwr_10(t *testing.T) {
	baseStr := "-5"
	exponentStr := "4"
	expectedNumStr := "625"
	maxPrecision := uint(0)

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

func TestBigIntMathPower_Pwr_11(t *testing.T) {
	baseStr := "-5"
	exponentStr := "5"
	expectedNumStr := "-3125"
	maxPrecision := uint(0)

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

func TestBigIntMathPower_Pwr_12(t *testing.T) {
	baseStr := "4"
	exponentStr := "0.25"
	expectedNumStr := "1.4142135623730950488016887242097"
	maxPrecision := uint(31)

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

func TestBigIntMathPower_Pwr_13(t *testing.T) {
	baseStr := "45"
	exponentStr := "120"
	expectedNumStr := "2429414689006507011047680668198610544614376056243160093821961151708217872022541081600097552192834563575596196518385368260399467853246465311637303360756939677768395657864175518625415861606597900390625"
	maxPrecision := uint(200)

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

func TestBigIntMathPower_Pwr_14(t *testing.T) {

	// Time:
	// 1-Milliseconds 1997-Microseconds 1997998-Nanoseconds
	baseStr := "19"
	exponentStr := "2.3"
	//                     12345648901234567890123456789
	expectedNumStr := "873.23931881701910176203214553167"
	maxPrecision := uint(29)

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

func TestBigIntMathPower_Pwr_15(t *testing.T) {
	// Time:
	// 1-Milliseconds 1999-Microseconds 1999498-Nanoseconds
	baseStr := "19"
	exponentStr := "-2.3"
	//                   12345648901234567890123456789012
	expectedNumStr := "0.00114516144480839927662319331457"
	maxPrecision := uint(32)

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
