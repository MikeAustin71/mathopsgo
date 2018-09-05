package mathops

import (
	"testing"
)

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
