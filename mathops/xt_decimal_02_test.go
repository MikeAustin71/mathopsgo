package mathops

import (
	"testing"
	)


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
		t.Errorf("Expected precision== %v .  Instead, precision== %v", precision, d2.GetPrecision())
	}

	if signVal != d2.GetSign() {
		t.Errorf("Expected sign Value== %v .  Instead, sign Value== %v", signVal, d2.GetSign())
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
		t.Errorf("Expected precision== %v .  Instead, precision== %v", precision, d2.GetPrecision())
	}

	if signVal != d2.GetSign() {
		t.Errorf("Expected sign Value== %v .  Instead, sign Value== %v", signVal, d2.GetSign())
	}

}


func TestDecimal_Mul_01(t *testing.T) {

	str1 := "575.63"
	str2 := "2014.123"
	expected := "1159389.62249"

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Mul(d2)

	if err != nil {
		t.Errorf("Error thrown on d1.Mul(d2). Error= %v", err)
	}

	if d3.GetNumStr() != expected {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, d3.GetNumStr())
	}

}

func TestDecimal_Mul_02(t *testing.T) {

	str1 := "-575.63"
	str2 := "2014.123"
	expected := "-1159389.62249"

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Mul(d2)

	if err != nil {
		t.Errorf("Error thrown on d1.Mul(d2). Error= %v", err)
	}

	if expected != d3.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, d3.GetNumStr())
	}

}
func TestDecimal_Mul_03(t *testing.T) {

	str1 := "-575.63"
	str2 := "-2014.123"
	expected := "1159389.62249"

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown on d1.SetNumStr(str2). str1= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Mul(d2)

	if err != nil {
		t.Errorf("Error thrown on d1.Mul(d2). Error= %v", err)
	}

	if expected != d3.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, d3.GetNumStr())
	}

}

func TestDecimal_MulThis_01(t *testing.T) {
	numStr := "3"
	mul, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStrDto) " +
			"numStrDto='%v'  Error = '%v' ",numStr, err.Error())
	}

	d, err := Decimal{}.NewNumStr("1")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"1\") " +
			"Error = '%v' ", err.Error())
	}

	for i := 0; i < 4; i++ {
		d.MulThis(mul)
	}

	expected := "81"

	if expected != d.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, d.GetNumStr())
	}

}


