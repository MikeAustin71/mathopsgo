package mathops

import (
	"math/big"
	"testing"
)

func TestNFactorial_CalcFactorialValueBigInt_01(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_02(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_03(t *testing.T) {

	n := big.NewInt(-5)

	lowerLimit := big.NewInt(0)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueBigInt_04(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(-2)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueBigInt_05(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_06(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_07(t *testing.T) {

	n := big.NewInt(23)

	lowerLimit := big.NewInt(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueBigInt_08(t *testing.T) {

	n := big.NewInt(5)

	lowerLimit := big.NewInt(9)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=5 lower limit =9")

	}

}

func TestNFactorial_CalcFactorialValueBigInt_09(t *testing.T) {

	n := big.NewInt(0)

	lowerLimit := big.NewInt(1)

	_, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=0 lower limit =1")

	}

}


func TestNFactorial_CalcFactorialValueBigInt_10(t *testing.T) {

	n := big.NewInt(23)

	lowerLimit := big.NewInt(20)

	expectedResultStr := "10626"

	result, err := NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueBigInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_01(t *testing.T) {

	n := 5

	lowerLimit := 1

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_02(t *testing.T) {

	n := 5

	lowerLimit := 0

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_03(t *testing.T) {

	n := -5

	lowerLimit := 0

	_, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueBigInt(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt_04(t *testing.T) {

	n := 5

	lowerLimit := -2

	_, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt_05(t *testing.T) {

	n := 5

	lowerLimit := 2

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_06(t *testing.T) {

	n := 5

	lowerLimit := 4

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt_07(t *testing.T) {

	n := 23

	lowerLimit := 0

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}


func TestNFactorial_CalcFactorialValueInt32_01(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_02(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_03(t *testing.T) {

	n := int32(-5)

	lowerLimit := int32(0)

	_, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt32(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt32_04(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(-2)

	_, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt32(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt32_05(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_06(t *testing.T) {

	n := int32(5)

	lowerLimit := int32(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt32_07(t *testing.T) {

	n := int32(23)

	lowerLimit := int32(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_01(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_02(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_03(t *testing.T) {

	n := int64(-5)

	lowerLimit := int64(0)

	_, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt64(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcFactorialValueInt64_04(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(-2)

	_, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueInt64(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcFactorialValueInt64_05(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_06(t *testing.T) {

	n := int64(5)

	lowerLimit := int64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueInt64_07(t *testing.T) {

	n := int64(23)

	lowerLimit := int64(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueInt64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_01(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_02(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_03(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(94)


	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Error("Expected error return from CalcNFactorialValue(n, lowerLimit) " +
			"n=-5")

	}

}

func TestNFactorial_CalcNFactorialValue_04(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(0)

	nFac.LowerLimit = uint64(2)

	_, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err == nil {
		t.Error("Expected error return from CalcNFactorialValue(n, lowerLimit) " +
			"lower limit =-2")

	}

}

func TestNFactorial_CalcNFactorialValue_05(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(nFac). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_06(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(5)

	nFac.LowerLimit = uint64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcNFactorialValue(nFac). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcNFactorialValue_07(t *testing.T) {

	nFac := FactorialDto{}

	nFac.UpperLimit = uint64(23)

	nFac.LowerLimit = uint64(1)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcNFactorialValue(nFac)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_01(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_02(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_03(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(9)

	_, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint_04(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(32)

	_, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint_05(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_06(t *testing.T) {

	n := uint(5)

	lowerLimit := uint(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint_07(t *testing.T) {

	n := uint(23)

	lowerLimit := uint(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_01(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_02(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_03(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(9)

	_, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint32(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint32_04(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(32)

	_, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint32(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint32_05(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_06(t *testing.T) {

	n := uint32(5)

	lowerLimit := uint32(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint32(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint32_07(t *testing.T) {

	n := uint32(23)

	lowerLimit := uint32(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint32(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_01(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(1)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_02(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(0)

	expectedResultStr := "120"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_03(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(9)

	_, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint64(n, lowerLimit) " +
			"n=5 lower limit=9")

	}

}

func TestNFactorial_CalcFactorialValueUint64_04(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(32)

	_, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err == nil {
		t.Error("Expected error return from CalcFactorialValueUint64(n, lowerLimit) " +
			"n=5 lower limit =32")

	}

}

func TestNFactorial_CalcFactorialValueUint64_05(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(2)

	expectedResultStr := "60"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_06(t *testing.T) {

	n := uint64(5)

	lowerLimit := uint64(4)

	expectedResultStr := "5"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueUint64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}

func TestNFactorial_CalcFactorialValueUint64_07(t *testing.T) {

	n := uint64(23)

	lowerLimit := uint64(0)

	expectedResultStr := "25852016738884976640000"

	result, err := NFactorial{}.CalcFactorialValueUint64(n, lowerLimit)

	if err != nil {
		t.Errorf("Error returned by NFactorial{}.CalcFactorialValueInt64(n, lowerLimit). " +
			"Error='%v' ", err.Error())
	}

	actualResultStr := result.GetNumStr()

	if expectedResultStr != actualResultStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedResultStr, actualResultStr)
	}

}
