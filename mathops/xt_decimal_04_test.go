package mathops

import "testing"

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

		d.SubtractFromThis(dx)
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

		d.SubtractFromThis(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_SubtractFromThisArray_01(t *testing.T) {

	nStrAry := []string{
		"5.50",
		"6.50",
		"7.00",
		"8.25",
	}

	decs, err := Decimal{}.NewNumStrArray(nStrAry)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrArray(nStrAry) " +
			"Error = '%v' ",
			err.Error())
	}

	total, err := Decimal{}.NewNumStr("0.0")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"0.0\") " +
			"Error = '%v' ",
			err.Error())
	}

	expected := "-27.25"

	err = total.SubtractFromThisArray(decs)

	if err != nil {
		t.Errorf("Error returned by total.SubtractFromThisArray(decs) " +
			"Error = '%v' ",
			err.Error())
	}

	if expected != total.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, total.GetNumStr())
	}

}

func TestDecimal_SubtractFromThisMultiple_01(t *testing.T) {

	dec1, err := Decimal{}.NewNumStr("5.50")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"5.50\") " +
			"Error = '%v' ",
			err.Error())
	}

	dec2, err := Decimal{}.NewNumStr("6.50")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"6.50\") " +
			"Error = '%v' ",
			err.Error())
	}

	dec3, err := Decimal{}.NewNumStr("7.00")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"7.00\") " +
			"Error = '%v' ",
			err.Error())
	}

	dec4, err := Decimal{}.NewNumStr("8.25")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"8.25\") " +
			"Error = '%v' ",
			err.Error())
	}

	total, err := Decimal{}.NewNumStr("0.0")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"0.0\") " +
			"Error = '%v' ",
			err.Error())
	}

	expected := "-27.25"

	err = total.SubtractFromThisMultiple(
												dec1,
												dec2,
												dec3,
												dec4,
									)

	if err != nil {
		t.Errorf("Error returned by total.SubtractFromThisMultiple(...) " +
			"Error = '%v' ",
			err.Error())
	}

	if expected != total.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, total.GetNumStr())
	}

}


