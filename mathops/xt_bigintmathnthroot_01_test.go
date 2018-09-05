package mathops

import "testing"

func TestBigIntMathNthRoot_GetNthRootBigNum_01(t *testing.T) {

	baseStr := "125"
	nthRootStr := "5"
	maxPrecision := uint(14)
	expectedResult := "2.62652780440377"

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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
	expectedResult := "13.3276982415963"

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.NthRoot() for negative number with even nthRoot. No Error triggered")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_06(t *testing.T) {

	baseStr := "-5604423.924"
	nthRootStr := "5"
	maxPrecision := uint(13)
	expectedResult := "-22.3720713464898"

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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
	expectedResult := "1"

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(baseStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", baseStr, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRootStr, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.NthRoot() for nthRoot == 1. No Error triggered")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_09(t *testing.T) {

	radicand := "0.027"
	nthRoot := "3"
	expectedStr := "0.300000"
	maxPrecision := uint(6)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
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

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_16(t *testing.T) {

	radicand := "2020020.1010205"
	nthRoot := "2"
	expectedStr := "1421.27411185193"
	maxPrecision := uint(11)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_17(t *testing.T) {

	radicand := "209050307.020509033"
	nthRoot := "2"
	expectedStr := "14458.5720947993"
	maxPrecision := uint(10)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_18(t *testing.T) {

	radicand := "500001"
	nthRoot := "5"
	expectedStr := "13.7973021335264"
	maxPrecision := uint(13)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_19(t *testing.T) {

	radicand := "500001.00000009"
	nthRoot := "5"
	expectedStr := "13.7973021335269"
	maxPrecision := uint(13)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_20(t *testing.T) {

	radicand := "500001.00000009"
	nthRoot := "3"
	expectedStr := "79.3701055117479"
	maxPrecision := uint(13)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	result, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from bIMathNthRoot.NthRoot(). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != result.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead result= %v .",
			expectedStr, result.GetNumStr())
	}
}

func TestBigIntMathNthRoot_GetNthRootBigNum_21(t *testing.T) {

	radicand := "-8000"
	nthRoot := "4"
	maxPrecision := uint(20)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected an Error. Negative Radicand with even NthRootInt. " +
			"Instead, NO ERROR WAS RETURNED.")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_22(t *testing.T) {

	radicand := "8"
	nthRoot := "0.4"
	expectedStr := "181.01933598375616624661615669884"
	maxPrecision := uint(29)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(baseStr) "+
			"baseStr='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRootStr) "+
			"nthRootStr='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	nthRootResult, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by bIMathNthRoot.NthRoot(bINumBase, "+
			"bINumNthRoot, maxPrecision) bINumBase='%v' bINumNthRoot='%v' "+
			" maxPrecision='%v' Error='%v'.", bINumBase.GetNumStr(),
			bINumNthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedStr != nthRootResult.GetNumStr() {
		t.Errorf("Error: Expected NthRootInt Result='%v'.  Instead Result='%v'",
			expectedStr, nthRootResult.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_23(t *testing.T) {

	radicand := "8"
	nthRoot := "-3"
	expectedStr := "0.5"
	maxPrecision := uint(1)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(radicand) "+
			"radicand='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRoot) "+
			"nthRoot='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	nthRootResult, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by bIMathNthRoot.NthRoot(bINumBase, "+
			"bINumNthRoot, maxPrecision) bINumBase='%v' bINumNthRoot='%v' "+
			" maxPrecision='%v' Error='%v'.", bINumBase.GetNumStr(),
			bINumNthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedStr != nthRootResult.GetNumStr() {
		t.Errorf("Error: Expected NthRootInt Result='%v'.  Instead Result='%v'",
			expectedStr, nthRootResult.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_24(t *testing.T) {

	radicand := "8"
	nthRoot := "-3.2"
	expectedStr := "0.52213689121370692016098323936996"
	maxPrecision := uint(32)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(radicand) "+
			"radicand='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRoot) "+
			"nthRoot='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	nthRootResult, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by bIMathNthRoot.NthRoot(bINumBase, "+
			"bINumNthRoot, maxPrecision) bINumBase='%v' bINumNthRoot='%v' "+
			" maxPrecision='%v' Error='%v'.", bINumBase.GetNumStr(),
			bINumNthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedStr != nthRootResult.GetNumStr() {
		t.Errorf("Error: Expected NthRootInt Result='%v'.  Instead Result='%v'",
			expectedStr, nthRootResult.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_25(t *testing.T) {

	radicand := "8.2"
	nthRoot := "-3.2"
	//              0.12345678901234567890123456789012
	expectedStr := "0.5181233574858042598812721854708"
	maxPrecision := uint(31)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(radicand) "+
			"radicand='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRoot) "+
			"nthRoot='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	nthRootResult, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by bIMathNthRoot.NthRoot(bINumBase, "+
			"bINumNthRoot, maxPrecision) bINumBase='%v' bINumNthRoot='%v' "+
			" maxPrecision='%v' Error='%v'.", bINumBase.GetNumStr(),
			bINumNthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedStr != nthRootResult.GetNumStr() {
		t.Errorf("Error: Expected NthRootInt Result='%v'.  Instead Result='%v'",
			expectedStr, nthRootResult.GetNumStr())
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_26(t *testing.T) {

	radicand := "-8.2"
	nthRoot := "-3.2"
	maxPrecision := uint(31)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(radicand) "+
			"radicand='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRoot) "+
			"nthRoot='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	_, err = bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err == nil {
		t.Error("Error: Expected err!=nil. Instead, err==nil. Should have received an error " +
			"on this calculation, but DID NOT RECEIVE AN ERROR!")
	}

}

func TestBigIntMathNthRoot_GetNthRootBigNum_27(t *testing.T) {

	radicand := "5.967"
	nthRoot := "-2.894"
	//              0.12345678901234567890123456789012
	expectedStr := "0.53944021275349493325378163087104"
	maxPrecision := uint(32)

	bINumBase, err := BigIntNum{}.NewNumStr(radicand)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(radicand) "+
			"radicand='%v' Error='%v'", radicand, err.Error())
	}

	bINumNthRoot, err := BigIntNum{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from BigIntNum{}.NewNumStr(nthRoot) "+
			"nthRoot='%v' Error='%v'", nthRoot, err.Error())
	}

	bIMathNthRoot := BigIntMathNthRoot{}

	nthRootResult, err := bIMathNthRoot.GetNthRoot(bINumBase, bINumNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by bIMathNthRoot.NthRoot(bINumBase, "+
			"bINumNthRoot, maxPrecision) bINumBase='%v' bINumNthRoot='%v' "+
			" maxPrecision='%v' Error='%v'.", bINumBase.GetNumStr(),
			bINumNthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	if expectedStr != nthRootResult.GetNumStr() {
		t.Errorf("Error: Expected NthRootInt Result='%v'.  Instead Result='%v'",
			expectedStr, nthRootResult.GetNumStr())
	}

}
