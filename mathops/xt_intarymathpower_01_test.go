package mathops

import "testing"

func TestIntAryMathPower_MinimumRequiredPrecision_01(t *testing.T) {

	base := IntAry{}.NewInt(312, 2)
	exponent := IntAry{}.NewInt(4, 0)
	expectedResult := 8

	result, err := IntAryMathPower{}.MinimumRequiredPrecision(&base, &exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}."+
			"MinimumRequiredPrecision(base, exponent)"+
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestIntAryMathPower_MinimumRequiredPrecision_02(t *testing.T) {

	base := IntAry{}.NewInt(312345, 5)
	exponent := IntAry{}.NewInt(18, 0)
	expectedResult := 90

	result, err := IntAryMathPower{}.MinimumRequiredPrecision(&base, &exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}."+
			"MinimumRequiredPrecision(base, exponent)"+
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestIntAryMathPower_MinimumRequiredPrecision_03(t *testing.T) {

	base := IntAry{}.NewInt(-312345, 5)
	exponent := IntAry{}.NewInt(18, 0)
	expectedResult := 90

	result, err := IntAryMathPower{}.MinimumRequiredPrecision(&base, &exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}."+
			"MinimumRequiredPrecision(base, exponent)"+
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestIntAryMathPower_MinimumRequiredPrecision_04(t *testing.T) {

	base := IntAry{}.NewInt(312345, 5)
	exponent := IntAry{}.NewInt(-18, 0)
	expectedResult := 90

	result, err := IntAryMathPower{}.MinimumRequiredPrecision(&base, &exponent)

	if err != nil {
		t.Errorf("Error returned by BigIntMathPower{}."+
			"MinimumRequiredPrecision(base, exponent)"+
			"Error='%v' ", err.Error())
	}

	if expectedResult != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedResult, result)
	}

}

func TestIntAryMathPower_MinimumRequiredPrecision_05(t *testing.T) {

	base := IntAry{}.NewInt(312345,5)
	exponent := IntAry{}.NewUint64(12345678901234567890, 0)


	result, err := IntAryMathPower{}.MinimumRequiredPrecision(&base, &exponent)

	if err == nil {
		t.Error("Error: Expected error be returned. NO ERROR RETURNED!")
	}

	if result != int(2147483646) {
		t.Errorf("Error: Expected result='2147483646'. Instead, result='%v'",
			result)
	}

}

func TestIntAryMathPower_Pwr_01(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds
	baseStr := "2"
	exponentStr := "4"
	expectedNumStr := "16"
	maxPrecision := 17
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_02(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds
	baseStr := "2"
	exponentStr := "-4"
	expectedNumStr := "0.0625"
	maxPrecision := 17
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_03(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds
	baseStr := "37.241"
	exponentStr := "8"
	expectedNumStr := "3699735472699.4101912057680101525"
	maxPrecision := 19
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_04(t *testing.T) {

	// Time
	// 2-Milliseconds 998-Microseconds 700-Nanoseconds
	baseStr := "37"
	exponentStr := "3.25"
	expectedNumStr := "124926.79641959048051506133768818"
	maxPrecision := 26
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_05(t *testing.T) {

	// Time
	// 3-Milliseconds 997-Microseconds 400-Nanosecond

	baseStr := "37"
	exponentStr := "-3.25"
	expectedNumStr := "0.0000080046877744411952288377104402677"
	maxPrecision := 37
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_06(t *testing.T) {

	// Time
	// 1-Milliseconds 998-Microseconds 300-Nanoseconds

	baseStr := "-37"
	exponentStr := "-3"
	expectedNumStr := "-0.000019742167295125658894833474818866"
	maxPrecision := 36
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_07(t *testing.T) {

	// Time
	// 3-Milliseconds 997-Microseconds 500-Nanoseconds

	baseStr := "32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := 18
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_08(t *testing.T) {

	// Time
	// 2-Milliseconds 999-Microseconds 300-Nanoseconds

	baseStr := "-32"
	exponentStr := "-3.6"
	expectedNumStr := "0.000003814697265625"
	maxPrecision := 18
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_09(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "5"
	exponentStr := "-3"
	expectedNumStr := "0.008"
	maxPrecision := 3
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_10(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-5"
	exponentStr := "4"
	expectedNumStr := "625"
	maxPrecision := 0
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_11(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "-5"
	exponentStr := "5"
	expectedNumStr := "-3125"
	maxPrecision := 0
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_12(t *testing.T) {

	// Time
	// 1-Milliseconds 998-Microseconds 900-Nanoseconds

	baseStr := "4"
	exponentStr := "0.25"
	expectedNumStr := "1.4142135623730950488016887242097"
	maxPrecision := 31
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_13(t *testing.T) {

	// Time
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45"
	exponentStr := "120"
	expectedNumStr := "2429414689006507011047680668198610544614376056243160093821961151708217872022541081600097552192834563575596196518385368260399467853246465311637303360756939677768395657864175518625415861606597900390625"
	maxPrecision := 200
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_14(t *testing.T) {

	// Time
	// 4-Milliseconds 997-Microseconds 300-Nanoseconds

	baseStr := "19"
	exponentStr := "2.3"
	//                     12345648901234567890123456789
	expectedNumStr := "873.23931881701910176203214553167"
	maxPrecision := 29
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_15(t *testing.T) {

	// Time
	// 8-Milliseconds 47-Microseconds 300-Nanoseconds

	baseStr := "19"
	exponentStr := "-2.3"
	//                   12345648901234567890123456789012
	expectedNumStr := "0.00114516144480839927662319331457"
	maxPrecision := 32
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_16(t *testing.T) {

	// Time
	// 2-Milliseconds 998-Microseconds 200-Nanoseconds

	baseStr := "45"
	exponentStr := "-2"
	//                      12345678901234567890123456789012345
	expectedNumStr := "0.00049382716049382716049382716049383"
	maxPrecision := 35
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_17(t *testing.T) {

	baseStr := "0"
	exponentStr := "5"
	//                      12345678901234567890123456789012345
	maxPrecision := 35
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err == nil {
		t.Error("Error: Expected a error return. NO ERROR WAS RETURNED!\n")
		return

	}
}

func TestIntAryMathPower_Pwr_18(t *testing.T) {

	// Time:
	//  0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45.1"
	exponentStr := "0"
	//                      12345678901234567890123456789012345
	expectedNumStr := "1"
	maxPrecision := 5
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_19(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanoseconds

	baseStr := "45.1"
	exponentStr := "1"
	//                      12345678901234567890123456789012345
	expectedNumStr := "45.1"
	maxPrecision := 5
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_20(t *testing.T) {

	// Time:
	// 0-Milliseconds 0-Microseconds 0-Nanosecond

	baseStr := "-45.1"
	exponentStr := "1"
	//                      12345678901234567890123456789012345
	expectedNumStr := "-45.1"
	maxPrecision := 5
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_21(t *testing.T) {

	// Time:
	// 6-Milliseconds 995-Microseconds 900-Nanoseconds

	baseStr := "-45.6"
	exponentStr := "-3.2"
	//                      12345678901234567890123456789012345
	//                    0.00000491261243811417457984700270545
	// 4.91261243811417457984700270545e-6

	expectedNumStr := "0.00000491261243811417457984700270545"
	maxPrecision := 35
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}

func TestIntAryMathPower_Pwr_22(t *testing.T) {

	baseStr := "-45.632"
	exponentStr := "-1.01579"
	maxPrecision := 15
	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err == nil {
		t.Error("Expected Error to be returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"NO ERROR WAS RETURNED! \n")
		return

	}

}

func TestIntAryMathPower_Pwr_23(t *testing.T) {

	baseStr := "2.125"
	exponentStr := "-5"
	//                   12345678901234567890123456789012
	expectedNumStr := "0.02307838042845159759046157465153"

	maxPrecision := 32

	minPrecision := 0

	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). "+
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). "+
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). "+
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}
