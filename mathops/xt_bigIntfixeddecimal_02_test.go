package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntFixedDecimal_MultiplyByTenToPwr_01(t *testing.T) {

	expectedNumStr := "105.6752"
	num := 1056752
	precision := uint(4)
	exponent := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_02(t *testing.T) {

	expectedNumStr := "-105.6752"
	num := -1056752
	precision := uint(4)
	exponent := uint(0)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_03(t *testing.T) {

	expectedNumStr := "1056.752"
	num := 1056752
	precision := uint(4)
	exponent := uint(1)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_04(t *testing.T) {

	expectedNumStr := "-1056.752"
	num := -1056752
	precision := uint(4)
	exponent := uint(1)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_05(t *testing.T) {

	expectedNumStr := "10567.52"
	num := 1056752
	precision := uint(4)
	exponent := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_06(t *testing.T) {

	expectedNumStr := "-10567.52"
	num := -1056752
	precision := uint(4)
	exponent := uint(2)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_07(t *testing.T) {

	expectedNumStr := "10567520000"
	num := 1056752
	precision := uint(4)
	exponent := uint(8)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_08(t *testing.T) {

	expectedNumStr := "0"
	num := 0
	precision := uint(4)
	exponent := uint(5)

	fixDec := BigIntFixedDecimal{}.NewInt(num, precision)

	fixDec.MultiplyByTenToPower(exponent)

	actualNumStr := fixDec.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_01(t *testing.T) {

	// multiplicand = 23.321
	multiplicandBInt := big.NewInt(23321)
	multiplicandPrecision := uint(3)
	exponent := uint(5)
	expectedResult := "746.272"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_02(t *testing.T) {

	// multiplicand = 8
	multiplicandBInt := big.NewInt(8)
	multiplicandPrecision := uint(0)
	exponent := uint(10)
	expectedResult := "8192"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_03(t *testing.T) {

	// multiplicand = 9.871234
	multiplicandBInt := big.NewInt(9871234)
	multiplicandPrecision := uint(6)
	exponent := uint(1)
	expectedResult := "19.742468"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_04(t *testing.T) {

	// multiplicand = -9.871234
	multiplicandBInt := big.NewInt(-9871234)
	multiplicandPrecision := uint(6)
	exponent := uint(3)
	expectedResult := "-78.969872"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_05(t *testing.T) {

	// multiplicand = 8
	multiplicandBInt := big.NewInt(8)
	multiplicandPrecision := uint(0)
	exponent := uint(0)
	expectedResult := "8"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_06(t *testing.T) {

	// (0.12345 x 2^15 = 4045.2096)
	multiplicandBInt := big.NewInt(12345)
	multiplicandPrecision := uint(5)
	exponent := uint(15)
	expectedResult := "4045.2096"

	fixDec := BigIntFixedDecimal{}.New(multiplicandBInt, multiplicandPrecision)

	fixDec.MultiplyByTwoToPower(exponent)

	if expectedResult != fixDec.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expectedResult, fixDec.GetNumStr())
	}
}

func TestBigIntFixedDecimal_NewNumStr_01(t *testing.T) {

	numStr := "89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_02(t *testing.T) {

	numStr := "-89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_03(t *testing.T) {

	numStr := "0.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_04(t *testing.T) {

	numStr := ".123456789012"
	expectedNumStr := "0.123456789012"
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_05(t *testing.T) {

	numStr := "-.123456789012"
	expectedNumStr := "-0.123456789012"
	expectedPrecision := uint(12)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_06(t *testing.T) {

	numStr := "10"
	expectedNumStr := "10"
	expectedPrecision := uint(0)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_07(t *testing.T) {

	numStr := "-52"
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_08(t *testing.T) {

	numStr := "-00052.1234"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_09(t *testing.T) {

	numStr := "(00052.1234)"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_10(t *testing.T) {

	numStr := "(00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_11(t *testing.T) {

	numStr := "+00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_12(t *testing.T) {

	numStr := "-00052.1234567890123456"
	expectedNumStr := "-52.1234567890123456"
	expectedPrecision := uint(16)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_13(t *testing.T) {

	numStr := "52 . 123 4567 8901 23456"
	expectedNumStr := "52.1234567890123456"
	expectedPrecision := uint(16)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_14(t *testing.T) {

	numStr := "5 2"
	expectedNumStr := "52"
	expectedPrecision := uint(0)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_15(t *testing.T) {

	numStr := "    (52)     "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_16(t *testing.T) {

	numStr := "    (52)    1234567 "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(numStr)

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

func TestBigIntFixedDecimal_NewNumStr_17(t *testing.T) {

	expectedPrecision := uint(1050)

	fixedDec, err := BigIntFixedDecimal{}.NewNumStr(EulersNum1050)

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

func TestBigIntFixedDecimal_NewNumStr_18(t *testing.T) {

	numStr := "abcdefghijklmnop"

	_, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err == nil {
		t.Error("Expected an Error from INVALID Number String. NO ERROR RETURNED!" )
	}
}

