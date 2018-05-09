package mathops

import (
	"testing"
	"math/big"
	"errors"
)

func TestBigIntNum_BigInt_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}


func TestBigIntNum_BigInt_02(t *testing.T) {

	nStr:="-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_BigInt_03(t *testing.T) {

	nStr:="0.000123456"
	expectedPrecision := uint(9)
	nbI64 := int64(123456)
	expectedScale := big.NewInt(1000000000)
	expectedSignVal := 1

	bOriginal := big.NewInt(nbI64)

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}


func TestBigIntNum_BigInt_04(t *testing.T) {

	nStr:="-0.000123456"
	expectedPrecision := uint(9)
	n := int64(-123456)
	expectedScale := big.NewInt(1000000000)
	expectedSignVal := -1

	bOriginal := big.NewInt(n)

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_Ceil_01(t *testing.T) {
	nStr := "5.95"
	expectedNumStr := "6"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_02(t *testing.T) {
	nStr := "5.05"
	expectedNumStr := "6"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_03(t *testing.T) {
	nStr := "5"
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_04(t *testing.T) {
	nStr := "-5.05"
	expectedNumStr := "-5"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_05(t *testing.T) {
	nStr := "2.4"
	expectedNumStr := "3"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_06(t *testing.T) {
	nStr := "2.9"
	expectedNumStr := "3"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_07(t *testing.T) {
	nStr := "-2.7"
	expectedNumStr := "-2"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_08(t *testing.T) {
	nStr := "-2"
	expectedNumStr := "-2"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_09(t *testing.T) {
	nStr := "0"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_10(t *testing.T) {
	nStr := "0.00000"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_12(t *testing.T) {
	nStr := "159876231.9999999999"
	expectedNumStr := "159876232"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_13(t *testing.T) {
	nStr := "-159876231.9999999999"
	expectedNumStr := "-159876231"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_14(t *testing.T) {
	nStr := "159876231.0000000000000001"
	expectedNumStr := "159876232"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_15(t *testing.T) {
	nStr := "-159876231.0000000000000001"
	expectedNumStr := "-159876231"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_16(t *testing.T) {
	nStr := "-0.0000000000000001"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Ceil_17(t *testing.T) {
	nStr := "0.0000000000000001"
	expectedNumStr := "1"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr, err := ceiling.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by ceiling.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_Decimal_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	dec, err := Decimal{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewDecimal(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewDecimal(dec) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_Decimal_02(t *testing.T) {

	nStr:="-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	dec, err := Decimal{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewDecimal(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewIntAry(ia) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_ExtendPrecision_01(t *testing.T) {

	num1Str := "654.123"
	expectedNumStr := "654.12300"

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.extendPrecision(2)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

}

func TestBigIntNum_ExtendPrecision_02(t *testing.T) {

	num1Str := "-654.123"
	expectedNumStr := "-654.12300"

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.extendPrecision(2)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

}

func TestBigIntNum_ExtendPrecision_03(t *testing.T) {

	num1Str := "7"
	expectedNumStr := "7.00"

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.extendPrecision(2)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

}

func TestBigIntNum_ExtendPrecision_04(t *testing.T) {

	num1Str := "0"
	expectedNumStr := "0.00"

	bNum1, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	expectedNum, err := BigIntNum{}.NewNumStr(expectedNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedNumStr). Error='%v' ",
			err.Error())
	}

	bNum1.extendPrecision(2)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

	actualNumStr, err := floor.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by floor.GetNumStrErr(). Error='%v'",
			err.Error())
	}

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

func TestBigIntNum_GetFractionalPart_01(t *testing.T) {

	nStr := "123.456"
	expectedNumStr := "0.456"
	expectedPrecision := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetFractionalPart_02(t *testing.T) {

	nStr := "-123.456"
	expectedNumStr := "-0.456"
	expectedPrecision := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetFractionalPart_03(t *testing.T) {

	nStr := "123"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetFractionalPart_04(t *testing.T) {

	nStr := "-123"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetFractionalPart_05(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetFractionalPart_06(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	fractionalPart := bINum1.GetFractionalPart()

	actualNumStr, err := fractionalPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by fractionalPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Fractional Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != fractionalPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Fractional Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, fractionalPart.GetPrecisionUint())
	}
}

func TestBigIntNum_GetDecimal_01(t *testing.T) {
	expectedStr := "-847921684.347"
	ePrecision := uint(3)

	bINum, err := BigIntNum{}.NewNumStr(expectedStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStr). Error='%v'", err.Error())
	}

	decActual, err := bINum.GetDecimal()

	if expectedStr != decActual.GetNumStr() {
		t.Errorf("Error: Expected decActual.GetNumStr()='%v'. " +
			"Instead, decActual.GetNumStr()='%v'.",
			expectedStr, decActual.GetNumStr())
	}

	if ePrecision != decActual.GetPrecisionUint() {
		t.Errorf("Error: Expected decActual.GetPrecisionUint()='%v'. " +
			"Instead, decActual.GetPrecisionUint()='%v'.",
			ePrecision, decActual.GetPrecisionUint())
	}

}

func TestBigIntNum_GetNumStrDto_01(t *testing.T) {
	expectedStr := "-847921684.347"
	ePrecision := uint(3)

	bINum, err := BigIntNum{}.NewNumStr(expectedStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStr). Error='%v'", err.Error())
	}

	numStrDto, err := bINum.GetNumStrDto()

	if expectedStr != numStrDto.GetNumStr() {
		t.Errorf("Error: Expected numStrDto.GetNumStr()='%v'. " +
			"Instead, numStrDto.GetNumStr()='%v'.",
			expectedStr, numStrDto.GetNumStr())
	}

	if ePrecision != numStrDto.GetPrecisionUint() {
		t.Errorf("Error: Expected numStrDto.GetPrecisionUint()='%v'. " +
			"Instead, numStrDto.GetPrecisionUint()='%v'.",
			ePrecision, numStrDto.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntAry_01(t *testing.T) {
	expectedStr := "-847921684.347"
	ePrecision := uint(3)

	bINum, err := BigIntNum{}.NewNumStr(expectedStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(expectedStr). Error='%v'", err.Error())
	}

	intAry, err := bINum.GetIntAry()

	if expectedStr != intAry.GetNumStr() {
		t.Errorf("Error: Expected intAry.GetNumStr()='%v'. " +
			"Instead, intAry.GetNumStr()='%v'.",
			expectedStr, intAry.GetNumStr())
	}

	if ePrecision != intAry.GetPrecisionUint() {
		t.Errorf("Error: Expected intAry.GetPrecisionUint()='%v'. " +
			"Instead, intAry.GetPrecisionUint()='%v'.",
			ePrecision, intAry.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_01(t *testing.T) {

	nStr := "123.456"
	expectedNumStr := "123"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_02(t *testing.T) {

	nStr := "-123.456"
	expectedNumStr := "-123"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_03(t *testing.T) {

	nStr := "123"
	expectedNumStr := "123"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_04(t *testing.T) {

	nStr := "-123"
	expectedNumStr := "-123"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_05(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_GetIntegerPart_06(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0"
	expectedPrecision := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	integerPart := bINum1.GetIntegerPart()

	actualNumStr, err := integerPart.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by integerPart.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Integer Part NumStr='%v'. " +
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

	if expectedPrecision != integerPart.GetPrecisionUint() {
		t.Errorf("Error: Expected Integer Part precision='%v' " +
			"Instead, precision='%v'",
			expectedPrecision, integerPart.GetPrecisionUint())
	}

}

func TestBigIntNum_IntAry_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	ia, err := IntAry{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewIntAry(ia)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewIntAry(ia) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_IntAry_02(t *testing.T) {

	nStr:="-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	ia, err := IntAry{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewIntAry(ia)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewIntAry(ia) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_NewINumMgr_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	dec, err := Decimal{}.NewNumStr(nStr)

	if err!= nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_NewINumMgr_02(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	ia, err := IntAry{}.NewNumStr(nStr)

	if err!= nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&ia)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&ia) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}


func TestBigIntNum_NewINumMgr_03(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	nDto0, err := NumStrDto{}.NewNumStr(nStr)

	if err!= nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&nDto0)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&nDto0) " +
			"Error='%v' ", err.Error())
	}

	nDto1, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto1.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto1.GetNumStr())
	}

}

func TestBigIntNum_NewINumMgr_04(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	dec, err := Decimal{}.NewNumStr(nStr)

	if err!= nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(dec.GetThisPointer())

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}


func TestBigIntNum_NewINumMgr_05(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	dec := Decimal{}.NewPtr()
	err := dec.SetNumStr(nStr)

	if err!= nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) " +
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_NewNumStr_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
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

	nStr:="-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
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

func TestBigIntNum_NewBigIntExponent_01(t *testing.T) {

	n:= 123456
	exponent := 3
	expectedPrecision := uint(0)
	expectedNumStr := "123456000"
	bExpected := big.NewInt(int64(123456000))
	expectedSignVal := 1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr, err := bINum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStrErr(). Error='%v'", err.Error())
	}

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

	n:= 123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "123.456"
	bExpected := big.NewInt(int64(123456))
	expectedSignVal := 1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr, err := bINum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStrErr(). Error='%v'", err.Error())
	}

	if expectedNumStr != actualNumStr  {
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

	n:= -123456
	exponent := 3
	expectedPrecision := uint(0)
	expectedNumStr := "-123456000"
	bExpected := big.NewInt(int64(-123456000))
	expectedSignVal := -1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr, err := bINum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStrErr(). Error='%v'", err.Error())
	}

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

	n:= -123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "-123.456"
	bExpected := big.NewInt(int64(-123456))
	expectedSignVal := -1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr, err := bINum.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStrErr(). Error='%v'", err.Error())
	}

	if expectedNumStr != actualNumStr  {
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

func TestBigIntNum_NumStrDto_01(t *testing.T) {

	nStr:="123.456"
	expectedPrecision := uint(3)
	nbStr := "123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := 1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

func TestBigIntNum_NumStrDto_02(t *testing.T) {

	nStr:="-123.456"
	expectedPrecision := uint(3)
	nbStr := "-123456"
	expectedScale := big.NewInt(1000)
	expectedSignVal := -1

	bOriginal, isOk := big.NewInt(0).SetString(nbStr, 10)

	if !isOk {
		errors.New("Error returned by big.NewInt(0).SetString(nbStr, 10).")
	}

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) " +
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

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}
