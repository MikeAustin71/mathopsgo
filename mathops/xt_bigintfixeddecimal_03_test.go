package mathops

import (
	"testing"
)


func TestBigIntFixedDecimal_RoundToDecPlace_01(t *testing.T) {

	expectedNumStr := "-123.57"
	num := -123567
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_RoundToDecPlace_02(t *testing.T) {

	expectedNumStr := "123.57"
	num := 123567
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_03(t *testing.T) {

	expectedNumStr := "123.567"
	num := 123567
	precision := uint(3)
	roundToDec := uint(3)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_04(t *testing.T) {

	expectedNumStr := "123.5670"
	num := 123567
	precision := uint(3)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_05(t *testing.T) {

	expectedNumStr := "-123.5670"
	num := -123567
	precision := uint(3)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_06(t *testing.T) {

	expectedNumStr := "0.00"
	num := 0
	precision := uint(3)
	roundToDec := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_07(t *testing.T) {

	expectedNumStr := "654.123"
	num := 654123456
	precision := uint(6)
	roundToDec := uint(3)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_08(t *testing.T) {

	expectedNumStr := "654.1235"
	num := 654123456
	precision := uint(6)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_09(t *testing.T) {

	num := -654123456
	expectedNumStr := "-654.123"
	precision := uint(6)
	roundToDec := uint(3)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_10(t *testing.T) {

	num := -654123456
	expectedNumStr := "-654.1235"
	precision := uint(6)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_11(t *testing.T) {

	num := 654
	expectedNumStr := "654.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_12(t *testing.T) {

	num := 654123
	expectedNumStr := "654.123000000"
	precision := uint(3)
	roundToDec := uint(9)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_13(t *testing.T) {

	num := -654123
	expectedNumStr := "-654.123000000"
	precision := uint(3)
	roundToDec := uint(9)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_14(t *testing.T) {

	num := -654
	expectedNumStr := "-654.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_RoundToDecPlace_15(t *testing.T) {

	num := 0
	expectedNumStr := "0.0000"
	precision := uint(0)
	roundToDec := uint(4)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.RoundToDecPlace(roundToDec)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}


func TestBigIntFixedDecimal_SetNumStr_01(t *testing.T) {

	origNumStr:= "57.64"
	numStr := "89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}


	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())

	}

}

func TestBigIntFixedDecimal_SetNumStr_02(t *testing.T) {

	origNumStr := "257.1"
	numStr := "-89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_03(t *testing.T) {

	origNumStr := "0.000005"
	numStr := "0.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}

}

func TestBigIntFixedDecimal_SetNumStr_04(t *testing.T) {

	origNumStr := "97.8"
	numStr := ".123456789012"
	expectedNumStr := "0.123456789012"
	expectedPrecision := uint(12)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_05(t *testing.T) {

	origNumStr := "87"
	numStr := "-.123456789012"
	expectedNumStr := "-0.123456789012"
	expectedPrecision := uint(12)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_06(t *testing.T) {
	origNumStr := "97.9821"
	numStr := "10"
	expectedNumStr := "10"
	expectedPrecision := uint(0)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}

}

func TestBigIntFixedDecimal_SetNumStr_07(t *testing.T) {

	origNumStr := "9845.61"
	numStr := "-52"
	expectedNumStr := "-52"
	expectedPrecision := uint(0)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_08(t *testing.T) {

	origNumStr := "97"
	numStr := "-00052.1234"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_09(t *testing.T) {

	origNumStr := "87.123456"
	numStr := "(00052.1234)"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_10(t *testing.T) {

	origNumStr:="22.414141414"
	numStr := "(00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_11(t *testing.T) {
  origNumStr := "98.123456"
	numStr := "+00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}

}

func TestBigIntFixedDecimal_SetNumStr_12(t *testing.T) {

	origNumStr := "97782345646"
	numStr := "-00052.1234567890123456"
	expectedNumStr := "-52.1234567890123456"
	expectedPrecision := uint(16)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}

}

func TestBigIntFixedDecimal_SetNumStr_13(t *testing.T) {
  origNumStr := "52"
	numStr := "52 . 123 4567 8901 23456"
	expectedNumStr := "52.1234567890123456"
	expectedPrecision := uint(16)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_14(t *testing.T) {

	origNumStr := "987.123456"
	numStr := "5 2"
	expectedNumStr := "52"
	expectedPrecision := uint(0)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_15(t *testing.T) {

	origNumStr := "-8794521.12345"
	numStr := "    (52)     "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_16(t *testing.T) {

	origNumStr := "98"
	numStr := "    (52)    1234567 "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)
	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())
	}
}

func TestBigIntFixedDecimal_SetNumStr_17(t *testing.T) {

	origNumStr := "98123456789"
	expectedPrecision := uint(1050)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	fixedDec.SetNumStr(EulersNum1050)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if EulersNum1050 != fixedDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			EulersNum1050, fixedDec.GetNumStr())
	}

	if expectedPrecision != fixedDec.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, fixedDec.GetPrecision())

	}

}


func TestBigIntFixedDecimal_SetNumStr_18(t *testing.T) {

	origNumStr := "98"
	numStr := "abcdefghijklmnop"

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = fixedDec.SetNumStr(numStr)

	if err == nil {
		t.Error("Expected an Error from INVALID Number String. NO ERROR RETURNED!" )
	}
}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_01(t *testing.T) {

	num := 456123000
	precision := uint(6)
	expectedNumStr := "456.123"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_02(t *testing.T) {

	num := -456123000
	precision := uint(6)
	expectedNumStr := "-456.123"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_03(t *testing.T) {

	num := 0
	precision := uint(3)
	expectedNumStr := "0"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_04(t *testing.T) {

	num := 70
	precision := uint(1)
	expectedNumStr := "7"

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TrimTrailingFracZeros_05(t *testing.T) {

	num := uint64(7000000000000000000)
	precision := uint(17)
	expectedNumStr := "70"

	fixedDec := BigIntFixedDecimal{}.NewUInt64(num, precision)

	fixedDec.TrimTrailingFracZeros()

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_TruncToDecPlace_01(t *testing.T) {

	num := -123567
	precision := uint(3)
	expectedNumStr := "-123.56"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_02(t *testing.T) {
	num := 123567
	precision := uint(3)
	expectedNumStr := "123.56"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_03(t *testing.T) {

	num := 123567
	precision := uint(3)
	expectedNumStr := "123.567"
	truncToDec := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_04(t *testing.T) {

	num := 123567
	precision := uint(3)
	expectedNumStr := "123.5670"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_05(t *testing.T) {

	num := -123567
	precision := uint(3)
	expectedNumStr := "-123.5670"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_06(t *testing.T) {
	num := 0
	precision := uint(3)
	expectedNumStr := "0.00"
	truncToDec := uint(2)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_07(t *testing.T) {

	num := 654123456
	precision := uint(6)
	expectedNumStr := "654.123"
	truncToDec := uint(3)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_08(t *testing.T) {

	num := 654123456789
	precision := uint(9)
	expectedNumStr := "654.1234"
	truncToDec := uint(4)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_09(t *testing.T) {

	num := 654123456789
	precision := uint(9)
	expectedNumStr := "654"
	truncToDec := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_10(t *testing.T) {

	num := 654
	precision := uint(0)
	expectedNumStr := "654.00000"
	truncToDec := uint(5)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_11(t *testing.T) {

	num := 654123
	precision := uint(3)
	expectedNumStr := "654.123000000"
	truncToDec := uint(9)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_12(t *testing.T) {

	num := 0
	precision := uint(0)
	expectedNumStr := "0.000000"
	truncToDec := uint(6)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_13(t *testing.T) {

	num := 0
	precision := uint(6)
	expectedNumStr := "0"
	truncToDec := uint(0)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}

func TestBigIntFixedDecimal_TruncToDecPlace_14(t *testing.T) {
	num := 654123456789015
	precision := uint(12)
	expectedNumStr := "654.12345678901"
	truncToDec := uint(11)

	fixedDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixedDec.TruncToDecPlace(truncToDec)

	actualNumStr := fixedDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}
}
