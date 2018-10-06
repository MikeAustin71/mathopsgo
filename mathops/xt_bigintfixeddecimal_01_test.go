package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntFixedDecimal_CopyIn_01(t *testing.T) {
	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_02(t *testing.T) {
	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_03(t *testing.T) {
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_04(t *testing.T) {
	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyIn_05(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyIn(fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_01(t *testing.T) {
	expectedNumStr := "-123.45"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_02(t *testing.T) {
	expectedNumStr := "123.45"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_03(t *testing.T) {
	expectedNumStr := "12345"

	originalNum := big.NewInt(12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_04(t *testing.T) {
	expectedNumStr := "-12345"

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyInPtr_05(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := BigIntFixedDecimal{}

	fD2.CopyInPtr(&fixedDec)

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_01(t *testing.T) {
	expectedNumStr := "894.1234"

	originalNum := big.NewInt(8941234)
	originalNumPrecision := uint(4)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_02(t *testing.T) {
	expectedNumStr := "-894.1234"

	originalNum := big.NewInt(-8941234)
	originalNumPrecision := uint(4)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_CopyOut_03(t *testing.T) {
	expectedNumStr := "0.000"

	originalNum := big.NewInt(0)
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)

	fD2 := fixedDec.CopyOut()

	if expectedNumStr != fD2.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fD2.GetNumStr())
	}

}

func TestBigIntFixedDecimal_FormatNumStr_01(t *testing.T) {

	expectedNumStr := "-123.45"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := big.NewInt(-12345)
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.New(originalNum, originalNumPrecision)


	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_02(t *testing.T) {

	expectedNumStr := "123.45"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)


	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_03(t *testing.T) {

	expectedNumStr := "(123.45)"
	mode := PARENTHESESNEGVALFMTMODE
	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_04(t *testing.T) {

	expectedNumStr := "-1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_05(t *testing.T) {

	expectedNumStr := "1234.56"
	mode := LEADMINUSNEGVALFMTMODE
	originalNum := 123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_06(t *testing.T) {

	expectedNumStr := "(1234.56)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -123456
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_07(t *testing.T) {

	expectedNumStr := "0"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_08(t *testing.T) {

	expectedNumStr := "0.000"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_09(t *testing.T) {

	expectedNumStr := "0.000"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 0000
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_12(t *testing.T) {

	expectedNumStr := "12345"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_13(t *testing.T) {


	expectedNumStr := "(12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_14(t *testing.T) {

	expectedNumStr := "-12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_15(t *testing.T) {

	expectedNumStr := "0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_FormatNumStr_16(t *testing.T) {

	expectedNumStr := "0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_17(t *testing.T) {

	expectedNumStr := "-0.00012345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_18(t *testing.T) {

	expectedNumStr := "-0.12345"
	mode := LEADMINUSNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_19(t *testing.T) {

	expectedNumStr := "(0.00012345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_20(t *testing.T) {

	expectedNumStr := "(0.12345)"
	mode := PARENTHESESNEGVALFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_21(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := -12345
	originalNumPrecision := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_22(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_23(t *testing.T) {

	expectedNumStr := "0012345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(7)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}
}

func TestBigIntFixedDecimal_FormatNumStr_24(t *testing.T) {

	expectedNumStr := "12345"
	mode := ABSOLUTEPURENUMSTRFMTMODE

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.FormatNumStr(mode)

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_01(t *testing.T) {

	expectedNumStr := "123.45"

	originalNum := 12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_02(t *testing.T) {

	expectedNumStr := "-123.45"

	originalNum := -12345
	originalNumPrecision := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_GetNumStr_03(t *testing.T) {

	expectedNumStr := "0.000"

	originalNum := 0
	originalNumPrecision := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	outStr := fixedDec.GetNumStr()

	if expectedNumStr != outStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, outStr)
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_01(t *testing.T) {

	expectedNumStr := "-0.12345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_02(t *testing.T) {

	expectedNumStr := "12.345"

	originalNum := 12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_03(t *testing.T) {

	expectedNumStr := "-12.345"

	originalNum := -12345
	originalNumPrecision := uint(2)
	exponent := uint(1)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_04(t *testing.T) {

	expectedNumStr := "0.00012345"

	originalNum := 12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_05(t *testing.T) {

	expectedNumStr := "-0.00012345"

	originalNum := -12345
	originalNumPrecision := uint(0)
	exponent := uint(8)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_06(t *testing.T) {

	expectedNumStr := "0.0057"

	originalNum := 57
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}

func TestBigIntFixedDecimal_DivideByTenToPwr_07(t *testing.T) {

	expectedNumStr := "0"

	originalNum := 0
	originalNumPrecision := uint(0)
	exponent := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(originalNum, originalNumPrecision)

	fixedDec.DivideByTenToPwr(exponent)

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

}
