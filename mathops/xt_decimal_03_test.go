package mathops

import (
	"testing"
	"math/big"
)


func TestDecimal_NewBigInt_01(t *testing.T) {

	bigI := big.NewInt(int64(123456123456))
	precision := uint(6)
	expected := "123456.123456"
	d, err  := Decimal{}.NewBigInt(bigI, precision)

	if err != nil {
		t.Errorf("Error Returned by Decimal{}.NewBigInt(bigI, precision). " +
			"bigI= '%v' precision= '%v' Error= %v",
			bigI.String(), precision, err)
	}


	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewI64_01(t *testing.T) {

	i64 := int64(123456)
	precision := uint(3)
	expected := "123.456"
	d, err := Decimal{}.NewInt64(i64, precision)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewInt64(i64, precision) " +
			"i64='%v' precision='%v'  Error = '%v' ",
			i64, precision, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewInt_01(t *testing.T) {
	iNum := int(123456)
	precision := uint(3)
	expected := "123.456"
	d, err := Decimal{}.NewInt(iNum, precision)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewInt(iNum, precision). " +
			"iNum='%v' precision='%v' Error='%v'",
			iNum, precision, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}
}

func TestDecimal_NewOne_01(t *testing.T) {
	expectedNumStr := "1.000"
	expectedPrecision := uint(3)

	bINum := Decimal{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewOne_02(t *testing.T) {
	expectedNumStr := "1"
	expectedPrecision := uint(0)

	bINum := Decimal{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewOne_03(t *testing.T) {
	expectedNumStr := "1.00000"
	expectedPrecision := uint(5)

	bINum := Decimal{}.NewOne(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewTwo_01(t *testing.T) {
	expectedNumStr := "2.000"
	expectedPrecision := uint(3)

	bINum := Decimal{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewTwo_02(t *testing.T) {
	expectedNumStr := "2"
	expectedPrecision := uint(0)

	bINum := Decimal{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewTwo_03(t *testing.T) {
	expectedNumStr := "2.00000"
	expectedPrecision := uint(5)

	bINum := Decimal{}.NewTwo(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewThree_01(t *testing.T) {
	expectedNumStr := "3.000"
	expectedPrecision := uint(3)

	bINum := Decimal{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewThree_02(t *testing.T) {
	expectedNumStr := "3"
	expectedPrecision := uint(0)

	bINum := Decimal{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewThree_03(t *testing.T) {
	expectedNumStr := "3.00000"
	expectedPrecision := uint(5)

	bINum := Decimal{}.NewThree(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewFive_01(t *testing.T) {
	expectedNumStr := "5.000"
	expectedPrecision := uint(3)

	bINum := Decimal{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewFive_02(t *testing.T) {
	expectedNumStr := "5"
	expectedPrecision := uint(0)

	bINum := Decimal{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewFive_03(t *testing.T) {
	expectedNumStr := "5.00000"
	expectedPrecision := uint(5)

	bINum := Decimal{}.NewFive(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}
}

func TestDecimal_NewTen_01(t *testing.T) {
	expectedNumStr := "10.000"
	expectedPrecision := uint(3)

	bINum := Decimal{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewTen_02(t *testing.T) {
	expectedNumStr := "10"
	expectedPrecision := uint(0)

	bINum := Decimal{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}

func TestDecimal_NewTen_03(t *testing.T) {
	expectedNumStr := "10.00000"
	expectedPrecision := uint(5)

	bINum := Decimal{}.NewTen(expectedPrecision)

	if expectedNumStr != bINum.GetNumStr() {
		t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
			expectedNumStr, bINum.GetNumStr())
	}

	if expectedPrecision != bINum.GetPrecisionUint() {
		t.Errorf("Error: Expected Precision='%v'. Instead, Precision='%v'",
			expectedPrecision, bINum.GetPrecisionUint())
	}

}


func TestDecimal_NewNumStr_01(t *testing.T) {
	inStr := "123.456"
	expected := "123.456"
	d, err := Decimal{}.NewNumStr(inStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(inStr) " +
			"inStr='%v' Error = '%v' ", inStr, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStr_02(t *testing.T) {

	inStr := "123456"
	expected := "123456"
	d, err := Decimal{}.NewNumStr(inStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(inStr) " +
			"inStr='%v' Error = '%v' ", inStr, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStr_03(t *testing.T) {

	inStr := "-123456"
	expected := "-123456"
	d, err := Decimal{}.NewNumStr(inStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(inStr) " +
			"inStr='%v' Error = '%v' ", inStr, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStr_04(t *testing.T) {

	inStr := "-123.456"
	expected := "-123.456"
	d, err := Decimal{}.NewNumStr(inStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(inStr) " +
			"inStr='%v' Error = '%v' ", inStr, err.Error())
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStrWithNumSeps_01(t *testing.T) {

	nStr := "123,456"

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

	dec, err := Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := dec.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := dec.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestDecimal_NewNumStrWithNumSeps_02(t *testing.T) {

	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}
	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

	dec, err := Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := dec.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := dec.GetNumericSeparatorsDto()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestDecimal_NewNumStrWithNumSeps_03(t *testing.T) {

	nStr := "123.456"

	expectedNumSeps := NumericSeparatorDto{}

	dec, err := Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps(nStr, expectedNumSeps). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := dec.GetNumStr()

	if nStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			nStr, actualNumStr)
	}

	actualNumSeps := dec.GetNumericSeparatorsDto()

	expectedNumSeps.SetDefaultsIfEmpty()

	if !expectedNumSeps.Equal(actualNumSeps) {
		t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'.",
			expectedNumSeps.String(), actualNumSeps.String())
	}
}

func TestDecimal_NewNumStrDto_01(t *testing.T) {
	nStr1 := "1.35"
	ePrecision := uint(2)
	eSignVal := 1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ", nDto.GetNumStr(), err.Error())
	}

	if nStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

func TestDecimal_NewNumStrDto_02(t *testing.T) {
	nStr1 := "-1.35"
	ePrecision := uint(2)
	eSignVal := -1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ", nDto.GetNumStr(), err.Error())
	}

	if nStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

func TestDecimal_NewNumStrDto_03(t *testing.T) {
	nStr1 := "0.00"
	ePrecision := uint(2)
	eSignVal := 1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ", nDto.GetNumStr(), err.Error())
	}

	if nStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

func TestDecimal_NewNumStrDto_04(t *testing.T) {

	nStr1 := "-0.00"
	eNumStr1 := "0.00"
	ePrecision := uint(2)
	eSignVal := 1
	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStrDto(nDto)

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

func TestDecimal_NewNumStrDto_05(t *testing.T) {

	nStr1 := "92"
	eNumStr1 := "92"
	ePrecision := uint(0)
	eSignVal := 1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ", nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStrDto(nDto)

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

func TestDecimal_NewNumStrPrecision_01(t *testing.T) {
	nStr := "123456"
	precision := uint(3)
	expected := "123456.000"

	d, err := Decimal{}.NewNumStrPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(nStr, precision, false). inStr= '%v' precision= '%v' Error= %v", nStr, precision, err)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStrPrecision_02(t *testing.T) {

	numStr := "0"
	expected := "0.00"

	precision := uint(2)

	d1, err := Decimal{}.NewNumStrPrecision(numStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(numStrDto, precision, false). inStr= '%v' precision= '%v' Error= %v", numStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

}
func TestDecimal_NewNumStrPrecision_03(t *testing.T) {

	numStr := "125"
	expected := "125"

	precision := uint(0)

	d1, err := Decimal{}.NewNumStrPrecision(numStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(numStrDto, precision, false). inStr= '%v' precision= '%v' Error= %v", numStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

}

func TestDecimal_NewNumStrPrecision_04(t *testing.T) {
	inStr := "-123.456"
	expected := "-123.4560"
	precision := uint(4)
	d, err := Decimal{}.NewNumStrPrecision(inStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(inStr, precision, false). inStr= '%v' precision= '%v' Error= %v", inStr, precision, err)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}


	if int(precision) != d.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v'", precision, d.GetPrecision())
	}

}

func TestDecimal_NewNumStrPrecision_06(t *testing.T) {
	inStr := "123.456"
	expected := "123.4560"
	precision := uint(4)

	d, err := Decimal{}.NewNumStrPrecision(inStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(inStr, precision, false). inStr= '%v' precision= '%v' Error= %v", inStr, precision, err)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NewNumStrPrecision_07(t *testing.T) {
	nStr := "123456"
	precision := uint(3)
	expected := "123456.000"

	d, err := Decimal{}.NewNumStrPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Error Returned from Decimal.NewNumStrPrecision(nStr, precision, false). inStr= '%v' precision= '%v' Error= %v", nStr, precision, err)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_NthRoot_01(t *testing.T) {
	numStr1 := "125"
	nthRootStr := "5"
	maxPrecision := uint(14)
	expected := "2.62652780440377"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",
			numStr1, err.Error())
	}

	decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error = '%v' ",
			nthRootStr, err.Error())
	}

	d2, err := d1.NthRoot(decNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.NthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_NthRoot_02(t *testing.T) {
	numStr1 := "5604423"
	nthRootStr := "6"
	maxPrecision := uint(13)
	expected := "13.3276982415963"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error = '%v' ",
			nthRootStr, err.Error())
	}

	d2, err := d1.NthRoot(decNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.NthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_NthRoot_03(t *testing.T) {
	numStr1 := "5604423.924"
	nthRootStr := "6"
	maxPrecision := uint(13)
	expected := "13.3276986078187"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error = '%v' ",
			nthRootStr, err.Error())
	}

	d2, err := d1.NthRoot(decNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.NthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_NthRoot_04(t *testing.T) {
	numStr1 := "-27"
	nthRootStr := "3"
	maxPrecision := uint(2)
	expected := "-3.00"
	eSignVal := -1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error = '%v' ",
			nthRootStr, err.Error())
	}

	d2, err := d1.NthRoot(decNthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.NthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_NthRoot_05(t *testing.T) {
	numStr1 := "-27"
	nthRootStr := "4"
	maxPrecision := uint(2)

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	decNthRoot, err := Decimal{}.NewNumStr(nthRootStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nthRootStr) " +
			"nthRootStr='%v' Error = '%v' ",
			nthRootStr, err.Error())
	}

	_, err = d1.NthRoot(decNthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from d1.NthRoot(nthRoot, maxPrecision) for negative number with even nthRoot. No Error triggered")
	}


}

func TestDecimal_NumStrToDecimal_01(t *testing.T) {
	d1 := Decimal{}.New()

	numStr1 := "123456"

	d2, err := d1.NumStrToDecimal(numStr1)

	if err != nil {
		t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
	}

	expected := numStr1

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_NumStrToDecimal_02(t *testing.T) {
	d1 := Decimal{}.New()

	numStr1 := "12345.6"

	d2, err := d1.NumStrToDecimal(numStr1)

	if err != nil {
		t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
	}

	expected := numStr1

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_NumStrToDecimal_03(t *testing.T) {
	d1 := Decimal{}.New()

	numStr1 := "-123456"

	d2, err := d1.NumStrToDecimal(numStr1)

	if err != nil {
		t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
	}

	expected := numStr1

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_NumStrToDecimal_04(t *testing.T) {
	d1 := Decimal{}.New()

	numStr1 := "-12345.6"

	d2, err := d1.NumStrToDecimal(numStr1)

	if err != nil {
		t.Errorf("Received error from d1.NumStrToDecimal(numStr1). numStr1:= %v", numStr1)
	}

	expected := numStr1

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_01(t *testing.T) {
	inStr := "123.456789"
	precision := uint(3)
	expected := "123.457"
	eSignVal := 1

	d := Decimal{}
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, true)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}


	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_02(t *testing.T) {

	inStr := "123456789"
	expected := "123456789.000"
	eSignVal := 1
	precision := uint(3)

	d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision, true)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}


	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_03(t *testing.T) {

	inStr := "123456789"
	expected := "123456789.000000000"
	eSignVal := 1

	precision := uint(9)
	d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_04(t *testing.T) {

	inStr := "123456789"
	expected := "123456789.0000000000"
	eSignVal := 1

	d := Decimal{}
	precision := uint(10)
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_05(t *testing.T) {

	inStr := "-123456789"
	expected := "-123456789.0000000000"
	eSignVal := -1

	d := Decimal{}
	precision := uint(10)
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_06(t *testing.T) {

	inStr := "-123456.789"
	expected := "-123456.789000"
	eSignVal := -1
	d := Decimal{}
	precision := uint(6)
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}


	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}

func TestDecimal_NumStrPrecisionToDecimal_07(t *testing.T) {

	inStr := "5"
	expected := "5.0"
	precision := uint(1)
	eSignVal := 1

	d := Decimal{}
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}
}

func TestDecimal_NumStrPrecisionToDecimal_08(t *testing.T) {

	inStr := "0.5"
	expected := "0.5"
	precision := uint(1)
	eSignVal := 1

	d := Decimal{}
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}


	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}
}

func TestDecimal_NumStrPrecisionToDecimal_09(t *testing.T) {

	inStr := "123456"
	expected := "123456.000"
	precision := uint(3)
	eSignVal := 1

	d := Decimal{}
	d1, err := d.NumStrPrecisionToDecimal(inStr, precision, false)

	if err != nil {
		t.Errorf("Error returned from d.NumStrPrecisionToDecimal(inStr, precision). inStr='%v' precision= %v Error= %v  \n", inStr, precision, err)
	}

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d1.GetNumStr())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}
