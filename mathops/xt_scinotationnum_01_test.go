package mathops

import "testing"

func TestSciNotationNum_SetNumStr_01(t *testing.T) {

	numStr := "2.652E4"

	expectedStr := "2.652e+4"

	mantissaLen := uint(3)

	sciNot1 := SciNotationNum{}.New()

	err := sciNot1.SetNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by sciNot1.SetNumStr(numStr). Error='%v'", err.Error())
	}

	resultStr, err := sciNot1.GetSciNotationStr(mantissaLen)

	if err != nil {
		t.Errorf("Error returned by sciNot1.GetSciNotationStr(mantissaLen). " +
			"Error='%v'", err.Error())
	}

	if expectedStr != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, resultStr)
	}
}
