package mathops

import "testing"

func TestNumStrDto_ScaleNumStr_01(t *testing.T) {
	nStr := "123456"
	scale := uint(3)
	expected := "123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, scale, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if scale != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", scale, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_02(t *testing.T) {
	nStr := "-123456"
	precision := uint(3)
	expected := "-123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := -1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_03(t *testing.T) {
	nStr := "123456"
	precision := uint(9)
	expected := "123456.000000000"
	iStr := "123456"
	fracStr := "000000000"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_04(t *testing.T) {
	nStr := "123456.7890"
	precision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_05(t *testing.T) {
	nStr := "123456.789"
	precision := uint(5)
	expected := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_06(t *testing.T) {
	nStr := "-123456.789"
	precision := uint(5)
	expected := "-123456.78900"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}
func TestNumStrDto_ScaleNumStr_07(t *testing.T) {
	nStr := "123456.789"
	precision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_01(t *testing.T) {
	nStr := "123456"
	precision := uint(3)
	expected := "123.456"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_02(t *testing.T) {
	nStr := "123.456"
	precision := uint(4)
	expected := "12.3456"
	signVal := 1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_03(t *testing.T) {
	nStr := "-123456"
	precision := uint(3)
	expected := "-123.456"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_04(t *testing.T) {
	nStr := "-123.456"
	precision := uint(4)
	expected := "-12.3456"
	signVal := -1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_05(t *testing.T) {
	nStr := "123456"
	precision := uint(7)
	expected := "0.0123456"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "0123456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_06(t *testing.T) {
	nStr := "123.456"
	precision := uint(3)
	expected := "123.456"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_07(t *testing.T) {
	nStr := "123456"
	precision := uint(4)
	expected := "12.3456"
	signVal := 1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_01(t *testing.T) {

	nStr := "123456789"
	precision := uint(7)
	roundResult := false
	outPrecision := uint(7)
	expected := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_02(t *testing.T) {

	nStr := "123456789"
	precision := uint(7)
	roundResult := true
	outPrecision := uint(7)
	expected := "123456789.0000000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_03(t *testing.T) {

	nStr := "-123456789"
	precision := uint(7)
	roundResult := false
	outPrecision := uint(7)
	expected := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_04(t *testing.T) {

	nStr := "-123456789"
	precision := uint(7)
	roundResult := true
	outPrecision := uint(7)
	expected := "-123456789.0000000"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := "0000000"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_05(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	roundResult := true
	outPrecision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_06(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	roundResult := false
	outPrecision := uint(2)
	expected := "123456.78"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_07(t *testing.T) {

	nStr := "123456.789"
	precision := uint(5)
	roundResult := false
	outPrecision := uint(5)
	expected := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_08(t *testing.T) {

	nStr := "123.456789"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "123.4"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_09(t *testing.T) {

	nStr := "123.456789"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "123.5"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_10(t *testing.T) {

	nStr := "-123.456789"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "-123.4"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "4"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_11(t *testing.T) {

	nStr := "-123.456789"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "-123.5"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "5"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_12(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	roundResult := true
	outPrecision := uint(0)
	expected := "123457"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SetPrecision_13(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	roundResult := true
	outPrecision := uint(0)
	expected := "-123457"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

