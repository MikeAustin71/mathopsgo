package mathops

import (
	"fmt"

	"math/big"
	"testing"
)

/*
	These tests are associated with the library routines contained in source code file
	decimal.go. The source code repository for these these tests is located at:

				https://github.com/MikeAustin71/decimalnum.git
*/

func TestNumStrUtility_ConvertNumStrToDecimal_01(t *testing.T) {

	str := "123456.654321"
	sint := "123456654321"
	nsu := NumStrUtility{}
	dec, err := nsu.ConvertNumStrToDecimal(str)

	if err != nil {
		t.Errorf("Error from nsu.ConvertNumStrToDecimal(str). str= '%v'. Error= %v", str, err.Error())
	}

	if dec.GetNumStr() != str {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", str, dec.GetNumStr())
	}

	if dec.GetPrecision() != 6 {
		t.Errorf("Expected precision = '6'. Instead, got %v", dec.GetPrecision())
	}

	allDigits, _ := dec.GetSignedAllDigitsStr()

	if allDigits != sint {
		t.Errorf("Expected nDto.SignedAllDigitsBigInt='%v'. Instead, got %v. ", sint, allDigits)
	}

	_, accuracy, err := dec.GetFloat64()

	if accuracy.String() != "Exact" {
		t.Errorf("Expected nDto.SignedFloat64Accuracy == 'Exact'. Instead, got %v.", accuracy.String())
	}

	bf, _ := dec.GetBigFloat()
	s := fmt.Sprintf("%s", bf.Text('f', dec.GetPrecision()))

	if s != str {
		t.Errorf("Expected nDto.SignedBigFloat='%v'. Instead, got %v. ", str, s)

	}

	if !dec.GetIsValid() {
		t.Errorf("Expected nDto.isValid == 'true'. Instead, got %v", dec.GetIsValid())
	}

}

func TestNumStrUtility_ConvertNumStrToDecimal_02(t *testing.T) {

	str := "-123456.654321"
	sint := "-123456654321"
	nsu := NumStrUtility{}
	dec, err := nsu.ConvertNumStrToDecimal(str)

	if err != nil {
		t.Errorf("Error from nsu.ConvertNumStrToDecimal(str). str= '%v'. Error= %v", str, err.Error())
	}

	if str != dec.GetNumStr() {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", str, dec.GetNumStr())
	}

	if dec.GetPrecision() != 6 {
		t.Errorf("Expected precision = '6'. Instead, got %v", dec.GetPrecision())
	}

	allDigits, _ := dec.GetSignedAllDigitsStr()

	if allDigits != sint {
		t.Errorf("Expected nDto.SignedAllDigitsBigInt='%v'. Instead, got %v. ", sint, allDigits)
	}

	_, accuracy, _ := dec.GetFloat64()

	if accuracy.String() != "Exact" {
		t.Errorf("Expected nDto.SignedFloat64Accuracy == 'Exact'. Instead, got %v.", accuracy.String())
	}

	bf, _ := dec.GetBigFloat()
	s := bf.Text('f', dec.GetPrecision())

	if s != str {
		t.Errorf("Expected nDto.SignedBigFloat='%v'. Instead, got %v. ", str, s)

	}

	if !dec.GetIsValid() {
		t.Errorf("Expected nDto.isValid == 'true'. Instead, got %v", dec.GetIsValid())
	}

}

func TestNumStrUtility_ConvertNumStrToDecimal_03(t *testing.T) {

	rawStr := "zyx -123456.654321 xyx"
	str := "-123456.654321"
	sint := "-123456654321"
	nsu := NumStrUtility{}
	dec, err := nsu.ConvertNumStrToDecimal(rawStr)

	if err != nil {
		t.Errorf("Error from nsu.ConvertNumStrToDecimal(str). str= '%v'. Error= %v", str, err.Error())
	}

	if str != dec.GetNumStr() {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", str, dec.GetNumStr())
	}

	if dec.GetPrecision() != 6 {
		t.Errorf("Expected precision = '6'. Instead, got %v", dec.GetPrecision())
	}

	allDigits, _ := dec.GetSignedAllDigitsStr()

	if allDigits != sint {
		t.Errorf("Expected nDto.SignedAllDigitsBigInt='%v'. Instead, got %v. ", sint, allDigits)
	}

	_, accuracy, _ := dec.GetFloat64()

	if accuracy.String() != "Exact" {
		t.Errorf("Expected nDto.SignedFloat64Accuracy == 'Exact'. Instead, got %v.", accuracy.String())
	}

	bf, _ := dec.GetBigFloat()
	s := bf.Text('f', dec.GetPrecision())

	if s != str {
		t.Errorf("Expected nDto.SignedBigFloat='%v'. Instead, got %v. ", str, s)

	}

	if !dec.GetIsValid() {
		t.Errorf("Expected nDto.isValid == 'true'. Instead, got %v", dec.GetIsValid())
	}

}

func TestNumStrUtility_ConvertNumStrToDecimal_04(t *testing.T) {
	rawStr := "Nothing"
	nsu := NumStrUtility{}

	_, err := nsu.ConvertNumStrToDecimal(rawStr)

	if err == nil {
		t.Errorf("Expected Error to be generated by INVALID Number String. " +
			"Instead, NO ERROR WAS RETURNED! rawStr='%v'", rawStr)
	}

}

func TestDecimal_AllDigitsNumStr_01(t *testing.T) {
	numStr := "x 123.456 x"
	expected := "123456"

	d := Decimal{}.New()

	outStr, err := d.AllDigitsNumStr(numStr)

	if err != nil {
		t.Errorf("Received error from d.AllDigitsNumStr(numStrDto). Error= %v", err)
	}

	if expected != outStr {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, outStr)
	}

}

func TestDecimal_Add_01(t *testing.T) {

	numStr1 := "35.50"
	numStr2 := "35.51"
	expected := "71.01"
	nu := NumStrUtility{}

	dec1, err := nu.ConvertNumStrToDecimal(numStr1)

	if err != nil {
		t.Errorf("Received error from nu.ConvertNumStrToDecimal(numStr1). numStr1= '%v'. Error= %v", numStr1, err)
	}

	dec2, _ := nu.ConvertNumStrToDecimal(numStr2)

	if err != nil {
		t.Errorf("Received error from nu.ConvertNumStrToDecimal(numStr2). numStr2= '%v'. Error= %v", numStr2, err)
	}

	dec3, err := dec1.Add(dec2)

	if err != nil {
		t.Errorf("Received error from dec1.Add(dec2). Error= %v", err)
	}

	if expected != dec3.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, dec3.GetNumStr())
	}

	if !dec3.GetIsValid() {
		t.Error("Expected dec3.isValid='true'. Instead, got 'false'")
	}

}

func TestDecimal_Add_02(t *testing.T) {

	numStr1 := "-35.50"
	numStr2 := "35.51"
	expected := "0.01"

	dec1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1). "+
			"numStr1='%v'  Error='%v'",
			numStr1, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). "+
			"numStr1='%v'  Error='%v'",
			numStr2, err.Error())
	}

	dec3, err := dec1.Add(dec2)

	if err != nil {
		t.Errorf("Received error from dec1.Add(dec2). Error= %v", err)
	}

	if expected != dec3.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, dec3.GetNumStr())
	}

	if !dec3.GetIsValid() {
		t.Error("Expected dec3.isValid='true'. Instead, got 'false'")
	}

}

func TestDecimal_Add_03(t *testing.T) {

	numStr1 := "35.50"
	numStr2 := "35.51"
	sub1 := "71.01"
	numStr3 := "0.5"
	sub2 := "71.51"
	numStr4 := "1.00"
	sub3 := "72.51"
	expected := ""

	dec1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1). "+
			"numStr1='%v'  Error='%v'",
			numStr1, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). "+
			"numStr2='%v'  Error='%v'",
			numStr2, err.Error())
	}

	dec3, err := dec1.Add(dec2)

	if err != nil {
		t.Errorf("Received error from dec1.Add(dec2). Error= %v", err)
	}

	if !dec3.GetIsValid() {
		t.Error("Expected dec3.isValid='true'. Instead, got 'false'")
	}

	expected = sub1 // 71.01
	if expected != dec3.GetNumStr() {
		t.Errorf("Expected After decy numStr3 NumStrOut='%v'. Instead, got '%v'", expected, dec3.GetNumStr())
	}

	decy, err := Decimal{}.NewNumStr(numStr3)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr3). "+
			"numStr3='%v'  Error='%v'",
			numStr3, err.Error())
	}

	expected = "0.5"
	if expected != decy.GetNumStr() {
		t.Errorf("Expected After decy numStr3 NumStrOut='%v'. Instead, got '%v'", expected, decy.GetNumStr())
	}

	dec4, err := dec3.Add(decy)

	expected = sub2 // 71.51
	if expected != dec4.GetNumStr() {
		t.Errorf("Expected After dec4 numStr3 NumStrOut='%v'. Instead, got '%v'", expected, dec4.GetNumStr())
	}

	dec2.SetNumStr(numStr4)

	expected = "1.00"

	if expected != dec2.GetNumStr() {
		t.Errorf("Expected After reuse of dec2 NumStrOut='%v'. Instead, got '%v'", expected, dec2.GetNumStr())
	}

	// Re-use dec3
	dec3, err = dec4.Add(dec2)

	expected = sub3 // 72.56

	if expected != dec3.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, dec3.GetNumStr())
	}

}

func TestDecimal_Add_04(t *testing.T) {
	numStr1 := "35.50"
	numStr2 := ".5"
	expected := "36.00"

	d1 := Decimal{}

	d1.SetNumStr(numStr1)
	d2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). "+
			"numStr2='%v'  Error='%v'",
			numStr2, err.Error())
	}

	d3, _ := d2.Add(d1)

	if expected != d3.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, d3.GetNumStr())
	}

}

func TestDecimal_Add_05(t *testing.T) {
	numStr1 := "35.50"
	numStr2 := ".5"
	numStr3 := "1.00"
	numStr4 := "9.32"
	numStr5 := "101.912"
	expected := "148.232"

	d1 := Decimal{}

	d1.SetNumStr(numStr1)
	d2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). "+
			"numStr2='%v'  Error='%v'",
			numStr2, err.Error())
	}

	d3, _ := d2.Add(d1)

	d2.SetNumStr(numStr3)
	d3, _ = d3.Add(d2)

	d2.SetNumStr(numStr4)
	d3, _ = d3.Add(d2)

	d2.SetNumStr(numStr5)
	d3, _ = d3.Add(d2)

	if expected != d3.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, d3.GetNumStr())
	}

}

func TestDecimal_AddToThis_01(t *testing.T) {

	nStrAry := []string{
		"35.50",
		"36.50",
		"5.5",
		"92.75",
	}

	expected := "170.25"

	d, err := Decimal{}.NewNumStr("0")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"0\"). "+
			"Error='%v'",
			err.Error())
	}

	for i := 0; i < len(nStrAry); i++ {
		dx := Decimal{}
		dx.SetNumStr(nStrAry[i])
		d.AddToThis(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, d.GetNumStr())
	}

}

func TestDecimal_AddToThis_02(t *testing.T) {

	nStrAry := []string{
		"35.50",
		"36.50",
		"5.5",
		"92.75",
	}

	expected := "170.25"

	d := Decimal{}.New()

	for i := 0; i < len(nStrAry); i++ {
		dx := Decimal{}
		dx.SetNumStr(nStrAry[i])
		d.AddToThis(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, d.GetNumStr())
	}

}

func TestDecimal_AddToThis_03(t *testing.T) {

	nStrAry := []string{
		"35.50",
		"36.50",
		"5.5",
		"92.75",
	}

	expected := "320.25"

	d, err := Decimal{}.NewNumStr("150")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\"). "+
			"Error='%v'",
			err.Error())
	}

	for i := 0; i < len(nStrAry); i++ {
		dx := Decimal{}
		dx.SetNumStr(nStrAry[i])
		d.AddToThis(dx)
	}

	if expected != d.GetNumStr() {
		t.Errorf("Expected NumStrOut='%v'. Instead, got '%v'", expected, d.GetNumStr())
	}

}

func TestDecimal_AddToThisArray_01(t *testing.T) {
	ePrecision := 4
	eSignVal := 1
	expected := "68139.6265"
	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") "+
			"Error = '%v' ", err.Error())
	}

	d, err := Decimal{}.NewNumStrsMultiple(
		"351.7",
		"6224.894",
		"34.8",
		"150",
		"150726.9",
		"-89421.6175",
		"47.23",
	)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrsMultiple(...) "+
			"Error = '%v' ", err.Error())
	}

	/*
		d:= [] Decimal { Decimal{}.NewNumStr("351.7"),
			Decimal{}.NewNumStr("6224.894"),
			Decimal{}.NewNumStr("34.8"),
			Decimal{}.NewNumStr("150"),
			Decimal{}.NewNumStr("150726.9"),
			Decimal{}.NewNumStr("-89421.6175"),
			Decimal{}.NewNumStr("47.23"),
		}
	*/

	err = base.AddToThisArray(d)

	if err != nil {
		t.Errorf("Received error from base.AddToThisArray(d). Error= %v", err)
	}

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected precision== %v . Instead, precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected sign Value == %v . Instead, sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisArray_02(t *testing.T) {
	ePrecision := 4
	eSignVal := -1
	expected := "-233314.1735"
	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") "+
			"Error = '%v' ", err.Error())
	}

	d, err := Decimal{}.NewNumStrsMultiple(
		"351.7",
		"6224.894",
		"34.8",
		"150",
		"-150726.9",
		"-89421.6175",
		"47.23",
	)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStrsMultiple(...) "+
			"Error = '%v' ", err.Error())
	}

	/*
		d:= [] Decimal { Decimal{}.NewNumStr("351.7"),
			Decimal{}.NewNumStr("6224.894"),
			Decimal{}.NewNumStr("34.8"),
			Decimal{}.NewNumStr("150"),
			Decimal{}.NewNumStr("-150726.9"),
			Decimal{}.NewNumStr("-89421.6175"),
			Decimal{}.NewNumStr("47.23"),
		}
	*/

	err = base.AddToThisArray(d)

	if err != nil {
		t.Errorf("Received error from base.AddToThisArray(d). Error= %v", err)
	}

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected precision== %v . Instead, precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected sign Value == %v . Instead, sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisSeries_01(t *testing.T) {

	ePrecision := 3
	eSignVal := 1

	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") "+
			"Error = '%v' ", err.Error())
	}

	d1, err := Decimal{}.NewNumStr("150")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\") "+
			"Error = '%v' ", err.Error())
	}

	d2, err := Decimal{}.NewNumStr("650.25")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"650.25\") "+
			"Error = '%v' ", err.Error())
	}

	d3, err := Decimal{}.NewNumStr("20.625")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"20.625\") "+
			"Error = '%v' ", err.Error())
	}

	d4, err := Decimal{}.NewNumStr("-16.59")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-16.59\") "+
			"Error = '%v' ", err.Error())
	}

	err = base.AddToThisSeries(d1, d2, d3, d4)

	if err != nil {
		t.Errorf("Received error from base.AddToThisSeries(d1, d2, d3, d4). Error= %v", err)
	}

	expected := "830.005"

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected precision== %v . Instead, precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected sign Value == %v . Instead, sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisSeries_02(t *testing.T) {

	ePrecision := 3
	eSignVal := -1

	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") "+
			"Error = '%v' ", err.Error())
	}

	d1, err := Decimal{}.NewNumStr("150")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\") "+
			"Error = '%v' ", err.Error())
	}

	d2, err := Decimal{}.NewNumStr("-6050.25")
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-6050.25\") "+
			"Error = '%v' ", err.Error())
	}

	d3, err := Decimal{}.NewNumStr("20.625")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"20.625\") "+
			"Error = '%v' ", err.Error())
	}

	d4, err := Decimal{}.NewNumStr("-16.59")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-16.59\") "+
			"Error = '%v' ", err.Error())
	}

	base.AddToThisSeries(d1, d2, d3, d4)

	expected := "-5870.495"

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected precision== %v . Instead, precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected sign Value == %v . Instead, sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_Cmp_01(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.455"
	expectedCmpResult := 1

	decNum1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	decNum2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := decNum1.Cmp(decNum2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestDecimal_Cmp_02(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.457"
	expectedCmpResult := -1

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := dec1.Cmp(dec2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestDecimal_Cmp_03(t *testing.T) {

	n1Str := "123.456"
	n2Str := "123.456"
	expectedCmpResult := 0

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := dec1.Cmp(dec2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestDecimal_Cmp_04(t *testing.T) {

	n1Str := "-123.456"
	n2Str := "-123.457"
	expectedCmpResult := 1

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := dec1.Cmp(dec2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestDecimal_Cmp_05(t *testing.T) {

	n1Str := "-123.456"
	n2Str := "-123.455"
	expectedCmpResult := -1

	dec1, err := Decimal{}.NewNumStr(n1Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
			"n1Str='%v' Error='%v'", n1Str, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(n2Str)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
			"n2Str='%v' Error='%v'", n2Str, err.Error())
	}

	cmpResult := dec1.Cmp(dec2)

	if expectedCmpResult != cmpResult {
		t.Errorf("Error: Expected Comparision Result='%v'.  Instead, Result='%v'",
			expectedCmpResult, cmpResult)
	}

}

func TestDecimal_CubeRoot_01(t *testing.T) {
	numStr1 := "2686.5"
	maxPrecision := uint(30)
	expected := "13.901519768959674425418091364468"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
			"numStr1='%v' Error = '%v' ", numStr1, err.Error())
	}

	decCubeRoot, err := d1.CubeRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.CubeRoot(maxPrecision). Error = %v ", err)
	}

	if expected != decCubeRoot.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, decCubeRoot.GetNumStr())
	}

	if eSignVal != decCubeRoot.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, decCubeRoot.GetSign())
	}

	if int(maxPrecision) != decCubeRoot.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, decCubeRoot.GetPrecision())
	}
}

func TestDecimal_CubeRoot_02(t *testing.T) {
	numStr1 := "390626"
	maxPrecision := uint(29)
	expected := "73.10050583431350346938081010498"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
			"numStr1='%v' Error = '%v' ", numStr1, err.Error())
	}

	decCubeRoot, err := d1.CubeRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.CubeRoot(maxPrecision). Error = %v ", err)
	}

	if expected != decCubeRoot.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, decCubeRoot.GetNumStr())
	}

	if eSignVal != decCubeRoot.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, decCubeRoot.GetSign())
	}

	if int(maxPrecision) != decCubeRoot.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, decCubeRoot.GetPrecision())
	}
}

func TestDecimal_CubeRoot_03(t *testing.T) {
	numStr1 := "-390626"
	maxPrecision := uint(29)
	expected := "-73.10050583431350346938081010498"
	eSignVal := -1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) "+
			"numStr1='%v' Error = '%v' ", numStr1, err.Error())
	}

	decCubeRoot, err := d1.CubeRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.CubeRoot(maxPrecision). Error = %v ", err)
	}

	if expected != decCubeRoot.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, decCubeRoot.GetNumStr())
	}

	if eSignVal != decCubeRoot.GetSign() {
		t.Errorf("Expected sign Value= '%v'. Instead, got sign Value= '%v' ", eSignVal, decCubeRoot.GetSign())
	}

	if int(maxPrecision) != decCubeRoot.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, decCubeRoot.GetPrecision())
	}

}

func TestDecimal_Divide_01(t *testing.T) {
	// str1 / str2
	str1 := "575.63"
	str2 := "2014.123"
	ePrecision := uint(20)
	expected := "0.28579684557497233287"

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown by d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown by d2.SetNumStr(str2). str2= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Divide(d2, ePrecision)

	if err != nil {
		t.Errorf("Error thrown by d1.Divide(d2, 20).  Error= %v", err)
	}

	outNumStr, _ := d3.GetBigFloatString(uint(d3.GetPrecision()))

	if outNumStr != expected {
		t.Errorf("Expected Quotient %v. Instead, got %v", expected, outNumStr)
	}

	actualResult := d3.GetNumStr()

	ia1, err := IntAry{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(str2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())
		return
	}

	ia3, err := ia1.DivideThisBy(&ia2, 0, 20)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). "+
			"Error='%v \n", err.Error())
		return
	}

	chkResult := ia3.GetNumStr()

	if chkResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. "+
			chkResult, actualResult)
	}

}

func TestDecimal_Divide_02(t *testing.T) {
	// str1 / str2
	str1 := "-8076.63"
	str2 := "12014.123"
	ePrecision := uint(20)

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown by d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown by d2.SetNumStr(str2). str2= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Divide(d2, ePrecision)

	if err != nil {
		t.Errorf("Error thrown by d1.Divide(d2, 20).  Error= %v", err)
	}

	actualResult := d3.GetNumStr()

	ia1, err := IntAry{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(str2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())
		return
	}

	ia3, err := ia1.DivideThisBy(&ia2, 0, 20)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). "+
			"Error='%v \n", err.Error())
		return
	}

	chkResult := ia3.GetNumStr()

	if chkResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. "+
			chkResult, actualResult)
	}

}

func TestDecimal_Divide_03(t *testing.T) {
	// str1 / str2
	str1 := "100"
	str2 := "33"
	maxPrecision := uint(20)

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown by d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown by d2.SetNumStr(str2). str2= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Divide(d2, maxPrecision)

	if err != nil {
		t.Errorf("Error thrown by d1.Divide(d2, 20).  Error= %v", err)
	}

	actualResult := d3.GetNumStr()

	ia1, err := IntAry{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(str2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())
		return
	}

	ia3, err := ia1.DivideThisBy(&ia2, 0, 20)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). "+
			"Error='%v \n", err.Error())
		return
	}

	chkResult := ia3.GetNumStr()

	if chkResult != actualResult {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. "+
			chkResult, actualResult)
	}

}

func TestDecimal_Divide_04(t *testing.T) {
	// str1 / str2
	str1 := "975.69"
	str2 := "589.7654321"
	excelResult := "1.654369597"
	maxPrecision := uint(9)

	d1 := Decimal{}.New()

	err := d1.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error thrown by d1.SetNumStr(str1). str1= '%v' Error= %v", str1, err)
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStr(str2)

	if err != nil {
		t.Errorf("Error thrown by d2.SetNumStr(str2). str2= '%v' Error= %v", str2, err)
	}

	d3, err := d1.Divide(d2, maxPrecision)

	if err != nil {
		t.Errorf("Error thrown by d1.Divide(d2, 20).  Error= %v", err)
	}

	actualResult := d3.GetNumStr()

	ia1, err := IntAry{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str1). "+
			"str1='%v' Error='%v \n", str1, err.Error())
	}

	ia2, err := IntAry{}.NewNumStr(str2)

	if err != nil {
		t.Errorf("Error returned from IntAry{}.NewNumStr(str2). "+
			"str2='%v' Error='%v \n", str2, err.Error())

	}

	ia3, err := ia1.DivideThisBy(&ia2, 0, int(maxPrecision))

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). "+
			"Error='%v \n", err.Error())

	}

	chkResult := ia3.GetNumStr()

	if chkResult != actualResult {
		t.Errorf("Error: IntAry Expected result='%v'. Instead, result='%v'. "+
			chkResult, actualResult)
	}

	if excelResult != actualResult {
		t.Errorf("Error: Excel Expected result='%v'. Instead, result='%v'. ",
			excelResult, actualResult)
	}

}

func TestDecimal_GetNumStr_01(t *testing.T) {

	str1 := "575.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetNumStr()

	if str1 != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			str1, actualStr)
	}

}

func TestDecimal_GetNumStr_02(t *testing.T) {

	str1 := "-575.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetNumStr()

	if str1 != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			str1, actualStr)
	}

}

func TestDecimal_GetNumParen_01(t *testing.T) {

	str1 := "575.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetNumParen()

	if str1 != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			str1, actualStr)
	}

}

func TestDecimal_GetNumParen_02(t *testing.T) {

	str1 := "-575.63"
	expectedStr := "(575.63)"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetThouStr_01(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetThouStr_02(t *testing.T) {

	str1 := "-2567894.63"

	expectedStr := "-2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetThouStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetThouParen_01(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetThouParen_02(t *testing.T) {

	str1 := "-2567894.63"

	expectedStr := "(2,567,894.63)"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetThouParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyStr_01(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "$2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyStr_02(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "€2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	// '\U000020ac', // Euro €  													 7
	d1.SetCurrencySymbol('\U000020ac')

	actualStr := d1.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyStr_03(t *testing.T) {

	str1 := "-2567894.63"

	expectedStr := "-$2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyParen_01(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "$2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	actualStr := d1.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyParen_02(t *testing.T) {

	str1 := "2567894.63"

	expectedStr := "€2,567,894.63"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	// '\U000020ac', // Euro €  													 7
	d1.SetCurrencySymbol('\U000020ac')

	actualStr := d1.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetCurrencyParen_03(t *testing.T) {

	str1 := "-2567894.63"

	expectedStr := "(£2,567,894.63)"

	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1). "+
			"Error='%v' ", err.Error())
	}

	// '\U000000a3', // United Kingdom Pound (£)					29
	d1.SetCurrencySymbol('\U000000a3')

	actualStr := d1.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead NumStr='%v' ",
			expectedStr, actualStr)
	}

}

func TestDecimal_GetAbsoluteAllDigitsStr_01(t *testing.T) {
	str1 := "-123456.123456"
	expected := "123456123456"
	d := Decimal{}.New()
	d.SetNumStr(str1)

	absStr, err := d.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("d.GetAbsoluteAllDigitsStr() returned an Error. Error= %v", err)
	}

	if absStr != expected {
		t.Errorf("Expected %v. Instead, got %v", expected, absStr)
	}

}

func TestDecimal_GetAbsoluteAllDigitsStr_02(t *testing.T) {
	str1 := "123456.123456"
	expected := "123456123456"
	d := Decimal{}.New()
	d.SetNumStr(str1)

	absStr, err := d.GetAbsoluteAllDigitsStr()

	if err != nil {
		t.Errorf("d.GetAbsoluteAllDigitsStr() returned an Error. Error= %v", err)
	}

	if absStr != expected {
		t.Errorf("Expected %v. Instead, got %v", expected, absStr)
	}

}

func TestDecimal_GetAbsoluteValue_01(t *testing.T) {
	str1 := "-123456.123456"
	expected := "123456.123456"

	d := Decimal{}.New()

	err := d.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from NewNumStr(str1). str1='%v' Error= %v", str1, err)
	}

	dAbs := d.GetAbsoluteValue()

	if expected != dAbs.GetNumStr() {
		t.Errorf("Expected absolute value = '%v'. Instead, got '%v'.", expected, dAbs.GetNumStr())
	}

}

func TestDecimal_GetAbsoluteValue_02(t *testing.T) {
	str1 := "123456.123456"
	expected := "123456.123456"

	d := Decimal{}.New()

	err := d.SetNumStr(str1)

	if err != nil {
		t.Errorf("Error returned from NewNumStr(str1). str1='%v' Error= %v", str1, err)
	}

	dAbs := d.GetAbsoluteValue()

	if expected != dAbs.GetNumStr() {
		t.Errorf("Expected absolute value = '%v'. Instead, got '%v'.", expected, dAbs.GetNumStr())
	}

}

func TestDecimal_GetBigIntNum_01(t *testing.T) {

	bigI := big.NewInt(int64(123456123456))
	precision := uint(6)

	exStr := "123456.123456"

	expectedBigIntNum := BigIntNum{}.NewBigInt(bigI, precision)

	dec := Decimal{}.NewBigInt(bigI, precision)

	bigINum, err := dec.GetBigIntNum()

	if err != nil {
		t.Errorf("Error returned by dec.GetBigIntNum(). Error='%v'",
			err.Error())
	}

	if !expectedBigIntNum.Equal(bigINum) {
		t.Errorf("Error: Expected BigIntNum NOT Equal to Actual BigIntNum! "+
			"expectedBi='%v', expectedPrecision='%v'. actualBi='%v' actualPrecision='%v'",
			expectedBigIntNum.bigInt.Text(10), expectedBigIntNum.precision,
			bigINum.bigInt.Text(10), bigINum.precision)
	}

	actualNumStr := bigINum.GetNumStr()

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if err != nil {
		t.Errorf("Error returned by bigINum.GetNumStrErr(). "+
			"Error='%v' ", err.Error())
	}

	if exStr != actualNumStr {
		t.Errorf("Error: Expected NumStr='%v'.  Instead, NumStr='%v'",
			exStr, actualNumStr)
	}
}

func TestDecimal_GetIntAry_01(t *testing.T) {
	bigI := big.NewInt(int64(123456123456))
	precision := uint(6)
	exStr := "123456.123456"
	d := Decimal{}.NewBigInt(bigI, precision)

	signVal := 1

	ia, err := d.GetIntAry()

	if err != nil {
		t.Errorf("Error returned from d.GetIntAryElements(). Error= %v ", err)
	}

	if exStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.Precsion== %v .   Instead, ia.precision== %v", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
	}

}

func TestDecimal_GetIntAry_02(t *testing.T) {
	bigI := big.NewInt(int64(-123456123456))
	precision := uint(6)
	exStr := "-123456.123456"
	d := Decimal{}.NewBigInt(bigI, precision)

	signVal := -1

	ia, err := d.GetIntAry()

	if err != nil {
		t.Errorf("Error returned from d.GetIntAryElements(). Error= %v ", err)
	}

	if exStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.Precsion== %v .   Instead, ia.precision== %v", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
	}

}

func TestDecimal_GetRelevantPrecision(t *testing.T) {
	str1 := "-2.0105000"
	expected := uint(4)
	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) "+
			"str1='%v' Error = '%v' ", str1, err.Error())
	}

	rP := d1.GetRelevantPrecision()

	if rP != expected {
		t.Errorf("Expected Relevant precision = %v. Instead, got %v", expected, rP)
	}

}

func TestDecimal_Inverse_01(t *testing.T) {

	numStr := "25"
	precision := uint(0)
	expected := "0.04"

	d1, err := Decimal{}.NewNumStrPrecision(numStr, precision, false)

	if err != nil {
		t.Errorf("Error Returned from Decimal d1.NewNumStrPrecision(numStrDto, precision, false). numStrDto= '%v' precision= '%v' Error= %v", numStr, precision, err)
	}

	d2, err := d1.Inverse(2)

	if err != nil {
		t.Errorf("Received error from d1.Inverse(). Error= %v", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr= '%v'. Instead, got %v.", expected, d2.GetNumStr())
	}

}
