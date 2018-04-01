package mathops

import "testing"

func TestNumStrDto_ShiftPrecisionLeft_01(t *testing.T) {

	nStr := "123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_02(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	outPrecision := uint(5)
	expected := "1234.56789"
	signVal := 1
	absIntRuneStr := "1234"
	absFracRuneStr := "56789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_03(t *testing.T) {

	nStr := "123456.789"
	precision := uint(6)
	outPrecision := uint(9)
	expected := "0.123456789"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "123456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_04(t *testing.T) {

	nStr := "123456789"
	precision := uint(6)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_05(t *testing.T) {

	nStr := "123"
	precision := uint(5)
	outPrecision := uint(5)
	expected := "0.00123"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00123"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}


func TestNumStrDto_ShiftPrecisionLeft_06(t *testing.T) {

	nStr := "0"
	precision := uint(3)
	outPrecision := uint(3)
	expected := "0.000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "000"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_07(t *testing.T) {

	nStr := "0.000"
	precision := uint(2)
	outPrecision := uint(5)
	expected := "0.00000"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00000"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_08(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_09(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "-123456.789"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_10(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_11(t *testing.T) {

	nStr := "-123456789"
	precision := uint(6)
	outPrecision := uint(6)
	expected := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_01(t *testing.T) {

	nStr := "123456.789"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "123456789"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_02(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	outPrecision := uint(1)
	expected := "12345678.9"
	signVal := 1
	absIntRuneStr := "12345678"
	absFracRuneStr := "9"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_03(t *testing.T) {

	nStr := "123456.789"
	precision := uint(6)
	outPrecision := uint(0)
	expected := "123456789000"
	signVal := 1
	absIntRuneStr := "123456789000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_04(t *testing.T) {

	nStr := "123456789"
	precision := uint(6)
	outPrecision := uint(0)
	expected := "123456789000000"
	signVal := 1
	absIntRuneStr := "123456789000000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_05(t *testing.T) {

	nStr := "123"
	precision := uint(5)
	outPrecision := uint(0)
	expected := "12300000"
	signVal := 1
	absIntRuneStr := "12300000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_06(t *testing.T) {

	nStr := "0"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "0"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_07(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_08(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "-123456.789"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_09(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "-123456789"
	signVal := -1
	absIntRuneStr := "123456789"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_10(t *testing.T) {

	nStr := "-123456789"
	precision := uint(6)
	outPrecision := uint(0)
	expected := "-123456789000000"
	signVal := -1
	absIntRuneStr := "123456789000000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionRight(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.GetNumStr() != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.GetNumStr())
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
	}

	if signVal != nsDto.GetSign() {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.GetSign())
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.GetAbsIntRunes())

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.GetAbsFracRunes())

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SubtractNumStrs_01(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "73.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_02(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-73.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_03(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "691.1"
	nStr3 := "-623.579"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_04(t *testing.T) {

	nStr1 := "691.1"
	nStr2 := "67.521"
	nStr3 := "623.579"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_05(t *testing.T) {

	nStr1 := "691.1"
	nStr2 := "0"
	nStr3 := "691.1"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_06(t *testing.T) {

	nStr1 := "0"
	nStr2 := "691.1"
	nStr3 := "-691.1"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.GetNumStr()
	expected := nResult.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.GetAbsIntRunes())
	iStr := string(nResult.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())
	fracStr := string(nResult.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.GetSign(), nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.GetPrecision(), nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

