package mathops

import "testing"

func TestDecimal_ShiftPrecisionLeft_01(t *testing.T) {

	nStr1 := "123456789"
	expectedNumStr := "123456.789"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(3)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionLeft_02(t *testing.T) {

	nStr1 := "12345"
	expectedNumStr := "0.012345"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(6)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionLeft_03(t *testing.T) {

	nStr1 := "12345"
	expectedNumStr := "0.000012345"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(9)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionLeft_04(t *testing.T) {

	nStr1 := "-123456789"
	expectedNumStr := "-123456.789"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(3)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionLeft_05(t *testing.T) {

	nStr1 := "-12345"
	expectedNumStr := "-0.012345"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(6)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionLeft_06(t *testing.T) {

	nStr1 := "-12345"
	expectedNumStr := "-0.000012345"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionLeft(9)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionRight_01(t *testing.T) {

	nStr1 := "123456789"
	expectedNumStr := "123456789000"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(3)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}


func TestDecimal_ShiftPrecisionRight_02(t *testing.T) {

	nStr1 := "12345"
	expectedNumStr := "12345000000"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(6)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionRight_03(t *testing.T) {

	nStr1 := "1.2345"
	expectedNumStr := "1234500000"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(9)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionRight_04(t *testing.T) {

	nStr1 := "-123456.789"
	expectedNumStr := "-123456789"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(3)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionRight_05(t *testing.T) {

	nStr1 := "-12.345"
	expectedNumStr := "-1234.5"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(2)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}

func TestDecimal_ShiftPrecisionRight_06(t *testing.T) {

	nStr1 := "-12.345"
	expectedNumStr := "-123450"

	d1, err := Decimal{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(nStr1). " +
			"Error='%v' ", err.Error())
	}

	err = d1.ShiftPrecisionRight(4)

	if err != nil {
		t.Errorf("Error returned by d1.ShiftPrecisionLeft(3). " +
			"Error='%v' ", err.Error())
	}

	actualNumStr := d1.GetNumStr()

	if expectedNumStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'.",
			expectedNumStr, actualNumStr)
	}

}


func TestDecimal_SquareRoot_01(t *testing.T) {
	numStr1 := "2686.5"
	maxPrecision := 30
	expected := "51.831457629512986714934518985668"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.SquareRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.SquareRoot(maxPrecision). Error = %v ", err)
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

func TestDecimal_SquareRoot_02(t *testing.T) {
	numStr1 := "390626"
	maxPrecision := 29
	expected := "625.00079999948800065535895142588"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.SquareRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.SquareRoot(maxPrecision). Error = %v ", err)
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

func TestDecimal_SquareRoot_03(t *testing.T) {
	numStr1 := "-390626"
	maxPrecision := 29

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	_, err = d1.SquareRoot(maxPrecision)

	if err == nil {
		t.Error("It was expected that an error would be generated when attempting to calculate the square root of a negative number. However, no such error was triggered!")
	}

}

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

		err = d.SubtractFromThis(dx)

		if err != nil {
			t.Errorf("Error returned from d.SubtractFromThis(dx) " +
				"Error='%v' ", err.Error())
		}
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

		err = d.SubtractFromThis(dx)

		if err != nil {
			t.Errorf("Error returned from d.SubtractFromThis(dx). " +
				"Error='%v' ", err.Error())
		}

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


