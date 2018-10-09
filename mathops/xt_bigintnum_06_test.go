package mathops

import (
	"math/big"
	"testing"
)

func TestBigIntNum_NumStrDto_01(t *testing.T) {

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

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) "+
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

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
			"Error='%v' ", err.Error())
	}

	bINum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStrDto(nDto) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_02(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.57"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_03(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.567"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_04(t *testing.T) {

	nStr := "123.567"
	expectedNumStr := "123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_05(t *testing.T) {

	nStr := "-123.567"
	expectedNumStr := "-123.5670"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_06(t *testing.T) {

	nStr := "0.000"
	expectedNumStr := "0.00"
	roundToDec := uint(2)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_07(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_08(t *testing.T) {

	nStr := "654.123456"
	expectedNumStr := "654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_09(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.123"
	roundToDec := uint(3)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_10(t *testing.T) {

	nStr := "-654.123456"
	expectedNumStr := "-654.1235"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_11(t *testing.T) {

	nStr := "654"
	expectedNumStr := "654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_12(t *testing.T) {

	nStr := "654.123"
	expectedNumStr := "654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_13(t *testing.T) {

	nStr := "-654.123"
	expectedNumStr := "-654.123000000"
	roundToDec := uint(9)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_14(t *testing.T) {

	nStr := "-654"
	expectedNumStr := "-654.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_RoundToDecPlace_15(t *testing.T) {

	nStr := "0"
	expectedNumStr := "0.0000"
	roundToDec := uint(4)

	bINum1, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
			" nStr='%v'  Error='%v'",
			nStr, err.Error())
	}

	bINum1.RoundToDecPlace(roundToDec)

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_SetNumStr_01(t *testing.T) {

	origNumStr:= "57.64"
	numStr := "89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_02(t *testing.T) {

	origNumStr := "257.1"
	numStr := "-89765.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_03(t *testing.T) {

	origNumStr := "0.000005"
	numStr := "0.123456789012"
	expectedNumStr := numStr
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_04(t *testing.T) {

	origNumStr := "97.8"
	numStr := ".123456789012"
	expectedNumStr := "0.123456789012"
	expectedPrecision := uint(12)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_05(t *testing.T) {

	origNumStr := "87"
	numStr := "-.123456789012"
	expectedNumStr := "-0.123456789012"
	expectedPrecision := uint(12)


	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_06(t *testing.T) {
	origNumStr := "97.9821"
	numStr := "10"
	expectedNumStr := "10"
	expectedPrecision := uint(0)


	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_07(t *testing.T) {

	origNumStr := "9845.61"
	numStr := "-52"
	expectedNumStr := "-52"
	expectedPrecision := uint(0)


	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_08(t *testing.T) {

	origNumStr := "97"
	numStr := "-00052.1234"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_09(t *testing.T) {

	origNumStr := "87.123456"
	numStr := "(00052.1234)"
	expectedNumStr := "-52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_10(t *testing.T) {

	origNumStr:="22.414141414"
	numStr := "(00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_11(t *testing.T) {
	origNumStr := "98.123456"
	numStr := "+00052.1234"
	expectedNumStr := "52.1234"
	expectedPrecision := uint(4)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_12(t *testing.T) {

	origNumStr := "97782345646"
	numStr := "-00052.1234567890123456"
	expectedNumStr := "-52.1234567890123456"
	expectedPrecision := uint(16)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_13(t *testing.T) {
	origNumStr := "52"
	numStr := "52 . 123 4567 8901 23456"
	expectedNumStr := "52.1234567890123456"
	expectedPrecision := uint(16)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_14(t *testing.T) {

	origNumStr := "987.123456"
	numStr := "5 2"
	expectedNumStr := "52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_15(t *testing.T) {

	origNumStr := "-8794521.12345"
	numStr := "    (52)     "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_16(t *testing.T) {

	origNumStr := "98"
	numStr := "    (52)    1234567 "
	expectedNumStr := "-52"
	expectedPrecision := uint(0)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

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

func TestBigIntNum_SetNumStr_17(t *testing.T) {

	origNumStr := "98123456789"
	expectedPrecision := uint(1050)

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	bigINum.SetNumStr(EulersNum1050Str)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	if EulersNum1050Str != bigINum.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			EulersNum1050Str, bigINum.GetNumStr())
	}

	if expectedPrecision != bigINum.GetPrecisionUint() {
		t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
			expectedPrecision, bigINum.GetPrecision())
	}

}


func TestBigIntNum_SetNumStr_18(t *testing.T) {

	origNumStr := "98"
	numStr := "abcdefghijklmnop"

	bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

	if err!= nil {
		t.Errorf("Error returned by %v", err.Error() )
	}

	err = bigINum.SetNumStr(numStr)

	if err == nil {
		t.Error("Expected an Error from INVALID Number String. NO ERROR RETURNED!" )
	}
}


func TestBigIntNum_ShiftPrecisionLeft_01(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "123.456789"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_02(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "1234.56789"
	shiftPlacesLeft := uint(2)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_03(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "0.123456789"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_04(t *testing.T) {
	basicNumStr := "123456789"
	expectedResult := "123.456789"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_05(t *testing.T) {
	basicNumStr := "123"
	expectedResult := "0.00123"
	shiftPlacesLeft := uint(5)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_06(t *testing.T) {
	basicNumStr := "0"
	expectedResult := "0"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_07(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "123456.789"
	shiftPlacesLeft := uint(0)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_08(t *testing.T) {
	basicNumStr := "-123456.789"
	expectedResult := "-123456.789"
	shiftPlacesLeft := uint(0)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_09(t *testing.T) {
	basicNumStr := "-123456.789"
	expectedResult := "-123.456789"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionLeft_10(t *testing.T) {
	basicNumStr := "-123456789"
	expectedResult := "-123.456789"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_01(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "123456789"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_02(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "12345678.9"
	shiftPlacesLeft := uint(2)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_03(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "123456789000"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_04(t *testing.T) {
	basicNumStr := "123456789"
	expectedResult := "123456789000000"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_05(t *testing.T) {
	basicNumStr := "123"
	expectedResult := "12300000"
	shiftPlacesLeft := uint(5)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_06(t *testing.T) {
	basicNumStr := "0"
	expectedResult := "0"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_07(t *testing.T) {
	basicNumStr := "123456.789"
	expectedResult := "123456.789"
	shiftPlacesLeft := uint(0)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_08(t *testing.T) {
	basicNumStr := "-123456.789"
	expectedResult := "-123456.789"
	shiftPlacesLeft := uint(0)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_09(t *testing.T) {
	basicNumStr := "-123456.789"
	expectedResult := "-123456789"
	shiftPlacesLeft := uint(3)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}

func TestBigIntNum_ShiftPrecisionRight_10(t *testing.T) {
	basicNumStr := "-123456789"
	expectedResult := "-123456789000000"
	shiftPlacesLeft := uint(6)

	bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
			"basicNumStr='%v' Error='%v' ",
			basicNumStr, err.Error())
	}

	bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

	actualResult := bIntNum.GetNumStr()

	if expectedResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
			expectedResult, actualResult)
	}

}
