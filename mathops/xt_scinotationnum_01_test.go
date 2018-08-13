package mathops

import "testing"

func TestSciNotationNum_SetNumStr_01(t *testing.T) {

	numStr := "2.652E4"

	expectedStr := "2.652e+4"

	sciNot1 := SciNotationNum{}.New()

	err := sciNot1.SetNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by sciNot1.SetNumStr(numStr). Error='%v'", err.Error())
	}

	resultStr := sciNot1.GetSciNotationStr()

	if expectedStr != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, resultStr)
	}
}
