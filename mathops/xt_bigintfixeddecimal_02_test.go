package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntFixedDecimal_FormatNumStr_01(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_01"
	expectedNumStr := "-123.45"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).New(originalNum, originalNumPrecision)\n\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, originalNum.Text(10), originalNumPrecision, err.Error())
		return
	}

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_02(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_02"
	expectedNumStr := "123.45"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_03(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_03"
	expectedNumStr := "(123.45)"
	mode := PARENTHESESNEGVALFMTMODE
	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_04(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_04"
	expectedNumStr := "-1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_05(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_05"
	expectedNumStr := "1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 123456
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_06(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_06"
	expectedNumStr := "(1234.56)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_07(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_FormatNumStr_07"
	expectedNumStr := "0"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0
	originalNumPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_08(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_08"
	expectedNumStr := "0.000"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_09(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_09"
	expectedNumStr := "0.000"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_12(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_12"
	expectedNumStr := "12345"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_13(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_13"
	expectedNumStr := "(12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_14(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_14"
	expectedNumStr := "-12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_15(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_15"
	expectedNumStr := "0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(8)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_16(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_16"
	expectedNumStr := "0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(5)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_17(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_17"
	expectedNumStr := "-0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_18(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_18"
	expectedNumStr := "-0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_19(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_19"
	expectedNumStr := "(0.00012345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_20(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_20"
	expectedNumStr := "(0.12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_21(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_21"
	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_22(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_22"
	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_23(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_23"
	expectedNumStr := "0012345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(7)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_FormatNumStr_24(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_FormatNumStr_24"
	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected  fixedDec.FormatNumStr = '%v'\n"+
			"Instead,  fixedDec.FormatNumStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_01(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_01"
	numStr := "8596.1234567"
	expectedInt := "8596"
	expectedFrac := "0.1234567"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_02(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_02"
	numStr := "-8596.1234567"
	expectedInt := "-8596"
	expectedFrac := "-0.1234567"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_03(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_03"
	numStr := "0"
	expectedInt := "0"
	expectedFrac := "0"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_04(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_04"
	numStr := "859"
	expectedInt := "859"
	expectedFrac := "0"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_05(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_05"
	numStr := "-859"
	expectedInt := "-859"
	expectedFrac := "0"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_06(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_06"
	numStr := "-859.00123456"
	expectedInt := "-859"
	expectedFrac := "-0.00123456"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_07(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetIntegerFractionalParts_07"
	numStr := "859.00123456"
	expectedInt := "859"
	expectedFrac := "0.00123456"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPart, fracPart, err := fixedDec.GetIntegerFractionalParts()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	intPartNumStr, err := intPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"intPartNumStr, err := intPart.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	fracPartNumStr, err := fracPart.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fracPartNumStr, err := fracPart.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedInt != intPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected intPartNumStr = '%v'\n"+
			"Instead, intPartNumStr = '%v'\n\n",
			ePrefix, expectedInt, intPartNumStr)

		return
	}

	if expectedFrac != fracPartNumStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected fracPartNumStr = '%v'\n"+
			"Instead, fracPartNumStr = '%v'\n\n",
			ePrefix, expectedFrac, fracPartNumStr)
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_01(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetMagnitude_01"
	numStr := "963256"
	expectedMagnitude := big.NewInt(5)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_02(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetMagnitude_02"
	numStr := "2"
	expectedMagnitude := big.NewInt(0)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_03(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetMagnitude_03"
	numStr := "32"
	expectedMagnitude := big.NewInt(1)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_04(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetMagnitude_04"
	numStr := "8456123921"
	expectedMagnitude := big.NewInt(9)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_05(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetMagnitude_05"
	numStr := "2.2"
	expectedMagnitude := big.NewInt(0)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_06(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetMagnitude_06"
	numStr := "8456123912.123"
	expectedMagnitude := big.NewInt(9)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_07(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetMagnitude_07"
	numStr := "-643,212.123"
	expectedMagnitude := big.NewInt(5)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetMagnitude_08(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetMagnitude_08"
	numStr := "324.123456"
	expectedMagnitude := big.NewInt(2)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\n"+
			"Error='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	fixedDecNumStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	magnitude, err := fixedDec.GetMagnitude()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"magnitude, err := fixedDec.GetMagnitude()\n"+
			"fixedDecNumStr= '%v'\nError='%v'\n\n",
			ePrefix, fixedDecNumStr, err.Error())
		return
	}

	cmpResult := expectedMagnitude.Cmp(magnitude)

	if cmpResult != 0 {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected magnitude = '%v'\n"+
			"Instead, magnitude = '%v'\n\n",
			ePrefix, expectedMagnitude.Text(10), magnitude.Text(10))
	}

	return
}

func TestBigIntFixedDecimal_GetNumStr_01(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetNumStr_01"
	expectedNumStr := "123.45"

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected outStr = '%v'\n"+
			"Instead, outStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_GetNumStr_02(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_GetNumStr_02"
	expectedNumStr := "-123.45"

	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected outStr = '%v'\n"+
			"Instead, outStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_GetNumStr_03(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetNumStr_03"

	expectedNumStr := "0.000"

	originalNum := 0
	originalNumPrecision := uint(3)

	fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)

	outStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec := new(BigIntFixedDecimal).NewInt(originalNum, originalNumPrecision)\n"+
			"originalNum= '%v'\n"+
			"originalNumPrecision= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, originalNum, originalNumPrecision, err.Error())
		return
	}

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected outStr = '%v'\n"+
			"Instead, outStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_GetNumStr_04(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_GetNumStr_04"
	expectedNumStr := "9876543210.12345678901234"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(expectedNumStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(expectedNumStr, '.')\n"+
			"expectedNumStr='%v'\nError='%v'\n\n", ePrefix, expectedNumStr, err.Error())
		return
	}

	outStr, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"outStr, err := fixedDec.GetNumStr()\n"+
			"expectedNumStr= '%v'\n"+
			"Error='%v'\n\n",
			ePrefix, expectedNumStr, err.Error())
		return
	}

	if expectedNumStr != outStr {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected outStr = '%v'\n"+
			"Instead, outStr = '%v'\n\n",
			ePrefix, expectedNumStr, outStr)
	}

	return
}

func TestBigIntFixedDecimal_Inverse_01(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_Inverse_01"
	numStr := "4"
	expectedValue := "0.25"
	maxPrecision := uint(5)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return
}

func TestBigIntFixedDecimal_Inverse_02(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_Inverse_02"
	numStr := "85.12345"
	expectedValue := "0.01174764415680990373392995702124"
	maxPrecision := uint(32)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return
}

func TestBigIntFixedDecimal_Inverse_03(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_Inverse_03"
	numStr := "-4"
	expectedValue := "-0.25"
	maxPrecision := uint(32)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return
}

func TestBigIntFixedDecimal_Inverse_04(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_Inverse_04"
	numStr := "97842.123456"
	expectedValue := "0.000010220546781670208316702050788328"
	maxPrecision := uint(36)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return
}

func TestBigIntFixedDecimal_Inverse_05(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_Inverse_05"
	numStr := "-5.12345678"
	expectedValue := "-0.19518072327722456165620274833274"
	maxPrecision := uint(32)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return

}

func TestBigIntFixedDecimal_Inverse_06(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_Inverse_06"
	numStr := "0"
	expectedValue := "0"
	maxPrecision := uint(32)

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	err = fixedDec.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"err = fixedDec.Inverse(maxPrecision)\n"+
			"maxPrecision= '%v'\n"+
			"Error='%v'\n\n", ePrefix, maxPrecision, err.Error())
		return
	}

	inverseValue, err := fixedDec.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"inverseValue, err := fixedDec.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if expectedValue != inverseValue {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected inverseValue = '%v'\n"+
			"Instead, inverseValue = '%v'\n\n",
			ePrefix, expectedValue, inverseValue)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_01(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_01"
	numStr := "4"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, true, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_02(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_IsEven_02"
	numStr := "-4"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, true, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_03(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_IsEven_03"

	numStr := "0"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, true, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_04(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_04"
	numStr := "5"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_05(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_05"
	numStr := "-5"

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_06(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_06"
	numStr := "2.2"
	//expectedValue := false

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_07(t *testing.T) {
	ePrefix := "BigIntFixedDecimal_IsEven_07"
	numStr := "-2.2"
	//expectedValue := false

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_08(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_08"
	numStr := "8388608"
	//expectedValue := true

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, true, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_09(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_09"
	numStr := "-8388608"
	//expectedValue := true

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if true != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, true, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_10(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_10"
	numStr := "8388609"
	//expectedValue := false

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}

func TestBigIntFixedDecimal_IsEven_11(t *testing.T) {

	ePrefix := "BigIntFixedDecimal_IsEven_11"
	numStr := "-8388609"
	//expectedValue := false

	fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
			"numStr='%v'\nError='%v'\n\n", ePrefix, numStr, err.Error())
		return
	}

	isEven := fixedDec.IsEven()

	if false != isEven {
		t.Errorf("%v\n"+
			"Error: Unexpected Result!\n"+
			"Expected isEven = '%v'\n"+
			"Instead, isEven = '%v'\n\n",
			ePrefix, false, isEven)
	}

	return
}
