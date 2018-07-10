package mathops

import "testing"

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

	err = IntAryMathPower{}.Pwr(&iaBase, &iaExponent, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathPower{}.PwrByMultiplication(...). " +
			"Error='%v' \n", err.Error())
		return

	}

	actualResultStr := iaBase.GetNumStr()

	if expectedNumStr != actualResultStr {
		t.Errorf("Error: Expected Result='%v'. Instead, Result='%v'",
			expectedNumStr, actualResultStr)
	}

}
