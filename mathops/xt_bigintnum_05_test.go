package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_NewNumStr_01(t *testing.T) {

	nStr := "123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Expected bigInt='%v'  Instead, bigInt='%v'. ",
			bOriginal.Text(10), bINum.bigInt.Text(10))
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Expected precision='%v' Instead, precision='%v' ",
			expectedPrecision, bINum.precision)
	}

	if bINum.scaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.scaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.absBigInt) != 0 {
		t.Errorf("Expected absBigInt='%v'  Instead, absBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.absBigInt.Text(10))
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Expected sign Value='%v'. Instead, sign Value='%v'. ",
			expectedSignVal, bINum.sign)
	}

}

func TestBigIntNum_NewNumStr_02(t *testing.T) {

	nStr := "-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Expected bigInt='%v'  Instead, bigInt='%v'. ",
			bOriginal.Text(10), bINum.bigInt.Text(10))
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Expected precision='%v' Instead, precision='%v' ",
			expectedPrecision, bINum.precision)
	}

	if bINum.scaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.scaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.absBigInt) != 0 {
		t.Errorf("Expected absBigInt='%v'  Instead, absBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.absBigInt.Text(10))
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Expected sign Value='%v'. Instead, sign Value='%v'. ",
			expectedSignVal, bINum.sign)
	}

}


func TestBigIntNum_NewNumStr_03(t *testing.T) {

	numStr := "0.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_04(t *testing.T) {

	numStr := ".123456789012"
	expectedNumStr := "0.123456789012"
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_05(t *testing.T) {

	numStr := "-.123456789012"
	expectedNumStr := "-0.123456789012"
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_06(t *testing.T) {

	numStr := "10"
	expectedNumStr := "10"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_07(t *testing.T) {

	numStr := "-52"
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_08(t *testing.T) {

	numStr := "-00052.1234"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_09(t *testing.T) {

	numStr := "(00052.1234)"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_10(t *testing.T) {

	numStr := "(00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_11(t *testing.T) {

	numStr := "+00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_12(t *testing.T) {

	numStr := "-00052.1234567890123456"
	expectedNumStr := "-52.1234567890123456"
	expectedPrecision := uint(16)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_13(t *testing.T) {

	numStr := "52 . 123 4567 8901 23456"
	expectedNumStr := "52.1234567890123456"
	expectedPrecision := uint(16)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_14(t *testing.T) {

	numStr := "5 2"
	expectedNumStr := "52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_15(t *testing.T) {

	numStr := "    (52)     "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_16(t *testing.T) {

	numStr := "    (52)    1234567 "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(numStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if expectedNumStr != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedNumStr, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}
}

func TestBigIntNum_NewNumStr_17(t *testing.T) {

	expectedPrecision := uint(1050)

	bigINum, err := BigIntFixedDecimal{}.NewNumStr(EulersNum1050Str)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if EulersNum1050Str != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			EulersNum1050Str, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecision() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())

	}

}

func TestBigIntNum_NewNumStr_18(t *testing.T) {

	numStr := "abcdefghijklmnop"

	_, err := BigIntFixedDecimal{}.NewNumStr(numStr)

	if err == nil {
		t.Error("Expected an Error from INVALID Number String. NO ERROR RETURNED!" )
	}
}


func TestBigIntNum_NewNumStrWithNumSeps_01(t *testing.T) {

	nStr := "123,456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	bigINum, err := BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). "+
			"Error='%v'", err.Error())
	}

	actualNumStr := bigINum.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}
}

func TestBigIntNum_NewNumStrWithNumSeps_02(t *testing.T) {

	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.SetDefaultsIfEmpty()

	bigINum, err := BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). "+
			"Error='%v'", err.Error())
	}

	actualNumStr := bigINum.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}
}

func TestBigIntNum_NewBigIntExponent_01(t *testing.T) {

	n := 123456
	exponent := 3
	expectedPrecision := uint(3)
	expectedNumStr := "123456.000"
	bExpected := big.NewInt(int64(123456000))
	expectedSignVal := 1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

	if actualNumStr != expectedNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			actualNumStr, expectedNumStr)
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Error: Expected sign Value='%v'.  Instead, sign Value='%v'.",
			expectedSignVal, bINum.sign)
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Error: Expected precision='%v'.  Instead, precision='%v'.",
			expectedPrecision, bINum.precision)
	}

	if bExpected.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.bigInt.Text(10))
	}

}

func TestBigIntNum_NewBigIntExponent_02(t *testing.T) {

	n := 123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "123.456"
	bExpected := big.NewInt(int64(123456))
	expectedSignVal := 1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Error: Expected sign Value='%v'.  Instead, sign Value='%v'.",
			expectedSignVal, bINum.sign)
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Error: Expected precision='%v'.  Instead, precision='%v'.",
			expectedPrecision, bINum.precision)
	}

	if bExpected.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.bigInt.Text(10))
	}

}

func TestBigIntNum_NewBigIntExponent_03(t *testing.T) {

	n := -123456
	exponent := 3
	expectedPrecision := uint(3)
	expectedNumStr := "-123456.000"
	bExpected := big.NewInt(int64(-123456000))
	expectedSignVal := -1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

	if actualNumStr != expectedNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			actualNumStr, expectedNumStr)
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Error: Expected sign Value='%v'.  Instead, sign Value='%v'.",
			expectedSignVal, bINum.sign)
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Error: Expected precision='%v'.  Instead, precision='%v'.",
			expectedPrecision, bINum.precision)
	}

	if bExpected.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.bigInt.Text(10))
	}

}

func TestBigIntNum_NewBigIntExponent_04(t *testing.T) {

	n := -123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "-123.456"
	bExpected := big.NewInt(int64(-123456))
	expectedSignVal := -1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	if expectedSignVal != bINum.sign {
		t.Errorf("Error: Expected sign Value='%v'.  Instead, sign Value='%v'.",
			expectedSignVal, bINum.sign)
	}

	if expectedPrecision != bINum.precision {
		t.Errorf("Error: Expected precision='%v'.  Instead, precision='%v'.",
			expectedPrecision, bINum.precision)
	}

	if bExpected.Cmp(bINum.bigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.bigInt.Text(10))
	}

}


func TestBigIntNum_NewOne_01(t *testing.T) {
	expectedNumStr := "1.000"
	expectedPrecision := uint(3)

	bINum := BigIntNum{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewOne_02(t *testing.T) {
	expectedNumStr := "1"
	expectedPrecision := uint(0)

	bINum := BigIntNum{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewOne_03(t *testing.T) {
	expectedNumStr := "1.00000"
	expectedPrecision := uint(5)

	bINum := BigIntNum{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewTwo_01(t *testing.T) {
	expectedNumStr := "2.000"
	expectedPrecision := uint(3)

	bINum := BigIntNum{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewTwo_02(t *testing.T) {
	expectedNumStr := "2"
	expectedPrecision := uint(0)

	bINum := BigIntNum{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewTwo_03(t *testing.T) {
	expectedNumStr := "2.00000"
	expectedPrecision := uint(5)

	bINum := BigIntNum{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewThree_01(t *testing.T) {
	expectedNumStr := "3.000"
	expectedPrecision := uint(3)

	bINum := BigIntNum{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewThree_02(t *testing.T) {
	expectedNumStr := "3"
	expectedPrecision := uint(0)

	bINum := BigIntNum{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewThree_03(t *testing.T) {
	expectedNumStr := "3.00000"
	expectedPrecision := uint(5)

	bINum := BigIntNum{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewFive_01(t *testing.T) {
	expectedNumStr := "5.000"
	expectedPrecision := uint(3)

	bINum := BigIntNum{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewFive_02(t *testing.T) {
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum := BigIntNum{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewFive_03(t *testing.T) {
	expectedNumStr := "5.00000"
	expectedPrecision := uint(5)

	bINum := BigIntNum{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}
}

func TestBigIntNum_NewTen_01(t *testing.T) {
	expectedNumStr := "10.000"
	expectedPrecision := uint(3)

	bINum := BigIntNum{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewTen_02(t *testing.T) {
	expectedNumStr := "10"
	expectedPrecision := uint(0)

	bINum := BigIntNum{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewTen_03(t *testing.T) {
	expectedNumStr := "10.00000"
	expectedPrecision := uint(5)

	bINum := BigIntNum{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestBigIntNum_NewUint_01(t *testing.T) {

	numInt := uint(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewUint(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint_02(t *testing.T) {

	numInt := uint(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewUint(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint_03(t *testing.T) {

	numInt := uint(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUint(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint_04(t *testing.T) {

	numInt := uint(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_01(t *testing.T) {

	numInt := uint(1234)
	exponent := 3
	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_02(t *testing.T) {

	numInt := uint(123456)
	exponent := -2
	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_03(t *testing.T) {

	numInt := uint(123456)
	exponent := 0
	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_04(t *testing.T) {

	numInt := uint(0)
	exponent := 0
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_05(t *testing.T) {

	numInt := uint(0)
	exponent := 3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUintExponent_06(t *testing.T) {

	numInt := uint(0)
	exponent := -3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUintExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32_01(t *testing.T) {

	num32Uint := uint32(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewUint32(num32Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32_02(t *testing.T) {

	num32Uint := uint32(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewUint32(num32Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32_03(t *testing.T) {

	num32Uint := uint32(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUint32(num32Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32_04(t *testing.T) {

	num32Uint := uint32(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint32(num32Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_01(t *testing.T) {

	numInt := uint32(1234)
	exponent := 3
	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_02(t *testing.T) {

	numInt := uint32(123456)
	exponent := -2
	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_03(t *testing.T) {

	numInt := uint32(123456)
	exponent := 0
	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_04(t *testing.T) {

	numInt := uint32(0)
	exponent := 0
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_05(t *testing.T) {

	numInt := uint32(0)
	exponent := 3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint32Exponent_06(t *testing.T) {

	numInt := uint32(0)
	exponent := -3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64_01(t *testing.T) {

	num64Uint := uint64(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewUint64(num64Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64_02(t *testing.T) {

	num64Uint := uint64(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewUint64(num64Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64_03(t *testing.T) {

	num64Uint := uint64(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUint64(num64Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64_04(t *testing.T) {

	num64Uint := uint64(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint64(num64Uint, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_01(t *testing.T) {

	numInt := uint64(1234)
	exponent := 3
	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_02(t *testing.T) {

	numInt := uint64(123456)
	exponent := -2
	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_03(t *testing.T) {

	numInt := uint64(123456)
	exponent := 0
	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_04(t *testing.T) {

	numInt := uint64(0)
	exponent := 0
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_05(t *testing.T) {

	numInt := uint64(0)
	exponent := 3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewUint64Exponent_06(t *testing.T) {

	numInt := uint64(0)
	exponent := -3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewUint64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

