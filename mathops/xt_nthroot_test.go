package mathops

import (
	"math/big"
	"testing"
)

/*
	These tests are designed to test library methods found in
	source file nthroot.go .

	This test file is located in source code repository:

       https://github.com/MikeAustin71/mathhlpr.git

*/

func TestNthRootOp_GetNthRootFloat32_01(t *testing.T) {
	num := float32(125.0)
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootFloat32(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootFloat32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootFloat64_01(t *testing.T) {
	num := float64(125.0)
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootFloat64(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootFloat64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootBigFloat_01(t *testing.T) {
	num := big.NewFloat(float64(125.0))
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootBigFloat(num, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootBigFloat() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootInt32_01(t *testing.T) {

	num := 125
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootInt(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootInt64_01(t *testing.T) {

	num := int64(125)
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootInt64(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootInt64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootBigInt_01(t *testing.T) {

	num := big.NewInt(125)
	nthRoot := 5
	maxPrecision := 14
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootBigInt(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootBigInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "125"
	nthRoot, err := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 14
	expected := "2.62652780440377"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_02(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423"

	nthRoot, err := IntAry{}.NewInt(6,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(6,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 13
	expected := "13.3276982415963"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"

	nthRoot, err := IntAry{}.NewInt(6,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(6,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 13
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_04(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"

	nthRoot, err := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(3,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 2
	expected := "-3.00"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_05(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"

	nthRoot, err := IntAry{}.NewInt(4,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(4,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 2
	originalNum.SetIntAryWithNumStr(numStr1)
	_, err = nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for negative number with even nthRoot. No Error triggered")
	}

}

func TestNthRootOp_GetNthRootIntAry_06(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "-5604423.924"

	nthRoot, err := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 13

	expected := "-22.3720713464898"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_07(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "5604423.924"

	nthRoot := IntAry{}.NewZero(0)

	maxPrecision := 1
	expected := "1.0"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_08(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "27"

	nthRoot := IntAry{}.NewOne(0)

	maxPrecision := 2
	originalNum.SetIntAryWithNumStr(numStr1)

	_, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for nthRoot == 1. No Error triggered")
	}

}

func TestNthRootOp_GetNthRootIntAry_09(t *testing.T) {
	nRt := NthRootOp{}
	numStr1 := "0.027"

	nthRoot, err := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(3,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 6

	expected := "0.300000"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_10(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "0.0005"

	nthRoot, err := IntAry{}.NewInt(9,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(9,0) Error='%v'",
			err.Error())
	}

	expected := "0.429752972587713"
	maxPrecision := 15

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}



func TestNthRootOp_GetNthRootIntAry_11(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "200000.000005"

	nthRoot := IntAry{}.NewTwo(0)

	expected := "447.213595505548"
	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_12(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "200001.100005"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "447.214825341245"
	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_13(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "2000000.0000005"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "1414.21356237327"
	maxPrecision := 11

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
			expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_14(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "20000000.00000005"

	nthRoot, err := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	expected := "271.441761659491"

	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_15(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "20000200.00020005"

	nthRoot, err := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(3,0) Error='%v'",
			err.Error())
	}

	expected := "271.442666463252"

	maxPrecision := 12

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_16(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "2020020.1010205"
	nthRoot := IntAry{}.NewTwo(0)
	expected := "1421.27411185193"
	maxPrecision := 11

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_17(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "209050307.020509033"

	nthRoot := IntAry{}.NewTwo(0)

	expected := "14458.5720947993"

	maxPrecision := 10

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
			expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_18(t *testing.T) {

	nRt := NthRootOp{}
	radicandStr := "500001"

	nthRoot, err := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	expected := "13.7973021335264"

	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .",
					expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_19(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "500001.00000009"

	nthRoot, err := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	expected := "13.7973021335269"
	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_20(t *testing.T) {
	nRt := NthRootOp{}
	radicandStr := "500001.00000009"

	nthRoot, err := IntAry{}.NewInt(3,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(3,0) Error='%v'",
			err.Error())
	}

	expected := "79.3701055117479"
	maxPrecision := 13

	originalNum, _ := IntAry{}.NewNumStr(radicandStr)

	ai, err := nRt.GetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_21(t *testing.T) {

	radicandStr := "-8000"
	maxPrecision := 20
	nthRoot, err := IntAry{}.NewInt(4,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(4,0) Error='%v'",
			err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicandStr)

	nthRt := NthRootOp{}
	_, err = nthRt.GetNthRootIntAry(&origRadicand, &nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected an Error. Negative Radicand with even NthRootInt. " +
			"Instead, NO ERROR WAS RETURNED.")
	}

}

func TestNthRootOp_GetNthRootIntAry_22(t *testing.T) {
	radicandStr := "8"
	nthRootStr := "0.4"
	expected := "181.01933598375616624661615669884"
	maxPrecision := 29

	nthRoot, err := IntAry{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRootStr, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicandStr)

	nthRt := NthRootOp{}

	ai, err := nthRt.GetNthRootIntAry(&origRadicand, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.GetNthRootIntAry(...) - %v",
				err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_23(t *testing.T) {
	radicand := "8"
	nthRoot := "-3"
	expected := "0.5"
	maxPrecision := 1

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)



	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
				err.Error())
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_24(t *testing.T) {
	radicand := "8"
	nthRoot := "-3.2"
	expectedStr := "0.52213689121370692016098323936996"
	maxPrecision := 32

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)



	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
				err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_25(t *testing.T) {

	radicand := "8.2"
	nthRoot := "-3.2"
	//              0.12345678901234567890123456789012
	expectedStr := "0.5181233574858042598812721854708"
	maxPrecision := 31

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
				err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_26(t *testing.T) {

	radicand := "-8.2"
	nthRoot := "-3.2"
	maxPrecision := 31

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	_, err = NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected a valid error object to be returned. Error: err==nil!" )
	}

}

func TestNthRootOp_GetNthRootIntAry_27(t *testing.T) {

	radicand := "5.967"
	nthRoot := "-2.894"
	//              0.12345678901234567890123456789012
	expectedStr := "0.53944021275349493325378163087104"
	maxPrecision := 32

	iaNthRoot, err := IntAry{}.NewNumStr(nthRoot)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error='%v'",
			nthRoot, err.Error())
	}

	origRadicand, _ := IntAry{}.NewNumStr(radicand)

	ai, err := NthRootOp{}.NewNthRoot(&origRadicand, &iaNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from NthRootOp{}.NewNthRoot(...) - %v",
			err.Error())
	}

	if expectedStr != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expectedStr, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootFloat32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float32(2686.5)


	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootFloat64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float64(2686.5)

	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigFloat_01(t *testing.T) {

	nRt := NthRootOp{}
	num, ok := big.NewFloat(0).SetString("2686.5")

	if !ok {
		t.Error("Conversion failed big.NewFloat(0).SetString(\"2686.5\")")
	}

	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigFloat(num, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigFloat() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int32(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int64(26865)
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := big.NewInt(int64(26865))
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "2686.5"
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetSquareRootIntAry(&originalNum, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "125"

	nthRoot, err := IntAry{}.NewInt(5,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(5,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 14
	expected := "2.62652780440377"
	originalNum.SetIntAryWithNumStr(numStr1)

	err = nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_02(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423"

	nthRoot, err := IntAry{}.NewInt(6,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(6,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 13
	expected := "13.3276982415963"
	originalNum.SetIntAryWithNumStr(numStr1)
	err = nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"

	nthRoot, err := IntAry{}.NewInt(6,0)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewInt(6,0) Error='%v'",
			err.Error())
	}

	maxPrecision := 13
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	err = nRt.SetNthRootIntAry(&originalNum, &nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.SetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}
