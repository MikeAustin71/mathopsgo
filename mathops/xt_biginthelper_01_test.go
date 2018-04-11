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

func TestBigIntPair_NewBase_01(t *testing.T) {

	//n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	//n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}
func TestBigIntPair_NewBase_02(t *testing.T) {

	//n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	//n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewBigIntNum_01(t *testing.T) {

	//n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	//n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	bPair := BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewBigIntNum_02(t *testing.T) {

	//n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	//n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	bPair := BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStr_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	bPair, err := BigIntPair{}.NewNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStr(n1Str, n2Str). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStr_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	bPair, err := BigIntPair{}.NewNumStr(n1Str, n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStr(n1Str, n2Str). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStrDto_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	nx1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	nx2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewNumStrDto_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	nx1Dto, err := NumStrDto{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	nx2Dto, err := NumStrDto{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewNumStrDto(nx1Dto, nx2Dto). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewIntAry_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_NewIntAry_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	ia1, err := IntAry{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_Decimal_01(t *testing.T) {

	n1Str:="1234567.8901"
	num1Str := "12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := 1
	b1AbsBigInt := big.NewInt(0).Set(b1)

	n2Str:="7654321.95"
	num2Str := "765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := 1
	b2AbsBigInt := big.NewInt(0).Set(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="76543219500"

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
			"Error='%v'. ", err.Error())
	}


	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

func TestBigIntPair_Decimal_02(t *testing.T) {

	n1Str:="-1234567.8901"
	num1Str := "-12345678901"
	b1, oK := big.NewInt(0).SetString(num1Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num1Str, 10)")
	}

	b1Precision := uint(4)
	b1Sign := -1
	b1AbsBigInt := big.NewInt(0).Neg(b1)

	n2Str:="-7654321.95"
	num2Str := "-765432195"
	b2, oK := big.NewInt(0).SetString(num2Str, 10)

	if !oK {
		errors.New("Error returned from big.NewInt(0).SetString(num2Str, 10)")
	}

	b2Precision := uint(2)
	b2Sign := -1
	b2AbsBigInt := big.NewInt(0).Neg(b2)

	reconciledPrecision := uint(4)

	n2StrReconciled :="-76543219500"

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(n1Str). " +
			"Error='%v'. ", err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(n2Str). " +
			"Error='%v'. ", err.Error())
	}

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		t.Errorf("Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
			"Error='%v'. ", err.Error())
	}

	if bPair.Big1.BigInt.Cmp(b1) != 0 {
		t.Errorf("Error: Expected Big1.BigInt='%v'.  Instead Big1.BigInt='%v'.",
			b1.Text(10), bPair.Big1.BigInt.Text(10))
	}

	if b1AbsBigInt.Cmp(bPair.Big1.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big1.AbsBigInt='%v'.  Instead Big1.AbsBigInt='%v'.",
			b1AbsBigInt.Text(10), bPair.Big1.AbsBigInt.Text(10))
	}

	if b1Precision !=  bPair.Big1.Precision {
		t.Errorf("Error: Expected Big1.Precision='%v'.  Instead, Big1.Precision='%v'.",
			b1Precision, bPair.Big1.Precision)
	}

	if b1Sign != bPair.Big1.Sign {
		t.Errorf("Error: Expected Big1.Sign='%v'.  Instead, Big1.Sign='%v'.",
			b1Sign, bPair.Big1.Sign)
	}

	if bPair.Big2.BigInt.Cmp(b2) != 0 {
		t.Errorf("Error. Expected Big2.BigInt='%v'.  Instead Big2.BigInt='%v'.",
			b2.Text(10), bPair.Big2.BigInt.Text(10))
	}


	if b2AbsBigInt.Cmp(bPair.Big2.AbsBigInt) != 0 {
		t.Errorf("Error: Expected Big2.AbsBigInt='%v'.  Instead Big2.AbsBigInt='%v'.",
			b2AbsBigInt.Text(10), bPair.Big2.AbsBigInt.Text(10))
	}

	if b2Precision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected Big2.Precision='%v'.  Instead, Big2.Precision='%v'.",
			b2Precision, bPair.Big2.Precision)
	}

	if b2Sign != bPair.Big2.Sign {
		t.Errorf("Error: Expected Big2.Sign='%v'.  Instead, Big2.Sign='%v'.",
			b2Sign, bPair.Big2.Sign)
	}

	bPair.MakePrecisionsEqual()

	if reconciledPrecision !=  bPair.Big2.Precision {
		t.Errorf("Error: Expected reconciled Big2.Precision='%v'. " +
			" Instead, reconciled Big2.Precision='%v'.",
			reconciledPrecision, bPair.Big2.Precision)
	}

	actualBig2Numstr := bPair.Big2.BigInt.Text(10)

	if n2StrReconciled != actualBig2Numstr {
		t.Errorf("Error: Expected reconciled Big2.BigInt='%v'. " +
			"Instead, reconciled Big2.BigInt='%v'. ",
			n2StrReconciled, actualBig2Numstr)
	}
}

