package mathops

import (
	"testing"
	"math/big"
	"errors"
)


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

func TestBigIntNum_RoundToDecPlace_01(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.57"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.57"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_08(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_09(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_10(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_11(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_12(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_13(t *testing.T) {

	nStr := "-654.123"
	expectedNumStr := "-654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_RoundToDecPlace_14(t *testing.T) {

	nStr := "-654"
	expectedNumStr := "-654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}


func TestBigIntNum_RoundToDecPlace_15(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_SetPrecision_01(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.123"
	newPrecision := uint(3)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_02(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.1235"
	newPrecision := uint(4)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_03(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654"
	newPrecision := uint(0)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_04(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.123"
	newPrecision := uint(3)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_05(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.1235"
	newPrecision := uint(4)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_06(t *testing.T) {

	num1Str := "654"
	expectedNumStr := "654.000"
	newPrecision := uint(3)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_07(t *testing.T) {

	num1Str := "654.123456"
	expectedNumStr := "654.123456000"
	newPrecision := uint(9)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_08(t *testing.T) {

	num1Str := "-654"
	expectedNumStr := "-654.000000000"
	newPrecision := uint(9)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_09(t *testing.T) {

	num1Str := "-654.123456"
	expectedNumStr := "-654.123456000"
	newPrecision := uint(9)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_10(t *testing.T) {

	num1Str := "0"
	expectedNumStr := "0.0000"
	newPrecision := uint(4)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_SetPrecision_11(t *testing.T) {

	num1Str := "0.0000"
	expectedNumStr := "0"
	newPrecision := uint(0)

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

	bNum1.SetPrecision(newPrecision)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum1='%v'" +
			"precision='%v' sign='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr(), bNum1.GetPrecisionUint(), bNum1.GetSign())
	}

	if newPrecision != bNum1.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v' .",
			newPrecision, bNum1.GetPrecisionUint())
	}

}

func TestBigIntNum_TruncToDecPlace_01(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.56"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}


func TestBigIntNum_TruncToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	truncToPlace := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToPlace)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	truncToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	truncToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_08(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654.1234"
	truncToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_09(t *testing.T) {

	nStr := "654.123456789"
	expectedNumStr := "654"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_10(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.00000"
	truncToDec := uint(5)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_11(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	truncToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_12(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.000000"
	truncToDec := uint(6)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_13(t *testing.T) {

	nStr := "0.000000"
	expectedNumStr := "0"
	truncToDec := uint(0)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

func TestBigIntNum_TruncToDecPlace_14(t *testing.T) {

	nStr := "654.123456789015"
	expectedNumStr := "654.12345678901"
	truncToDec := uint(11)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.TruncToDecPlace(truncToDec)

	actualNumStr, err := bINum1.GetNumStrErr()

	if err != nil {
		t.Errorf("Error returned by bINum1.GetNumStrErr1(). Error='%v'",
			err.Error())
	}

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}

