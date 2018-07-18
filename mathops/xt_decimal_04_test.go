package mathops

import (
	"testing"
	"math/big"
)


func TestDecimal_PowInt_01(t *testing.T) {

	numStr := "2.125"
	exp := 5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStrDto) " +
			"numStrDto='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.PowInt(exp, 15)

	if err != nil {
		t.Errorf("Received error from d1.PowInt(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "43.330596923828125"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_PowInt_02(t *testing.T) {

	numStr := "2.125"
	exp := -5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStrDto) " +
			"numStrDto='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.PowInt(exp, 32)

	if err != nil {
		t.Errorf("Received error from d1.PowInt(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "0.02307838042845159759046157465153"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_PowInt_03(t *testing.T) {

	numStr := "2.125"
	exp := -5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStrDto) " +
			"numStrDto='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.PowInt(exp, 250)

	if err != nil {
		t.Errorf("Received error from d1.PowInt(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "0.0230783804284515975904615746515318091892352539727592285702010836302529057503678187310412245740240038257373805953698154109885713843013768287933221444131345621425256205378428954465132756326869536861810731644102187755527493261645362878092652992519669234"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_SetPrecisionRound_01(t *testing.T) {

	str1 := "2.0105500"
	precision := uint(4)
	expected := "2.0106"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' ", str1, err.Error())
	}

	if str1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", str1, d1.GetNumStr())
	}

	d1.SetPrecisionRound(precision)

	if expected != d1.GetNumStr() {
		t.Errorf("Expected 2nd NumStr= '%v'. Instead, got %v.", expected, d1.GetNumStr())
	}

}

func TestDecimal_SetPrecisionRound_02(t *testing.T) {

	str1 := "-2.0105500"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' ", str1, err.Error())
	}

	if str1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", str1, d1.GetNumStr())
	}

	d1.SetPrecisionRound(4)

	expected := "-2.0106"

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d1.GetNumStr())
	}

}

func TestDecimal_SetPrecisionTrunc_01(t *testing.T) {
	str1 := "2.0105500"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' ", str1, err.Error())
	}

	if str1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", str1, d1.GetNumStr())
	}

	d1.SetPrecisionTrunc(4)

	expected := "2.0105"

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d1.GetNumStr())
	}

}

func TestDecimal_SetPrecisionTrunc_02(t *testing.T) {
	str1 := "-2.0105500"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' ", str1, err.Error())
	}

	if str1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", str1, d1.GetNumStr())
	}

	d1.SetPrecisionTrunc(4)

	expected := "-2.0105"

	if expected != d1.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d1.GetNumStr())
	}

}

func TestDecimal_SetNumStr_01(t *testing.T) {

	nStr1 := "1.35"
	ePrecision := uint(2)
	eSignVal := 1
	d1 := Decimal{}.New()

	d1.SetNumStr(nStr1)

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

func TestDecimal_SetNumStr_02(t *testing.T) {

	nStr1 := "-1.35"
	ePrecision := uint(2)
	eSignVal := -1
	d1 := Decimal{}.New()

	d1.SetNumStr(nStr1)


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

func TestDecimal_SetNumStr_03(t *testing.T) {

	nStr1 := "0.00"
	ePrecision := uint(2)
	eSignVal := 1
	d1 := Decimal{}.New()

	d1.SetNumStr(nStr1)

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

func TestDecimal_SetNumStr_04(t *testing.T) {

	nStr1 := "-0.00"
	eNumStr1 := "0.00"
	ePrecision := uint(2)
	eSignVal := 1
	d1 := Decimal{}.New()

	d1.SetNumStr(nStr1)

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

func TestDecimal_SetNumStr_05(t *testing.T) {

	nStr1 := "92"
	eNumStr1 := "92"
	ePrecision := uint(0)
	eSignVal := 1
	d1 := Decimal{}.New()

	d1.SetNumStr(nStr1)

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

func TestDecimal_SetNumStrDto_01(t *testing.T) {

	nStr0 := "-999999.99999"

	nStr1 := "1.35"
	ePrecision := uint(2)
	eSignVal := 1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ",
			nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStr(nStr0)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). " +
			"nStr0='%v' Error='%v' ",
				nStr0, err.Error())
	}

	err = d1.SetNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by d1.SetNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ",
			nDto.GetNumStr(), err.Error())
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

func TestDecimal_SetNumStrDto_02(t *testing.T) {

	nStr0 := "-999999.99999"

	nStr1 := "-1555666.35"
	ePrecision := uint(2)
	eSignVal := -1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ",
			nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStr(nStr0)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). " +
			"nStr0='%v' Error='%v' ",
				nStr0, err.Error())
	}

	err = d1.SetNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by d1.SetNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ",
			nDto.GetNumStr(), err.Error())
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

func TestDecimal_SetNumStrDto_03(t *testing.T) {

	nStr0 := "-999999.99999"

	nStr1 := "1555777.123456"
	ePrecision := uint(6)
	eSignVal := 1

	nDto, err := NumStrDto{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). " +
			"nStr1='%v' Error='%v' ",
			nStr1, err.Error())
	}

	d1, err := Decimal{}.NewNumStr(nStr0)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). " +
			"nStr0='%v' Error='%v' ",
				nStr0, err.Error())
	}

	err = d1.SetNumStrDto(nDto)

	if err != nil {
		t.Errorf("Error returned by d1.SetNumStrDto(nDto). " +
			"nDto.GetNumStr()='%v' Error='%v' ",
			nDto.GetNumStr(), err.Error())
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

func TestDecimal_SetFloat_01(t *testing.T) {

	fVal := float32(92.25)
	eNumStr1 := "92.25"
	ePrecision := uint(2)
	eSignVal := 1
	d1 := Decimal{}.New()

	err := d1.SetFloat32(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat32(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	d1.SetPrecisionRound(ePrecision)

	if  eNumStr1 != d1.GetNumStr() {
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

func TestDecimal_SetFloat_02(t *testing.T) {

	fVal := float32(-92.25)
	eNumStr1 := "-92.25"
	ePrecision := uint(2)
	eSignVal := -1
	d1 := Decimal{}.New()

	err := d1.SetFloat32(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat32(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	d1.SetPrecisionRound(ePrecision)

	if  eNumStr1 != d1.GetNumStr() {
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

func TestDecimal_SetFloat64_01(t *testing.T) {

	fVal := float64(92.256)
	eNumStr1 := "92.256"
	ePrecision := uint(3)
	eSignVal := 1
	d1 := Decimal{}.New()

	err := d1.SetFloat64(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat64(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	d1.SetPrecisionRound(ePrecision)

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

func TestDecimal_SetFloat64_02(t *testing.T) {

	fVal := float64(-92.25)
	eNumStr1 := "-92.25"
	ePrecision := uint(2)
	eSignVal := -1
	d1 := Decimal{}.New()

	err := d1.SetFloat64(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat64(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	d1.SetPrecisionRound(ePrecision)

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
func TestDecimal_SetFloatBig_01(t *testing.T) {

	eNumStr1 := "92.256"
	fVal, isOk := big.NewFloat(0.0).SetString(eNumStr1)

	if !isOk {
		t.Errorf("bigFloat.SetString failed to convert eNumStr1. eNumStr1= '%v'", eNumStr1)
	}

	ePrecision := uint(3)
	eSignVal := 1
	d1 := Decimal{}.New()

	err := d1.SetFloatBig(fVal)

	if err != nil {
		t.Errorf("d1.SetFloatBig(fVal) returned an error. fVal= '%v' Error= %v", eNumStr1, err)
	}

	d1.SetPrecisionRound(ePrecision)

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

func TestDecimal_SetFloatBig_02(t *testing.T) {

	eNumStr1 := "-92.25"
	ePrecision := uint(2)
	eSignVal := -1
	fVal, isOk := big.NewFloat(0.0).SetString(eNumStr1)

	if !isOk {
		t.Errorf("bigFloat.SetString failed to convert eNumStr1. eNumStr1= '%v'", eNumStr1)
	}

	d1 := Decimal{}.New()

	err := d1.SetFloatBig(fVal)

	if err != nil {
		t.Errorf("d1.SetFloatBig(fVal) returned an error. fVal= '%v' Error= %v", eNumStr1, err)
	}

	d1.SetPrecisionRound(2)

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

func TestDecimal_SetIntFracStrings_01(t *testing.T) {

	intStr := "123"
	fracStr := "456"
	eNumStr := "123.456"
	eSignVal := 1
	ePrecision := uint(3)

	d1 := Decimal{}.New()

	err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

	if err != nil {
		t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
	}

	if eNumStr != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
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

func TestDecimal_SetIntFracStrings_02(t *testing.T) {

	intStr := "123"
	fracStr := "0456"
	eNumStr := "-123.0456"
	eSignVal := -1
	ePrecision := uint(4)

	d1 := Decimal{}.New()

	err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

	if err != nil {
		t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
	}


	if eNumStr != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
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
