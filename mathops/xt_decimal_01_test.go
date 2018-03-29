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

	d2, err := nsu.ConvertNumStrToDecimal(rawStr)

	if err != nil {
		t.Errorf("Error generated by ConvertNumStrToDecimal(rawStr). rawStr= '%v' Error:= %v", rawStr, err.Error())
	}

	f64, _, _ := d2.GetFloat64()

	if f64 != 0.0 {
		t.Errorf("Expected result= 0.0 . Instead, got %v", f64)
	}

}

func TestDecimal_AllDigitsNumStr_01(t *testing.T) {
	numStr := "x 123.456 x"
	expected := "123456"

	d := Decimal{}.New()

	outStr, err := d.AllDigitsNumStr(numStr)

	if err != nil {
		t.Errorf("Received error from d.AllDigitsNumStr(numStr). Error= %v", err)
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1). " +
			"numStr1='%v'  Error='%v'",
				numStr1, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1). " +
			"numStr1='%v'  Error='%v'",
			numStr1, err.Error())
	}

	dec2, err := Decimal{}.NewNumStr(numStr2)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr3). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr2). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"0\"). " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\"). " +
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
	expected:= "68139.6265"
	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStrsMultiple(...) " +
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

	if err!=nil {
		t.Errorf("Received error from base.AddToThisArray(d). Error= %v", err)
	}

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected Precision== %v . Instead, Precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected Sign Value == %v . Instead, Sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisArray_02(t *testing.T) {
	ePrecision := 4
	eSignVal := -1
	expected:= "-233314.1735"
	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") " +
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
		t.Errorf("Error returned by Decimal{}.NewNumStrsMultiple(...) " +
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

	if err!=nil {
		t.Errorf("Received error from base.AddToThisArray(d). Error= %v", err)
	}

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected Precision== %v . Instead, Precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected Sign Value == %v . Instead, Sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisMultiple_01(t *testing.T) {

	ePrecision := 3
	eSignVal := 1

	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") " +
			"Error = '%v' ", err.Error())
	}



	d1, err:= Decimal{}.NewNumStr("150")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\") " +
			"Error = '%v' ", err.Error())
	}


	d2, err := Decimal{}.NewNumStr("650.25")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"650.25\") " +
			"Error = '%v' ", err.Error())
	}

	d3, err:= Decimal{}.NewNumStr("20.625")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"20.625\") " +
			"Error = '%v' ", err.Error())
	}

	d4, err := Decimal{}.NewNumStr("-16.59")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-16.59\") " +
			"Error = '%v' ", err.Error())
	}

	err = base.AddToThisMultiple(d1, d2, d3, d4)

	if err!=nil {
		t.Errorf("Received error from base.AddToThisMultiple(d1, d2, d3, d4). Error= %v", err)
	}

	expected := "830.005"

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected Precision== %v . Instead, Precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected Sign Value == %v . Instead, Sign Value == %v .", eSignVal, base.GetSign())
	}

}

func TestDecimal_AddToThisMultiple_02(t *testing.T) {

	ePrecision := 3
	eSignVal := -1

	base, err := Decimal{}.NewNumStr("25.72")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"25.72\") " +
			"Error = '%v' ", err.Error())
	}

	d1, err:= Decimal{}.NewNumStr("150")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"150\") " +
			"Error = '%v' ", err.Error())
	}

	d2, err:= Decimal{}.NewNumStr("-6050.25")
	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-6050.25\") " +
			"Error = '%v' ", err.Error())
	}

	d3, err:= Decimal{}.NewNumStr("20.625")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"20.625\") " +
			"Error = '%v' ", err.Error())
	}

	d4, err:= Decimal{}.NewNumStr("-16.59")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"-16.59\") " +
			"Error = '%v' ", err.Error())
	}

	base.AddToThisMultiple(d1, d2, d3, d4)

	expected := "-5870.495"

	if expected != base.GetNumStr() {
		t.Errorf("Expected NumStr== %v . Instead, NumStr== %v .", expected, base.GetNumStr())
	}

	if ePrecision != base.GetPrecision() {
		t.Errorf("Expected Precision== %v . Instead, Precision== %v .", ePrecision, base.GetPrecision())
	}

	if eSignVal != base.GetSign() {
		t.Errorf("Expected Sign Value == %v . Instead, Sign Value == %v .", eSignVal, base.GetSign())
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

func TestDecimal_MulTotal_01(t *testing.T) {
	numStr := "3"
	mul, err := Decimal{}.NewNumStr(numStr)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr) " +
			"numStr='%v'  Error = '%v' ",numStr, err.Error())
	}

	d, err := Decimal{}.NewNumStr("1")

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(\"1\") " +
			"Error = '%v' ", err.Error())
	}

	for i := 0; i < 4; i++ {
		d.MulTotal(mul)
	}

	expected := "81"

	if expected != d.GetNumStr() {
		t.Errorf("Error. Expected %v. Instead, got %v", expected, d.GetNumStr())
	}

}

func TestDecimal_Divide_01(t *testing.T) {

	str1 := "575.63"
	str2 := "2014.123"
	expected := "0.28579684557497233287"
	//actual   := "0.28579684557497231356"

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

	d3, err := d1.Divide(d2, 20)

	if err != nil {
		t.Errorf("Error thrown by d1.Divide(d2, 20).  Error= %v", err)
	}

	outNumStr, _ := d3.GetBigFloatString(uint(d3.GetPrecision()))

	if outNumStr != expected {
		t.Errorf("Expected Quotient %v. Instead, got %v", expected, outNumStr)
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

	dAbs, err := d.GetAbsoluteValue()

	if err != nil {
		t.Errorf("Error returned from d.GetAbsoluteValue(). Error= %v", err)
	}

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

	dAbs, err := d.GetAbsoluteValue()

	if err != nil {
		t.Errorf("Error returned from d.GetAbsoluteValue(). Error= %v", err)
	}

	if expected != dAbs.GetNumStr() {
		t.Errorf("Expected absolute value = '%v'. Instead, got '%v'.", expected, dAbs.GetNumStr())
	}

}

func TestDecimal_GetIntAry_01(t *testing.T) {
	bigI := big.NewInt(int64(123456123456))
	precision := uint(6)
	exStr := "123456.123456"
	d, err := Decimal{}.NewBigInt(bigI, precision)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewBigInt(bigI, precision). " +
			"Error='%v' ", err.Error())
	}


	signVal := 1

	ia, err := d.GetIntAry()

	if err != nil {
		t.Errorf("Error returned from d.GetIntAry(). Error= %v ", err)
	}

	if exStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.Precsion== %v .   Instead, ia.Precision== %v", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
	}

}

func TestDecimal_GetIntAry_02(t *testing.T) {
	bigI := big.NewInt(int64(-123456123456))
	precision := uint(6)
	exStr := "-123456.123456"
	d, err := Decimal{}.NewBigInt(bigI, precision)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewBigInt(bigI, precision). " +
			"bigI='%v' precision='%v' Error='%v' ",
				bigI.String(), precision, err.Error())
	}

	signVal := -1

	ia, err := d.GetIntAry()

	if err != nil {
		t.Errorf("Error returned from d.GetIntAry(). Error= %v ", err)
	}

	if exStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v .  Instead ia.GetNumStr() == %v ", exStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.Precsion== %v .   Instead, ia.Precision== %v", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.SignVal== %v .   Instead, ia.SignVal== %v", signVal, ia.GetSign())
	}

}

func TestDecimal_GetNthRoot_01(t *testing.T) {
	numStr1 := "125"
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",
			 numStr1, err.Error())
	}

	d2, err := d1.GetNthRoot(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetNthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_GetNthRoot_02(t *testing.T) {
	numStr1 := "5604423"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276982415963"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.GetNthRoot(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetNthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_GetNthRoot_03(t *testing.T) {
	numStr1 := "5604423.924"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276986078187"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.GetNthRoot(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetNthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_GetNthRoot_04(t *testing.T) {
	numStr1 := "-27"
	nthRoot := uint(3)
	maxPrecision := uint(2)
	expected := "-3.00"
	eSignVal := -1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.GetNthRoot(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetNthRoot(nthRoot, maxPrecision). Error= %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}

}

func TestDecimal_GetNthRoot_05(t *testing.T) {
	numStr1 := "-27"
	nthRoot := uint(4)
	maxPrecision := uint(2)

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	_, err = d1.GetNthRoot(nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from d1.GetNthRoot(nthRoot, maxPrecision) for negative number with even nthRoot. No Error triggered")
	}


}

func TestDecimal_GetRelevantPrecision(t *testing.T) {
	str1 := "-2.0105000"
	expected := uint(4)
	d1, err := Decimal{}.NewNumStr(str1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(str1) " +
			"str1='%v' Error = '%v' ",str1, err.Error())
	}

	rP := d1.GetRelevantPrecision()

	if rP != expected {
		t.Errorf("Expected Relevant Precision = %v. Instead, got %v", expected, rP)
	}

}

func TestDecimal_GetSquareRoot_01(t *testing.T) {
	numStr1 := "2686.5"
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.GetSquareRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetSquareRoot(maxPrecision). Error = %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}
}

func TestDecimal_GetSquareRoot_02(t *testing.T) {
	numStr1 := "390626"
	maxPrecision := uint(29)
	expected := "625.00079999948800065535895142588"
	eSignVal := 1

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	d2, err := d1.GetSquareRoot(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from d1.GetSquareRoot(maxPrecision). Error = %v ", err)
	}

	if expected != d2.GetNumStr() {
		t.Errorf("Expected NumStr: %v. Instead, got %v", expected, d2.GetNumStr())
	}

	if eSignVal != d2.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d2.GetSign())
	}

	if int(maxPrecision) != d2.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", maxPrecision, d2.GetPrecision())
	}
}

func TestDecimal_GetSquareRoot_03(t *testing.T) {
	numStr1 := "-390626"
	maxPrecision := uint(29)

	d1, err := Decimal{}.NewNumStr(numStr1)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewNumStr(numStr1) " +
			"numStr1='%v' Error = '%v' ",numStr1, err.Error())
	}

	_, err = d1.GetSquareRoot(maxPrecision)

	if err == nil {
		t.Error("It was expected that an error would be generated when attempting to calculate the square root of a negative number. However, no such error was triggered!")
	}

}
