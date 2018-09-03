package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_NewInt_01(t *testing.T) {

	intNum := 7
	precision := uint(0)

	expectedStr := "7"

	nDto, err := NumStrDto{}.NewInt(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt_02(t *testing.T) {

	intNum := 7
	precision := uint(1)

	expectedStr := "7.0"

	nDto, err := NumStrDto{}.NewInt(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt_03(t *testing.T) {

	intNum := 7
	precision := uint(3)

	expectedStr := "7.000"

	nDto, err := NumStrDto{}.NewInt(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt_04(t *testing.T) {

	intNum := -7
	precision := uint(3)

	expectedStr := "-7.000"

	nDto, err := NumStrDto{}.NewInt(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt_05(t *testing.T) {

	intNum := -7
	precision := uint(0)

	expectedStr := "-7"

	nDto, err := NumStrDto{}.NewInt(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewIntExponent_01(t *testing.T) {
	intNum := 7
	exponent := 3

	expectedStr := "7.000"

	nDto := NumStrDto{}.NewIntExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewIntExponent_02(t *testing.T) {
	intNum := 7123
	exponent := -3

	expectedStr := "7.123"

	nDto := NumStrDto{}.NewIntExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewIntExponent_03(t *testing.T) {
	intNum := -72
	exponent := 3

	expectedStr := "-72.000"

	nDto := NumStrDto{}.NewIntExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewIntExponent_04(t *testing.T) {
	intNum := -72123
	exponent := -3

	expectedStr := "-72.123"

	nDto := NumStrDto{}.NewIntExponent(intNum, exponent)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt32_01(t *testing.T) {

	intNum := int32(7)
	precision := uint(0)

	expectedStr := "7"

	nDto, err := NumStrDto{}.NewInt32(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt32(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt32_02(t *testing.T) {

	intNum := int32(7)
	precision := uint(1)

	expectedStr := "7.0"

	nDto, err := NumStrDto{}.NewInt32(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt32(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt32_03(t *testing.T) {

	intNum := int32(7)
	precision := uint(3)

	expectedStr := "7.000"

	nDto, err := NumStrDto{}.NewInt32(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt32(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt32_04(t *testing.T) {

	intNum := int32(-7)
	precision := uint(3)

	expectedStr := "-7.000"

	nDto, err := NumStrDto{}.NewInt32(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt32(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt32_05(t *testing.T) {

	intNum := int32(-7)
	precision := uint(0)

	expectedStr := "-7"

	nDto, err := NumStrDto{}.NewInt32(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt32(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt64_01(t *testing.T) {

	intNum := int64(7)
	precision := uint(0)

	expectedStr := "7"

	nDto, err := NumStrDto{}.NewInt64(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt64(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt64_02(t *testing.T) {

	intNum := int64(7)
	precision := uint(1)

	expectedStr := "7.0"

	nDto, err := NumStrDto{}.NewInt64(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt64(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt64_03(t *testing.T) {

	intNum := int64(7)
	precision := uint(3)

	expectedStr := "7.000"

	nDto, err := NumStrDto{}.NewInt64(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt64(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt64_04(t *testing.T) {

	intNum := int64(-7)
	precision := uint(3)

	expectedStr := "-7.000"

	nDto, err := NumStrDto{}.NewInt64(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt64(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewInt64_05(t *testing.T) {

	intNum := int64(-7)
	precision := uint(0)

	expectedStr := "-7"

	nDto, err := NumStrDto{}.NewInt64(intNum,precision)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewInt64(intNum,precision) " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v",
			false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
					precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_04(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). " +
			"nStr= '%v' Error= %v",nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
					precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_01(t *testing.T) {
	nStr := "123,456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps(" +
			"nStr, expectedNumSeps) Error='%v'", err.Error() )
	}

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_02(t *testing.T) {
	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.SetDefaultsIfEmpty()

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps(" +
			"nStr, expectedNumSeps) Error='%v'", err.Error() )
	}

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewNumStrWithNumSeps_03(t *testing.T) {
	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}

	nDto, err := NumStrDto{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStrWithNumSeps(" +
			"nStr, expectedNumSeps) Error='%v'", err.Error() )
	}

	expectedNumSeps.SetDefaultsIfEmpty()

	actualNumStr := nDto.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := nDto.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}

}

func TestNumStrDto_NewBigIntNum_01(t *testing.T) {
	nStr := "123.456"

	bINum, err := BigIntNum{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). " +
			"nStr='%v' Error='%v ",
				nStr, err.Error())
	}

	nDto, err := NumStrDto{}.NewBigIntNum(bINum)


	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewBigIntNum(bINum). " +
			"bINum='%v' Error='%v ",
			bINum.GetNumStr(), err.Error())
	}

	if nStr != nDto.GetNumStr() {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			nStr, nDto.GetNumStr())
	}

}

func TestNumStrDto_NewZero_01(t *testing.T) {

	expectedStr := "0"

	nDto := NumStrDto{}.NewZero(0)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_NewZero_02(t *testing.T) {

	expectedStr := "0.00"

	nDto := NumStrDto{}.NewZero(2)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_NewZero_03(t *testing.T) {

	expectedStr := "0.0000"

	nDto := NumStrDto{}.NewZero(4)

	actualNumStr := nDto.GetNumStr()

	if expectedStr != actualNumStr {
		t.Errorf("Expected nDto.GetNumStr()='%v'. Instead, nDto.GetNumStr()='%v'.",
			expectedStr, nDto.GetNumStr())
	}
}

func TestNumStrDto_ParseNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). " +
				"nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_04(t *testing.T) {
	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
					precision, nDto.GetPrecisionUint())
	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_06(t *testing.T) {
	nStr := "-123.4567#"
	nStrOut := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseNumStr_07(t *testing.T) {
	nStr := "-123.4#-567#"
	nStrOut := "-123.4567"
	iStr := "123"
	fracStr := "4567"
	signVal := -1
	precision := uint(4)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.GetNumStr()

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if nDto.IsFractionalValue() != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseSignedBigInt_01(t *testing.T) {

	signedAbsNumStr := "-123456789"
	absAllNumStr := "123456789"
	nStr := "-123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := -1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}

func TestNumStrDto_ParseSignedBigInt_02(t *testing.T) {

	signedAbsNumStr := "123456789"
	absAllNumStr := "123456789"
	nStr := "123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := 1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("bigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits() {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits())
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if precision != nDto.GetPrecisionUint() {
		t.Errorf("Expected precision= '%v'. Instead, got %v",
				precision, nDto.GetPrecisionUint())

	}

	err = nDto.IsValid("Test 'nDto' is INVALID! ")

	if err != nil {
		t.Errorf("Error returned by nDto.IsValid() Error='%v'", err.Error())
	}

}
