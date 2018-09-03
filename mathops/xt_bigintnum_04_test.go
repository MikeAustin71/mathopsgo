package mathops

import (
	"testing"
	"math/big"
	)


func TestBigIntNum_Multiply_01(t *testing.T) {

	str1 := "575.63"
	str2 := "2014.123"
	expected := "1159389.62249"

	bINum1 := BigIntNum{}.New()

	err := bINum1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	bINum2 := BigIntNum{}.New()

	err = bINum2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	bINum3 := bINum1.Multiply(bINum2)

	if bINum3.GetNumStr() != expected {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, bINum3.GetNumStr())
	}

}

func TestBigIntNum_Multiply_02(t *testing.T) {

	str1 := "-575.63"
	str2 := "2014.123"
	expected := "-1159389.62249"

	bINum1 := BigIntNum{}.New()

	err := bINum1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	bINum2 := BigIntNum{}.New()

	err = bINum2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	bINum3 := bINum1.Multiply(bINum2)

	if expected != bINum3.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, bINum3.GetNumStr())
	}

}

func TestBigIntNum_Multiply_03(t *testing.T) {

	str1 := "-575.63"
	str2 := "-2014.123"
	expected := "1159389.62249"

	bINum1 := BigIntNum{}.New()

	err := bINum1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	bINum2 := BigIntNum{}.New()

	err = bINum2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	bINum3 := bINum1.Multiply(bINum2)

	if expected != bINum3.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, bINum3.GetNumStr())
	}

}

func TestBigIntNum_Multiply_04(t *testing.T) {

	str1 := "0"
	str2 := "-2014.123"
	expected := "0"

	bINum1 := BigIntNum{}.New()

	err := bINum1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	bINum2 := BigIntNum{}.New()

	err = bINum2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on bINum1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	bINum3 := bINum1.Multiply(bINum2)

	if expected != bINum3.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, bINum3.GetNumStr())
	}

}

func TestBigIntNum_Multiply_05(t *testing.T) {
	numStr := "3"
	mul, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrDto) " +
			"numStrDto='%v'  Error = '%v' ",numStr, err.Error())
	}

	bigIntNum, err := BigIntNum{}.NewNumStr("1")

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(\"1\") " +
			"Error = '%v' ", err.Error())
	}

	for i := 0; i < 4; i++ {

		bigIntNum = bigIntNum.Multiply(mul)

		if err != nil {
			t.Errorf("Error returned by bigIntNum.Multiply(mul). " +
				"i='%v' bigIntNum='%v', mul='%v' Error='%v'",
				i, bigIntNum.GetNumStr(), mul.GetNumStr(), err.Error())
		}
	}

	expected := "81"

	if expected != bigIntNum.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, bigIntNum.GetNumStr())
	}

}

func TestBigIntNum_MultiplyByTenToPower_01(t *testing.T) {
	numStr := "85.621"
	exponent := uint(3)
	expectedStr := "85621"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_MultiplyByTenToPower_02(t *testing.T) {
	numStr := "85.621"
	exponent := uint(5)
	expectedStr := "8562100"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_MultiplyByTenToPower_03(t *testing.T) {
	numStr := "-85.621"
	exponent := uint(3)
	expectedStr := "-85621"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_MultiplyByTenToPower_04(t *testing.T) {
	numStr := "-85.621"
	exponent := uint(5)
	expectedStr := "-8562100"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_MultiplyByTenToPower_06(t *testing.T) {
	numStr := "-85.621"
	exponent := uint(0)
	expectedStr := "-85.621"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPower(exponent)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestBigIntNum_NewBigFloat_01(t *testing.T) {

	bfloat := big.NewFloat(32.123)

	expectedNumStr := "32.123"

	bINum, err := BigIntNum{}.NewBigFloat(bfloat, 3)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bfloat, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewBigFloat_02(t *testing.T) {

	bFloat := big.NewFloat(float64(32.129))

	expectedNumStr := "32.13"

	bINum, err := BigIntNum{}.NewBigFloat(bFloat, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bFloat, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewBigFloat_03(t *testing.T) {

	bFloat := big.NewFloat(-32.129)

	expectedNumStr := "-32.13"

	bINum, err := BigIntNum{}.NewBigFloat(bFloat, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bFloat, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat32_01(t *testing.T) {

	numf32 := float32(32.123)

	expectedNumStr := "32.123"

	maxPrecision := uint(4)

	bINum, err := BigIntNum{}.NewFloat32(numf32, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat32_02(t *testing.T) {

	numf32 := float32(32.129)

	expectedNumStr := "32.13"

	bINum, err := BigIntNum{}.NewFloat32(numf32, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat32_03(t *testing.T) {

	numf32 := float32(-32.129)

	expectedNumStr := "-32.13"

	bINum, err := BigIntNum{}.NewFloat32(numf32, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat64_01(t *testing.T) {

	numf32 := float64(32.123)

	expectedNumStr := "32.123"

	maxPrecision := uint(3)

	bINum, err := BigIntNum{}.NewFloat64(numf32, maxPrecision)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat64_02(t *testing.T) {

	numf32 := float64(32.129)

	expectedNumStr := "32.13"

	bINum, err := BigIntNum{}.NewFloat64(numf32, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewFloat64_03(t *testing.T) {

	numf32 := float64(-32.129)

	expectedNumStr := "-32.13"

	bINum, err := BigIntNum{}.NewFloat64(numf32, 2)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_01(t *testing.T) {

	numInt := 1234

	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewIntExponent(numInt, 3)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_02(t *testing.T) {

	numInt := 123456

	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewIntExponent(numInt, -2)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_03(t *testing.T) {

	numInt := 123456

	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewIntExponent(numInt, 0)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32Exponent_01(t *testing.T) {

	numInt := int32(1234)

	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, 3)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32Exponent_02(t *testing.T) {

	numInt := int32(123456)

	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, -2)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32Exponent_03(t *testing.T) {

	numInt := int32(123456)

	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, 0)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_01(t *testing.T) {

	numInt := int64(1234)

	expectedNumStr := "1234.000"

	bINum := BigIntNum{}.NewInt64Exponent(numInt, 3)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_02(t *testing.T) {

	numInt := int64(123456)

	expectedNumStr := "1234.56"

	bINum := BigIntNum{}.NewInt64Exponent(numInt, -2)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_03(t *testing.T) {

	numInt := int64(123456)

	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewInt64Exponent(numInt, 0)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_04(t *testing.T) {

	numInt := int64(-123456)

	expectedNumStr := "-123456"

	exponent := 0

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_05(t *testing.T) {

	numInt := int64(-123456)

	expectedNumStr := "-123.456"

	exponent := -3

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_06(t *testing.T) {

	numInt := int64(-123456)

	expectedNumStr := "-123456.000"

	exponent := 3

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). " +
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). " +
			"Error='%v'", err.Error())
	}

	actualNumStr := bigINum.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}
}

func TestBigIntNum_NewBigIntExponent_01(t *testing.T) {

	n:= 123456
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

	n:= 123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "123.456"
	bExpected := big.NewInt(int64(123456))
	expectedSignVal := 1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

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

	n:= -123456
	exponent := -3
	expectedPrecision := uint(3)
	expectedNumStr := "-123.456"
	bExpected := big.NewInt(int64(-123456))
	expectedSignVal := -1

	bOriginal := big.NewInt(int64(n))

	bINum := BigIntNum{}.NewBigIntExponent(bOriginal, exponent)

	actualNumStr := bINum.GetNumStr()

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

func TestBigIntNum_NumStrDto_01(t *testing.T) {

	nStr:="123.456"
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
		t.Error("Error returned by big.NewInt(0).SetString(nbStr, 10).")
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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

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

	actualNumStr := bINum1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
			expectedNumStr, actualNumStr )
	}

}
