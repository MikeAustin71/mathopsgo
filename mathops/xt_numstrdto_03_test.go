package mathops

import "testing"

func TestNumStrDto_ScaleNumStr_01(t *testing.T) {
	// "123456.789"				  3						"123.456789"
	nStr := "123456.789"
	scaleMode := SCALEPRECISIONLEFT
	scale := uint(3)
	ePrecision := uint(6)
	expected := "123.456789"
	iStr := "123"
	fracStr := "456789"
	signVal := 1


	nsDto, err := NumStrDto{}.NewPtr().ScaleNumStr(nStr, scale, scaleMode)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if ePrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", scale, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if precision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_14(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	roundResult := false
	outPrecision := uint(0)
	expected := "123456"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_15(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(0)
	roundResult := false
	outPrecision := uint(0)
	expected := "-123456"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_16(t *testing.T) {

	nStr := "123457"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_17(t *testing.T) {

	nStr := "123457"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "123457.0"
	signVal := 1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_18(t *testing.T) {

	nStr := "-123457"
	precision := uint(1)
	roundResult := false
	outPrecision := uint(1)
	expected := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetPrecision_19(t *testing.T) {

	nStr := "-123457"
	precision := uint(1)
	roundResult := true
	outPrecision := uint(1)
	expected := "-123457.0"
	signVal := -1
	absIntRuneStr := "123457"
	absFracRuneStr := "0"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

func TestNumStrDto_SetThisPrecision_01(t *testing.T) {

	nStr := "123456789"
	initialPrecision := uint(0)
	precision := uint(3)
	roundResult := true
	outPrecision := uint(3)
	expected := "123456789.000"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := "000"

	nsDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). " +
			"Err='%v' ", err.Error())
	}

	if initialPrecision != nsDto.GetPrecision() {
		t.Errorf("Error: Expected Initial Precision='%v'. Actual Initial Precision='%v' ",
			initialPrecision, nsDto.GetPrecision())
	}

  nsDto.SetThisPrecision(precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). " +
			"precision= '%v'. Error= %v", precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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


func TestNumStrDto_SetThisPrecision_02(t *testing.T) {

	nStr := "123456.789"
	initialPrecision := uint(3)
	precision := uint(2)
	roundResult := true
	outPrecision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr). " +
			"Err='%v' ", err.Error())
	}

	if initialPrecision != nsDto.GetPrecision() {
		t.Errorf("Error: Expected Initial Precision='%v'. Actual Initial Precision='%v' ",
			initialPrecision, nsDto.GetPrecision())
	}

  nsDto.SetThisPrecision(precision, roundResult)

	if err != nil {
		t.Errorf("Received error from nsu.SetThisPrecision(precision, roundResult). " +
			"precision= '%v'. Error= %v", precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.GetPrecision() {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.GetPrecision())
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

