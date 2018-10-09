package mathops

import (
	"math/big"
	"testing"
)

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

func TestBigIntFixedDecimal_GetIntegerFractionalParts_01(t *testing.T) {

	numStr := "8596.1234567"
	expectedInt := "8596"
	expectedFrac:= "0.1234567"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	intPart, fracPart := fixedDec.GetIntegerFractionalParts()


	if expectedInt != intPart.GetNumStr() {
		t.Errorf("Error: Expected integer part='%v'. Instead, integer part='%v'. ",
			expectedInt, intPart.GetNumStr())
	}

	if expectedFrac != fracPart.GetNumStr() {
		t.Errorf("Error: Expected fractional part='%v'. Instead, fractional part='%v'. ",
			expectedFrac, fracPart.GetNumStr())
	}

}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_02(t *testing.T) {

	numStr := "-8596.1234567"
	expectedInt := "-8596"
	expectedFrac:= "-0.1234567"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	intPart, fracPart := fixedDec.GetIntegerFractionalParts()


	if expectedInt != intPart.GetNumStr() {
		t.Errorf("Error: Expected integer part='%v'. Instead, integer part='%v'. ",
			expectedInt, intPart.GetNumStr())
	}

	if expectedFrac != fracPart.GetNumStr() {
		t.Errorf("Error: Expected fractional part='%v'. Instead, fractional part='%v'. ",
			expectedFrac, fracPart.GetNumStr())
	}

}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_03(t *testing.T) {

	numStr := "0"
	expectedInt := "0"
	expectedFrac:= "0"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	intPart, fracPart := fixedDec.GetIntegerFractionalParts()


	if expectedInt != intPart.GetNumStr() {
		t.Errorf("Error: Expected integer part='%v'. Instead, integer part='%v'. ",
			expectedInt, intPart.GetNumStr())
	}

	if expectedFrac != fracPart.GetNumStr() {
		t.Errorf("Error: Expected fractional part='%v'. Instead, fractional part='%v'. ",
			expectedFrac, fracPart.GetNumStr())
	}

}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_04(t *testing.T) {

	numStr := "859"
	expectedInt := "859"
	expectedFrac:= "0"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	intPart, fracPart := fixedDec.GetIntegerFractionalParts()


	if expectedInt != intPart.GetNumStr() {
		t.Errorf("Error: Expected integer part='%v'. Instead, integer part='%v'. ",
			expectedInt, intPart.GetNumStr())
	}

	if expectedFrac != fracPart.GetNumStr() {
		t.Errorf("Error: Expected fractional part='%v'. Instead, fractional part='%v'. ",
			expectedFrac, fracPart.GetNumStr())
	}

}

func TestBigIntFixedDecimal_GetIntegerFractionalParts_05(t *testing.T) {

	numStr := "-859"
	expectedInt := "-859"
	expectedFrac:= "0"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntFixedDecimal{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'", numStr, err.Error())
	}

	intPart, fracPart := fixedDec.GetIntegerFractionalParts()


	if expectedInt != intPart.GetNumStr() {
		t.Errorf("Error: Expected integer part='%v'. Instead, integer part='%v'. ",
			expectedInt, intPart.GetNumStr())
	}

	if expectedFrac != fracPart.GetNumStr() {
		t.Errorf("Error: Expected fractional part='%v'. Instead, fractional part='%v'. ",
			expectedFrac, fracPart.GetNumStr())
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

