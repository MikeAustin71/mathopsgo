package mathops

import (
	"testing"
	)

func TestIntAryMathPower_PwrByMultiplication_01(t *testing.T) {

	// Time: 0-Nanoseconds
	baseStr := "45"
	exponentStr := "12"
	expectedResultStr := "68952523554931640625"
	minPrecision := 0
	maxPrecision := 0


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestIntAryMathPower_PwrByMultiplication_02(t *testing.T) {

	// Time: 0-Nanoseconds
	baseStr := "45"
	exponentStr := "12"
	expectedResultStr := "68952523554931640625"
	minPrecision := 0
	maxPrecision := 0


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestIntAryMathPower_PwrByMultiplication_03(t *testing.T) {

	// Time: 0-Nanoseconds
	baseStr := "45"
	exponentStr := "120"
	expectedResultStr := "2429414689006507011047680668198610544614376056243160093821961151708217872022541081600097552192834563575596196518385368260399467853246465311637303360756939677768395657864175518625415861606597900390625"
	minPrecision := 0
	maxPrecision := 0


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestIntAryMathPower_PwrByMultiplication_04(t *testing.T) {

	// 1-Milliseconds 1011-Microseconds 1011498-Nanoseconds
	baseStr := "45"
	exponentStr := "-2"
	//                      12345678901234567890123456789012345
	expectedResultStr := "0.00049382716049382716049382716049383"
	minPrecision := 0
	maxPrecision := 35


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestIntAryMathPower_PwrByMultiplication_05(t *testing.T) {

	// 5-Milliseconds 5997-Microseconds 5997690-Nanoseconds
	baseStr := "19"
	exponentStr := "2.3"
	//                        12345678901234567890123456789012345
	expectedResultStr := "873.23931881701910176203214553167"
	minPrecision := 0
	maxPrecision := 29


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestIntAryMathPower_PwrByMultiplication_06(t *testing.T) {

	// 8-Milliseconds 8994-Microseconds 8994984-Nanoseconds
	baseStr := "19"
	exponentStr := "-2.3"
	//                      12345648901234567890123456789012
	expectedResultStr := "0.00114516144480839927662319331457"
	minPrecision := 0
	maxPrecision := 32


	iaBase, err := IntAry{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(baseStr). " +
			"baseStr='%v' Error='%v' \n", baseStr, err.Error())
		return

	}

	iaExponent, err := IntAry{}.NewNumStr(exponentStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(exponentStr). " +
			"exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
		return

	}

	actualResult, err := IntAryMathPower{}.PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := actualResult.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}
