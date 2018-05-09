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
