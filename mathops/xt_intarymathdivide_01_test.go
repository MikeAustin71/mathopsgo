package mathops

import "testing"

func TestIntAryMathDivide_Divide_01(t *testing.T) {

	dividendStr := "25.6"
	divisorStr := "2.3"
	maxPrecision := 30
	minPrecision := 5
	//                          1         2         3
	//               0.123456789012345678901234567890
	expectedStr := "11.130434782608695652173913043478"

	dividend, err := IntAry{}.NewNumStr(dividendStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). " +
			"dividendStr='%v' Error='%v' ", dividendStr, err.Error())
	}

	divisor, err := IntAry{}.NewNumStr(divisorStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). " +
			"divisorStr='%v' Error='%v' ", divisorStr, err.Error())
	}

	quotient, err := IntAryMathDivide{}.Divide(&dividend, &divisor, minPrecision, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by IntAryMathDivide{}.Divide() " +
			"dividend='%v' divisor='%v' Error='%v'", dividend, divisor, err.Error())
	}

	if expectedStr != quotient.GetNumStr() {
		t.Errorf("Error: Expected quotient='%v'. Instead, quotient='%v' ",
			expectedStr, quotient.GetNumStr() )
	}

}
