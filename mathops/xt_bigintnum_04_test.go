package mathops

import (
	"math/big"
	"testing"
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
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStrDto) "+
			"numStrDto='%v'  Error = '%v' ", numStr, err.Error())
	}

	bigIntNum, err := BigIntNum{}.NewNumStr("1")

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(\"1\") "+
			"Error = '%v' ", err.Error())
	}

	for i := 0; i < 4; i++ {

		bigIntNum = bigIntNum.Multiply(mul)

		if err != nil {
			t.Errorf("Error returned by bigIntNum.Multiply(mul). "+
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

func TestBigIntNum_MultiplyByTenToPowerAdd_01(t *testing.T) {

	numStr := "85621"
	exponent := uint(1)
	addendStr := "9"
	expectedStr := "856219"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bIAddend, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPowerAdd(exponent, bIAddend)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}
}

func TestBigIntNum_MultiplyByTenToPowerAdd_02(t *testing.T) {

	numStr := "8562.1"
	exponent := uint(2)
	addendStr := "9"
	expectedStr := "856219"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bIAddend, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPowerAdd(exponent, bIAddend)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}
}

func TestBigIntNum_MultiplyByTenToPowerAdd_03(t *testing.T) {

	numStr := "8562.123"
	exponent := uint(4)
	addendStr := "9"
	expectedStr := "85621239"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bIAddend, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPowerAdd(exponent, bIAddend)

	actualStr := bINum.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'.  Instead, result='%v'",
			expectedStr, actualStr)
	}
}

func TestBigIntNum_MultiplyByTenToPowerAdd_04(t *testing.T) {

	numStr := "8562.123"
	exponent := uint(4)
	addendStr := "-10"
	expectedStr := "85621220"

	bINum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(numStr). "+
			"Error='%v'", err.Error())
	}

	bIAddend, err := BigIntNum{}.NewNumStr(addendStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(addendStr). "+
			"Error='%v'", err.Error())
	}

	bINum.MultiplyByTenToPowerAdd(exponent, bIAddend)

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
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bfloat, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bFloat, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewBigFloat(bFloat, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat32(numf32, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) "+
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
		t.Errorf("Error returned by BigIntNum{}.NewFloat64(numf32, 4) "+
			"Error='%v' ", err.Error())
	}

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_01(t *testing.T) {

	numInt := int(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewInt(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_02(t *testing.T) {

	numInt := int(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewInt(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_03(t *testing.T) {

	numInt := int(-1234)
	precision := uint(3)
	expectedNumStr := "-1.234"

	bINum := BigIntNum{}.NewInt(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_04(t *testing.T) {

	numInt := int(-1234)
	precision := uint(0)
	expectedNumStr := "-1234"

	bINum := BigIntNum{}.NewInt(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_05(t *testing.T) {

	numInt := int(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewInt(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt_06(t *testing.T) {

	numInt := int(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt(numInt, precision)

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
	exponent := 0
	expectedNumStr := "123456"

	bINum := BigIntNum{}.NewIntExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_04(t *testing.T) {

	numInt := 0
	exponent := 0
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewIntExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_05(t *testing.T) {

	numInt := 0
	exponent := 3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewIntExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewIntExponent_06(t *testing.T) {

	numInt := 0
	exponent := -3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewIntExponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_01(t *testing.T) {

	numInt := int32(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_02(t *testing.T) {

	numInt := int32(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_03(t *testing.T) {

	numInt := int32(-1234)
	precision := uint(3)
	expectedNumStr := "-1.234"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_04(t *testing.T) {

	numInt := int32(-1234)
	precision := uint(0)
	expectedNumStr := "-1234"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_05(t *testing.T) {

	numInt := int32(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32_06(t *testing.T) {

	numInt := int32(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt32(numInt, precision)

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

func TestBigIntNum_NewInt32Exponent_04(t *testing.T) {

	numInt := int32(0)
	exponent := 0
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32Exponent_05(t *testing.T) {

	numInt := int32(0)
	exponent := 3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt32Exponent_06(t *testing.T) {

	numInt := int32(0)
	exponent := -3
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt32Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_01(t *testing.T) {

	numInt := int64(1234)
	precision := uint(3)
	expectedNumStr := "1.234"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_02(t *testing.T) {

	numInt := int64(1234)
	precision := uint(0)
	expectedNumStr := "1234"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_03(t *testing.T) {

	numInt := int64(-1234)
	precision := uint(3)
	expectedNumStr := "-1.234"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_04(t *testing.T) {

	numInt := int64(-1234)
	precision := uint(0)
	expectedNumStr := "-1234"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_05(t *testing.T) {

	numInt := int64(0)
	precision := uint(0)
	expectedNumStr := "0"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64_06(t *testing.T) {

	numInt := int64(0)
	precision := uint(3)
	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt64(numInt, precision)

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

func TestBigIntNum_NewInt64Exponent_07(t *testing.T) {

	numInt := int64(0)

	expectedNumStr := "0"

	exponent := 0

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_08(t *testing.T) {

	numInt := int64(0)

	expectedNumStr := "0.000"

	exponent := 3

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewInt64Exponent_09(t *testing.T) {

	numInt := int64(0)

	exponent := -3

	expectedNumStr := "0.000"

	bINum := BigIntNum{}.NewInt64Exponent(numInt, exponent)

	actualNumStr := bINum.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
			expectedNumStr, actualNumStr)
	}

}

func TestBigIntNum_NewINumMgr_01(t *testing.T) {

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

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) "+
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

func TestBigIntNum_NewINumMgr_02(t *testing.T) {

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

	ia, err := IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by IntAry{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&ia)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&ia) "+
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

func TestBigIntNum_NewINumMgr_03(t *testing.T) {

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

	nDto0, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(&nDto0)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&nDto0) "+
			"Error='%v' ", err.Error())
	}

	nDto1, err := NumStrDto{}.NewBigInt(bINum.bigInt, bINum.precision)

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

	if nStr != nDto1.GetNumStr() {
		t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'. ",
			nStr, nDto1.GetNumStr())
	}

}

func TestBigIntNum_NewINumMgr_04(t *testing.T) {

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

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(dec.GetThisPointer())

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) "+
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

func TestBigIntNum_NewINumMgr_05(t *testing.T) {

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

	dec := Decimal{}.NewPtr()
	err := dec.SetNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr). Error='%v' ",
			err.Error())
	}

	bINum, err := BigIntNum{}.NewINumMgr(dec)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewINumMgr(&dec) "+
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
