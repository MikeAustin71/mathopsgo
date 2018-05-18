package mathops

import "testing"

func TestBigIntMathNthRoot_GetNthRootBigNum_01(t *testing.T) {

	baseStr := "125"
	nthRootStr := "5"
	maxPrecision := uint(14)
	expectedResult := "2.62652780440377"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_02(t *testing.T) {

	baseStr := "5604423"
	nthRootStr := "6"
	maxPrecision := uint(13)
	expectedResult  := "13.3276982415963"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_03(t *testing.T) {

	baseStr := "5604423.924"
	nthRootStr := "6"
	maxPrecision := uint(13)
	expectedResult := "13.3276986078187"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}


}

func TestBigIntMathNthRoot_GetNthRootBigNum_04(t *testing.T) {

	baseStr := "-27"
	nthRootStr := "3"
	maxPrecision := uint(2)
	expectedResult := "-3.00"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_05(t *testing.T) {

	baseStr := "-27"
	nthRootStr := "4"
	maxPrecision := uint(2)

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRoot() for negative number with even nthRoot. No Error triggered")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_06(t *testing.T) {

	baseStr := "-5604423.924"
	nthRootStr := "5"
	maxPrecision := uint(13)
	expectedResult := "-22.3720713464898"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_07(t *testing.T) {

	baseStr := "5604423.924"
	nthRootStr := "0"
	maxPrecision := uint(1)
	expectedResult := "1.0"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_08(t *testing.T) {

	baseStr := "27"
	nthRootStr := "1"
	maxPrecision := uint(2)

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRoot() for nthRoot == 1. No Error triggered")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_09(t *testing.T) {

	baseStr := "0.027"
	nthRootStr := "3"
	maxPrecision := uint(6)
	expectedResult := "0.300000"

	bINumBase, err  := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedResult != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedResult, result.GetNumStr())
	}

}

