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
	d, err := Decimal{}.NewI64(i64, precision)

	if err != nil {
		t.Errorf("Error returned by Decimal{}.NewI64(i64, precision) " +
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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Precision= '%v'. Intead, got Precision= '%v' ", ePrecision, d1.GetPrecision())
	}

	if eSignVal != d1.GetSign() {
		t.Errorf("Expected Sign Value= '%v'. Intead, got Sign Value = '%v' ", eSignVal, d1.GetSign())

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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
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
		t.Errorf("Expected Sign Value= '%v'. Instead, got Sign Value= '%v' ", eSignVal, d1.GetSign())
	}

	if int(precision) != d1.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got precision= '%v' ", precision, d1.GetPrecision())
	}

}
