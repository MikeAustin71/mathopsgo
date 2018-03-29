package mathops

import (
	"testing"
	"math/big"
)

func TestDecimal_Subtract_01(t *testing.T) {

	nStr1 := "123456"
	expected := "123455.5"

	precision1 := uint(0)

	d1, err := Decimal{}.NewNumStrPrecision(nStr1, precision1, true)

	if err != nil {
		t.Errorf("Error Returned from Decimal d1.NewNumStrPrecision(nStr1, precision1, false). nStr1= '%v' precision1= '%v' Error= %v", nStr1, precision1, err)
	}

	nStr2 := "0.5"
	precision2 := uint(1)

	d2, err := Decimal{}.NewNumStrPrecision(nStr2, precision2, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal d2.NewNumStrPrecision(nStr2, precision2, false). nStr2= '%v' precision2= '%v' Error= %v", nStr2, precision2, err)
	}

	d3, err := d1.Subtract(d2)

	if err != nil {
		t.Errorf("Error received from Subtract. Error= %v", err)
	}

	if expected != d3.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d3.GetNumStr())
	}

}

func TestDecimal_Subtract_02(t *testing.T) {

	nStr1 := "123.222"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1) " +
			"nStr1='%v' Error = '%v' ", nStr1, err.Error())
	}

	nStr2 := "-1.2223"

	d2, err := Decimal{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr2) " +
			"nStr2='%v' Error = '%v' ", nStr2, err.Error())
	}

	d3, err := d1.Subtract(d2)

	if err != nil {
		t.Errorf("Error received from Subtract. Error= %v", err)
	}

	expected := "124.4443"

	if expected != d3.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d3.GetNumStr())
	}

}

func TestDecimal_Subtract_03(t *testing.T) {

	nStr1 := "-4"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1) " +
			"nStr1='%v' Error = '%v' ", nStr1, err.Error())
	}

	nStr2 := "-2"

	d2, err := Decimal{}.NewNumStr(nStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr2) " +
			"nStr2='%v' Error = '%v' ", nStr2, err.Error())
	}

	d3, err := d1.Subtract(d2)

	if err != nil {
		t.Errorf("Error received from Subtract. Error= %v", err)
	}

	expected := "-2"

	if expected != d3.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d3.GetNumStr())
	}

}

func TestDecimal_SubtractTotal_01(t *testing.T) {
	nStrAry := []string{
		"5.50",
		"6.50",
		"7.00",
		"8.25",
	}

	d, err := Decimal{}.NewNumStr("500.00")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"500.00\") " +
			"Error = '%v' ", err.Error())
	}

	expected := "472.75"

	for i := 0; i < len(nStrAry); i++ {
		dx, err := Decimal{}.NewNumStr(nStrAry[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(nStrAry[i]) " +
				"nStrAry[i]='%v' Error = '%v' ", nStrAry[i], err.Error())
		}

		d.SubtractTotal(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_SubtractTotal_02(t *testing.T) {
	nStrAry := []string{
		"5.50",
		"6.50",
		"7.00",
		"8.25",
	}

	d := Decimal{}.New()

	expected := "-27.25"

	for i := 0; i < len(nStrAry); i++ {
		dx, err := Decimal{}.NewNumStr(nStrAry[i])

		if err != nil {
			t.Errorf("Error returned by Decimal{}.NewNumStr(nStrAry[i]) " +
				"nStrAry[i]='%v' Error = '%v' ", nStrAry[i], err.Error())
		}

		d.SubtractTotal(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_MakeDecimalBigIntPrecision_01(t *testing.T) {
	numStr := "123456123456"
	bigI, isOk := big.NewInt(0).SetString(numStr, 10)

	if !isOk {
		t.Errorf("Failed to convert string to BigInt from numStr='%v'", numStr)
	}

	precision := uint(3)

	d1 := Decimal{}
	d2, err := d1.MakeDecimalBigIntPrecision(bigI, precision)

	if err != nil {
		t.Errorf("Error from MakeDecimalBigIntPrecision. Error:= %v \n", err)
	}

	expected := "123456123.456"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_MakeDecimalBigIntPrecision_02(t *testing.T) {
	numStr := "0"
	expected := "0.00"
	precision := uint(2)

	bigI, isOk := big.NewInt(0).SetString(numStr, 10)

	if !isOk {
		t.Errorf("Failed to convert string to BigInt from numStr='%v'", numStr)
	}

	d1 := Decimal{}
	d2, err := d1.MakeDecimalBigIntPrecision(bigI, precision)

	if err != nil {
		t.Errorf("Error from MakeDecimalBigIntPrecision. Error:= %v \n", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

}

func TestDecimal_MakeDecimalFromIntAry_01(t *testing.T) {

	numStr := "982.123456"
	precision := uint(6)
	signVal := 1

	ia, _ := IntAry{}.NewNumStr(numStr)

	d1 := Decimal{}.New()
	d2, err := d1.MakeDecimalFromIntAry(&ia)

	if err != nil {
		t.Errorf("Error retruned from d1.MakeDecimalFromIntAry(&ia). Error= %v", err)
	}

	if d2.GetNumStr() != ia.GetNumStr() {
		t.Errorf("Expected NumStr== %v .  Instead, NumStr== %v", d2.GetNumStr(), ia.GetNumStr())
	}

	if int(precision) != d2.GetPrecision() {
		t.Errorf("Expected Precision== %v .  Instead, Precision== %v", precision, d2.GetPrecision())
	}

	if signVal != d2.GetSign() {
		t.Errorf("Expected Sign Value== %v .  Instead, Sign Value== %v", signVal, d2.GetSign())
	}

}

func TestDecimal_MakeDecimalFromIntAry_02(t *testing.T) {

	numStr := "-982.123456"
	precision := uint(6)
	signVal := -1

	ia, _ := IntAry{}.NewNumStr(numStr)

	d1 := Decimal{}.New()
	d2, err := d1.MakeDecimalFromIntAry(&ia)

	if err != nil {
		t.Errorf("Error retruned from d1.MakeDecimalFromIntAry(&ia). Error= %v", err)
	}

	if err != nil {
		t.Errorf("Error retruned from d2.GetNumStr(). Error= %v", err)
	}

	if d2.GetNumStr() != ia.GetNumStr() {
		t.Errorf("Expected NumStr== %v .  Instead, NumStr== %v", d2.GetNumStr(), ia.GetNumStr())
	}


	if int(precision) != d2.GetPrecision() {
		t.Errorf("Expected Precision== %v .  Instead, Precision== %v", precision, d2.GetPrecision())
	}

	if signVal != d2.GetSign() {
		t.Errorf("Expected Sign Value== %v .  Instead, Sign Value== %v", signVal, d2.GetSign())
	}

}

func TestDecimal_Pow_01(t *testing.T) {

	numStr := "2.125"
	exp := 5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr) " +
			"numStr='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.Pow(exp, 15)

	if err != nil {
		t.Errorf("Received error from d1.Pow(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "43.330596923828125"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_Pow_02(t *testing.T) {

	numStr := "2.125"
	exp := -5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr) " +
			"numStr='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.Pow(exp, 32)

	if err != nil {
		t.Errorf("Received error from d1.Pow(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "0.02307838042845159759046157465153"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_Pow_03(t *testing.T) {

	numStr := "2.125"
	exp := -5
	d1, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr) " +
			"numStr='%v' Error = '%v' ",numStr, err.Error())
	}

	d2, err := d1.Pow(exp, 250)

	if err != nil {
		t.Errorf("Received error from d1.Pow(exp). exp='%v' Error= %v", exp, err)
	}

	expected := "0.0230783804284515975904615746515318091892352539727592285702010836302529057503678187310412245740240038257373805953698154109885713843013768287933221444131345621425256205378428954465132756326869536861810731644102187755527493261645362878092652992519669234"

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}

func TestDecimal_Inverse_01(t *testing.T) {

	numStr := "25"
	precision := uint(0)
	expected := "0.04"

	d1, err := Decimal{}.NewNumStrPrecision(numStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal d1.NewNumStrPrecision(numStr, precision, false). numStr= '%v' precision= '%v' Error= %v", numStr, precision, err)
	}

	d2, err := d1.Inverse(2)

	if err != nil {
		t.Errorf("Received error from d1.Inverse(). Error= %v", err)
	}

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}


	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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

	err := d1.SetFloat(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	if  eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}


	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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

	err := d1.SetFloat(fVal)

	if err != nil {
		t.Errorf("Received error from d1.SetFloat(fVal). fVal= '%v' Error= %v ", fVal, err)
	}

	if  eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())
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

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())
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

	if eNumStr1 != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())
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

	err := d1.SetIntFracStrings(eSignVal, intStr, fracStr)

	if err != nil {
		t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
	}

	if eNumStr != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())
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

	err := d1.SetIntFracStrings(eSignVal, intStr, fracStr)

	if err != nil {
		t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
	}


	if eNumStr != d1.GetNumStr() {
		t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
	}

	if int(ePrecision) != d1.GetPrecision() {
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())
	}

	if !d1.GetIsValid() {
		t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
	}

}

