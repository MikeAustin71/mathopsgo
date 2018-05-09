package mathops

import (
	"testing"
	"math/big"
)

func TestBigIntNum_Floor_01(t *testing.T) {
	nStr := "5.95"
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_02(t *testing.T) {
	nStr := "5.05"
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_03(t *testing.T) {
	nStr := "5"
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_04(t *testing.T) {
	nStr := "-5.05"
	expectedNumStr := "-6"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_05(t *testing.T) {
	nStr := "2.4"
	expectedNumStr := "2"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_06(t *testing.T) {
	nStr := "2.9"
	expectedNumStr := "2"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_07(t *testing.T) {
	nStr := "-2.7"
	expectedNumStr := "-3"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_08(t *testing.T) {
	nStr := "-2"
	expectedNumStr := "-2"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_09(t *testing.T) {
	nStr := "0"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_10(t *testing.T) {
	nStr := "18972.0000000000001"
	expectedNumStr := "18972"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_11(t *testing.T) {
	nStr := "-18972.0000000000001"
	expectedNumStr := "-18973"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_12(t *testing.T) {
	nStr := "0.0000000000001"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_13(t *testing.T) {
	nStr := "-0.0000000000001"
	expectedNumStr := "-1"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_14(t *testing.T) {
	nStr := "-189765342891.0000000000001"
	expectedNumStr := "-189765342892"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_Floor_15(t *testing.T) {
	nStr := "189765342891.0000000000001"
	expectedNumStr := "189765342891"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	floor := bINum1.Floor()

	actualNumStr := floor.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Floor NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != floor.GetPrecisionUint() {
		t.Errorf("Error: Expected Floor precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, floor.GetPrecisionUint())
	}
}

func TestBigIntNum_FormatNumStr_01(t *testing.T) {

	originalNumStr := "-123.45"
	expectedNumStr := "-123.45"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_02(t *testing.T) {

	originalNumStr := "123.45"
	expectedNumStr := "123.45"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_03(t *testing.T) {

	originalNumStr := "-123.45"
	expectedNumStr := "(123.45)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_04(t *testing.T) {

	originalNumStr := "-1234.56"
	expectedNumStr := "-1234.56"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_05(t *testing.T) {

	originalNumStr := "1234.56"
	expectedNumStr := "1234.56"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_06(t *testing.T) {

	originalNumStr := "-1234.56"
	expectedNumStr := "(1234.56)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_07(t *testing.T) {

	originalNumStr := "0"
	expectedNumStr := "0"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_08(t *testing.T) {

	originalNumStr := "0.000"
	expectedNumStr := "0.000"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_09(t *testing.T) {

	originalNumStr := "0.000"
	expectedNumStr := "0.000"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_12(t *testing.T) {

	originalNumStr := "12345"
	expectedNumStr := "12345"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_13(t *testing.T) {

	originalNumStr := "-12345"
	expectedNumStr := "(12345)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_14(t *testing.T) {

	originalNumStr := "-12345"
	expectedNumStr := "-12345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_15(t *testing.T) {

	originalBInt := big.NewInt(12345)
	precision := uint(8)
	expectedNumStr := "0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_16(t *testing.T) {

	originalBInt := big.NewInt(12345)
	precision := uint(5)
	expectedNumStr := "0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_17(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(8)
	expectedNumStr := "-0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_18(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(5)
	expectedNumStr := "-0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntNum_FormatNumStr_19(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(8)
	expectedNumStr := "(0.00012345)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatNumStr_20(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(5)
	expectedNumStr := "(0.12345)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntNum_FormatThousandsStr_01(t *testing.T) {

	originalNumStr := "1234"
	expectedNumStr := "1,234"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_02(t *testing.T) {

	originalNumStr := "123"
	expectedNumStr := "123"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_03(t *testing.T) {

	originalNumStr := "-1234"
	expectedNumStr := "(1,234)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_04(t *testing.T) {

	originalNumStr := "-1234"
	expectedNumStr := "-1,234"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_05(t *testing.T) {

	originalNumStr := "1234.567"
	expectedNumStr := "1,234.567"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_06(t *testing.T) {

	originalNumStr := "-1234.567"
	expectedNumStr := "-1,234.567"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_07(t *testing.T) {

	originalNumStr := "-1234.567"
	expectedNumStr := "(1,234.567)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_08(t *testing.T) {

	originalNumStr := "0"
	expectedNumStr := "0"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_09(t *testing.T) {

	originalNumStr := "0.0000"
	expectedNumStr := "0.0000"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_10(t *testing.T) {

	originalNumStr := "0.0000"
	expectedNumStr := "0.0000"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_11(t *testing.T) {

	originalNumStr := "1234567890.12"
	expectedNumStr := "1,234,567,890.12"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_12(t *testing.T) {

	originalNumStr := "-1234567890.12"
	expectedNumStr := "-1,234,567,890.12"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_13(t *testing.T) {

	originalNumStr := "-1234567890.12"
	expectedNumStr := "(1,234,567,890.12)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_14(t *testing.T) {

	originalNumStr := "1234567890"
	expectedNumStr := "1,234,567,890"
	mode := PARENTHESESNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_15(t *testing.T) {

	originalNumStr := "-1234567890"
	expectedNumStr := "-1,234,567,890"
	mode := LEADMINUSNEGVALFMTMODE

	bINum, err := BigIntNum{}.NewNumStr(originalNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(originalNumStr). "+
			" num1Str= '%v' Error='%v' ",
			originalNumStr, err.Error())
	}

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_16(t *testing.T) {

	originalBInt := big.NewInt(12345)
	precision := uint(8)
	expectedNumStr := "0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_17(t *testing.T) {

	originalBInt := big.NewInt(12345)
	precision := uint(5)
	expectedNumStr := "0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_18(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(8)
	expectedNumStr := "-0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_19(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(5)
	expectedNumStr := "-0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntNum_FormatThousandsStr_20(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(8)
	expectedNumStr := "(0.00012345)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntNum_FormatThousandsStr_21(t *testing.T) {

	originalBInt := big.NewInt(-12345)
	precision := uint(5)
	expectedNumStr := "(0.12345)"
	mode := PARENTHESESNEGVALFMTMODE

	bINum := BigIntNum{}.NewBigInt(originalBInt, precision)

	outStr := bINum.FormatThousandsStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

