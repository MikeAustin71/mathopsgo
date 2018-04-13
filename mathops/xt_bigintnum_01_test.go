package mathops

import (
	"testing"
	"math/big"
	"errors"
)

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

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

}

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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
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

	actualNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStr(). Error='%v'", err.Error())
	}

	if actualNumStr != expectedNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			actualNumStr, expectedNumStr)
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Error: Expected Sign Value='%v'.  Instead, Sign Value='%v'.",
			expectedSignVal, bINum.Sign)
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Error: Expected Precision='%v'.  Instead, Precision='%v'.",
			expectedPrecision, bINum.Precision)
	}

	if bExpected.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.BigInt.Text(10))
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

	actualNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStr(). Error='%v'", err.Error())
	}

	if expectedNumStr != actualNumStr  {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Error: Expected Sign Value='%v'.  Instead, Sign Value='%v'.",
			expectedSignVal, bINum.Sign)
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Error: Expected Precision='%v'.  Instead, Precision='%v'.",
			expectedPrecision, bINum.Precision)
	}

	if bExpected.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.BigInt.Text(10))
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

	actualNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStr(). Error='%v'", err.Error())
	}

	if actualNumStr != expectedNumStr {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			actualNumStr, expectedNumStr)
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Error: Expected Sign Value='%v'.  Instead, Sign Value='%v'.",
			expectedSignVal, bINum.Sign)
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Error: Expected Precision='%v'.  Instead, Precision='%v'.",
			expectedPrecision, bINum.Precision)
	}

	if bExpected.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.BigInt.Text(10))
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

	actualNumStr, err := bINum.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by bINum.GetNumStr(). Error='%v'", err.Error())
	}

	if expectedNumStr != actualNumStr  {
		t.Errorf("Error: Expected NumStr='%v' Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Error: Expected Sign Value='%v'.  Instead, Sign Value='%v'.",
			expectedSignVal, bINum.Sign)
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Error: Expected Precision='%v'.  Instead, Precision='%v'.",
			expectedPrecision, bINum.Precision)
	}

	if bExpected.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Error: Expected BigIntNum='%v'.  Instead, BigIntNum='%v'",
			bExpected.Text(10), bINum.BigInt.Text(10))
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
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) " +
			"Error='%v' ", err.Error())
	}

	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}


	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}


	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}


	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
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

	nDto, err := NumStrDto{}.NewBigInt(bINum.BigInt, bINum.Precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NumStrDto{}.NewNumStr(nStr) " +
			"Error='%v' ", err.Error())
	}


	if bOriginal.Cmp(bINum.BigInt) != 0 {
		t.Errorf("Expected BigInt='%v'  Instead, BigInt='%v'. ",
			bOriginal.Text(10), bINum.BigInt.Text(10))
	}

	if expectedPrecision != bINum.Precision {
		t.Errorf("Expected Precision='%v' Instead, Precision='%v' ",
			expectedPrecision, bINum.Precision)
	}

	if bINum.ScaleFactor.Cmp(expectedScale) != 0 {
		t.Errorf("Expected Scale Value='%v' Instead, Scale Value='%v' ",
			expectedScale.Text(10), bINum.ScaleFactor.Text(10))
	}

	if expectedAbsBigInt.Cmp(bINum.AbsBigInt) != 0 {
		t.Errorf("Expected AbsBigInt='%v'  Instead, AbsBigInt='%v'. ",
			expectedAbsBigInt.Text(10), bINum.AbsBigInt.Text(10))
	}

	if expectedSignVal != bINum.Sign {
		t.Errorf("Expected Sign Value='%v'. Instead, Sign Value='%v'. ",
			expectedSignVal, bINum.Sign)
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto.GetNumStr())
	}

}

