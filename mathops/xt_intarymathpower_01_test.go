package mathops

import (
	"testing"
	)

func TestIntAryMathPower_PwrByMultiplication_01(t *testing.T) {

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
