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

	radicand := "0.027"
	nthRoot := "3"
	expectedStr := "0.300000"
	maxPrecision := uint(6)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_10(t *testing.T) {

	radicand := "0.0005"
	nthRoot := "9"
	expectedStr := "0.429752972587713"
	maxPrecision := uint(15)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_11(t *testing.T) {

	radicand := "200000.000005"
	nthRoot := "2"
	expectedStr := "447.213595505548"
	maxPrecision := uint(12)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_12(t *testing.T) {

	radicand := "200001.100005"
	nthRoot := "2"
	expectedStr := "447.214825341245"
	maxPrecision := uint(12)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_13(t *testing.T) {

	radicand := "2000000.0000005"
	nthRoot := "2"
	expectedStr := "1414.21356237327"
	maxPrecision := uint(11)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_14(t *testing.T) {

	radicand := "20000000.00000005"
	nthRoot := "3"
	expectedStr := "271.441761659491"
	maxPrecision := uint(12)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_15(t *testing.T) {

	radicand := "20000200.00020005"
	nthRoot := "3"
	expectedStr := "271.442666463252"
	maxPrecision := uint(12)

	bINumBase, err  := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) " +
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.GetNthRoot(). " +
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

