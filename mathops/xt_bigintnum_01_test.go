package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_BigInt_01(t *testing.T) {

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

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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

	nStr := "0.000123456"
	expectedPrecision := uint(9)
	nbI64 := int64(123456)
	expectedScale := big.NewInt(1000000000)
	expectedSignVal := 1

	bOriginal := big.NewInt(nbI64)

	expectedAbsBigInt := big.NewInt(0).Set(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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

	nStr := "-0.000123456"
	expectedPrecision := uint(9)
	n := int64(-123456)
	expectedScale := big.NewInt(1000000000)
	expectedSignVal := -1

	bOriginal := big.NewInt(n)

	expectedAbsBigInt := big.NewInt(0).Neg(bOriginal)

	bINum := BigIntNum{}.NewBigInt(bOriginal, expectedPrecision)

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	ceiling := bINum1.Ceiling()

	actualNumStr := ceiling.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected Ceiling NumStr='%v'. "+
			"Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

	if expectedPrecision != ceiling.GetPrecisionUint() {
		t.Errorf("Error: Expected Ceiling precision='%v' "+
			"Instead, precision='%v'",
			expectedPrecision, ceiling.GetPrecisionUint())
	}
}

func TestBigIntNum_ChangeSign_01(t *testing.T) {
	nStr := "123.456"
	expectedStr := "-123.456"
	expectedSign := -1

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	bINum.ChangeSign()

	if expectedSign != bINum.GetSign() {
		t.Errorf("Error: After ChangeSign - Expected Sign='%v'.  Instead, Sign='%v' ",
			expectedSign, bINum.GetSign())
	}

	if expectedStr != bINum.GetNumStr() {
		t.Errorf("Error: After ChangeSign - Expected BigIntNum='%v'. Instead, BigIntNum='%v'",
			expectedStr, bINum.GetNumStr())
	}

}

func TestBigIntNum_ChangeSign_02(t *testing.T) {
	nStr := "-123.456"
	expectedStr := "123.456"
	expectedSign := 1

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	bINum.ChangeSign()

	if expectedSign != bINum.GetSign() {
		t.Errorf("Error: After ChangeSign - Expected Sign='%v'.  Instead, Sign='%v' ",
			expectedSign, bINum.GetSign())
	}

	if expectedStr != bINum.GetNumStr() {
		t.Errorf("Error: After ChangeSign - Expected BigIntNum='%v'. Instead, BigIntNum='%v'",
			expectedStr, bINum.GetNumStr())
	}

}

func TestBigIntNum_ChangeSign_03(t *testing.T) {
	nStr := "0.00"
	expectedStr := "0.00"
	expectedSign := 1

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	bINum.ChangeSign()

	if expectedSign != bINum.GetSign() {
		t.Errorf("Error: After ChangeSign - Expected Sign='%v'.  Instead, Sign='%v' ",
			expectedSign, bINum.GetSign())
	}

	if expectedStr != bINum.GetNumStr() {
		t.Errorf("Error: After ChangeSign - Expected BigIntNum='%v'. Instead, BigIntNum='%v'",
			expectedStr, bINum.GetNumStr())
	}

}

func TestBigIntNum_Cmp_01(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.455"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_02(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.457"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_03(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.456"
	expectedCmpResult := 0

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_04(t *testing.T) {

	n1Str := "-123.456"
	n2Str := "-123.457"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_05(t *testing.T) {

	n1Str := "-123.456"
	n2Str := "-123.455"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_06(t *testing.T) {

	n1Str := "-123.456"
	n2Str := "-123.456"
	expectedCmpResult := 0

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_07(t *testing.T) {

	n1Str := "123456000643218"
	n2Str := "123456000643217"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_08(t *testing.T) {

	n1Str := "123456000643218"
	n2Str := "123456000643219"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_09(t *testing.T) {

	n1Str := "123456000643218"
	n2Str := "123456000643218"
	expectedCmpResult := 0

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_10(t *testing.T) {

	n1Str := "-123456000643218"
	n2Str := "-123456000643217"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_11(t *testing.T) {

	n1Str := "-123456000643218"
	n2Str := "-123456000643219"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_12(t *testing.T) {

	n1Str := "-123456000643218"
	n2Str := "-123456000643218"
	expectedCmpResult := 0

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_13(t *testing.T) {

	n1Str := "1.7"
	n2Str := "-12345.6000643218"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_14(t *testing.T) {

	n1Str := "-1.7"
	n2Str := "0.01"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_15(t *testing.T) {

	n1Str := "17"
	n2Str := "-1"
	expectedCmpResult := 1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_16(t *testing.T) {

	n1Str := "-17"
	n2Str := "1"
	expectedCmpResult := -1

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Cmp_17(t *testing.T) {

	n1Str := "0"
	n2Str := "0.000"
	expectedCmpResult := 0

	bINum1, err := BigIntNum{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	bINum2, err := BigIntNum{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := bINum1.Cmp(bINum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestBigIntNum_Decimal_01(t *testing.T) {

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

	dec, err := Decimal{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewDecimal(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewDecimal(dec) "+
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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

	dec, err := Decimal{}.NewNumStr(nStr)

	bINum, err := BigIntNum{}.NewDecimal(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewIntAry(ia) "+
			"Error='%v' ", err.Error())
	}

	nDto, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision) "+
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

func TestBigIntNum_Decrement_01(t *testing.T) {

	numStr := "8"
	expectedNumStr := "7"
	expectedNumSeps := NumericSeparatorDto{}.New()

	bINum, err := BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
			"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	bINum.Decrement()

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_Decrement_02(t *testing.T) {

	numStr := "8.2"
	expectedNumStr := "7.2"
	expectedNumSeps := NumericSeparatorDto{}.New()

	bINum, err := BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
			"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	bINum.Decrement()

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_Decrement_03(t *testing.T) {

	numStr := "-8.2"
	expectedNumStr := "-9.2"
	expectedNumSeps := NumericSeparatorDto{}.New()

	bINum, err := BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
			"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	bINum.Decrement()

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := bINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntNum_Decrement_04(t *testing.T) {

	numStr := "8"
	expectedNumStr := "7"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	bINum, err := BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(numStr, expectedNumSeps). "+
			"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	bINum.Decrement()

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, actualNumStr)
	}

	actualNumSeps := bINum.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestBigIntNum_DivideByTwo_01(t *testing.T) {

	numStr := "658.78562347"
	// expectedNumStr := "329.392811735"
	expectedQuoStr := "329"
	expectedModStr := "0.78562347"

	maxPrecision := uint(10)

	base, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

	if err != nil {
		t.Errorf("Error returned by base.DivideByTwoQuoMod(maxPrecision). Error='%v' ",
			err.Error())
	}

	actualNumStr := quotient.GetNumStr()

	if expectedQuoStr != actualNumStr {
		t.Errorf("Error: Expected quotient NumStr='%v'. Instead, quotient NumStr='%v'. ",
			expectedQuoStr, actualNumStr)
	}

	actualNumStr = modulo.GetNumStr()

	if expectedModStr != actualNumStr {
		t.Errorf("Error: Expected modulo NumStr='%v'. Instead, modulo NumStr='%v'. ",
			expectedModStr, actualNumStr)
	}
}

func TestBigIntNum_DivideByTwo_02(t *testing.T) {

	numStr := "8"

	expectedQuoStr := "4"
	expectedModStr := "0"

	maxPrecision := uint(10)

	base, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	quotient, modulo, err := base.DivideByTwoQuoMod(maxPrecision)

	if err != nil {
		t.Errorf("Error returned by base.DivideByTwoQuoMod(maxPrecision). Error='%v' ",
			err.Error())
	}

	actualNumStr := quotient.GetNumStr()

	if expectedQuoStr != actualNumStr {
		t.Errorf("Error: Expected quotient NumStr='%v'. Instead, quotient NumStr='%v'. ",
			expectedQuoStr, actualNumStr)
	}

	actualNumStr = modulo.GetNumStr()

	if expectedModStr != actualNumStr {
		t.Errorf("Error: Expected modulo NumStr='%v'. Instead, modulo NumStr='%v'. ",
			expectedModStr, actualNumStr)
	}
}

func TestBigIntNum_DivideByTenToPower_01(t *testing.T) {
	num1Str := "654.123"
	expectedStr := "0.654123"
	exponent := uint(3)

	bINum, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	bINum.DivideByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_DivideByTenToPower_02(t *testing.T) {
	num1Str := "-654.123"
	expectedStr := "-0.654123"
	exponent := uint(3)

	bINum, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	bINum.DivideByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_DivideByTenToPower_03(t *testing.T) {
	num1Str := "654123"
	expectedStr := "0.000654123"
	exponent := uint(9)

	bINum, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	bINum.DivideByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_DivideByTenToPower_04(t *testing.T) {
	num1Str := "654123"
	expectedStr := "654123"
	exponent := uint(0)

	bINum, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	bINum.DivideByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_DivideByTenToPower_05(t *testing.T) {
	num1Str := "-654123"
	expectedStr := "-654123"
	exponent := uint(0)

	bINum, err := BigIntNum{}.NewNumStr(num1Str)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(num1Str). Error='%v' ",
			err.Error())
	}

	bINum.DivideByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
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

	bNum1.ExtendPrecision(2)

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

	bNum1.ExtendPrecision(2)

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

	bNum1.ExtendPrecision(2)

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

	bNum1.ExtendPrecision(2)

	if !expectedNum.Equal(bNum1) {
		t.Errorf("Error: Expected bNum1='%v'.  Instead, bNum2='%v'",
			expectedNum.GetNumStr(), bNum1.GetNumStr())
	}

}
