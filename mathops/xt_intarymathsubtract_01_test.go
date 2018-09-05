package mathops

import "testing"

func TestIntAryMathSubtract_SubtractTotal_01(t *testing.T) {
	nStr1 := "25.72"
	nStr2 := "8.4"
	expectedStr := "17.32"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	IntAryMathSubtract{}.SubtractTotal(&ia1, &ia2)

	if expectedStr != ia1.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, ia1.GetNumStr())
	}

}

func TestIntAryMathSubtract_Subtract_01(t *testing.T) {
	nStr1 := "25.72"
	nStr2 := "8.4"
	expectedStr := "17.32"

	ia1, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr1). "+
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr2). "+
			"nStr2='%v' Error='%v' ", nStr2, err.Error())
	}

	ia3 := IntAryMathSubtract{}.Subtract(&ia1, &ia2)

	if expectedStr != ia3.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, ia1.GetNumStr())
	}
}
