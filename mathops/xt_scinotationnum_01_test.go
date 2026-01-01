package mathops

import (
	"strconv"
	"testing"
)

func TestSciNotationNum_SetNumStr_01(t *testing.T) {

	ePrefix := "TestSciNotationNum_SetNumStr_01"

	numStr := "2.652E4"

	expectedStr := "2.652e+4"

	mantissaLen := uint(3)

	mantissaLenStr := strconv.FormatUint(uint64(mantissaLen), 10)

	sciNot1 := new(SciNotationNum).New()

	err := sciNot1.SetNumStr(numStr)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err := sciNot1.SetNumStr(numStr)\n"+
			"numStr= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, numStr, err.Error())
		return
	}

	resultStr, err := sciNot1.GetSciNotationStr(mantissaLen)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"resultStr, err := sciNot1.GetSciNotationStr(mantissaLen)\n"+
			"mantissaLen= '%v'\n"+
			"Error= '%v'\n\n",
			ePrefix, mantissaLenStr, err.Error())
		return
	}

	if expectedStr != resultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, resultStr)
	}

	if expectedStr != resultStr {
		t.Errorf("%v\n"+
			"Error: resultStr is INVALID!\n"+
			"Because  expectedStr != resultStr\n"+
			"Expected resultStr = '%v'\n"+
			"  Actual resultStr = '%v'\n\n",
			ePrefix, expectedStr, resultStr)

		return
	}

	return
}
