package mathops

import (
	"math/big"
	"testing"
)

func TestNumStrDto_AddNumStrs_01(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_01() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_02(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "-9211.40"
	nStr3 := "377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_02() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_03(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "9211.40"
	nStr3 := "18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_03() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_04(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "-9211.40"
	nStr3 := "-18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_04() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_05(t *testing.T) {
	nStr1 := "2"
	nStr2 := "3"
	nStr3 := "5"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_05() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_06(t *testing.T) {
	nStr1 := "2"
	nStr2 := "0.0"
	nStr3 := "2.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_06() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_07(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0.0"
	nStr3 := "0.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_07() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_08(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_08() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_09(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_09() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_10(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())
	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_11(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.GetNumStr()
	expected := nExpected.GetNumStr()

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsIntRunes())
	expected = string(nExpected.GetAbsIntRunes())

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.GetAbsFracRunes())
	expected = string(nExpected.GetAbsFracRunes())

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.GetSign() != nResult.GetSign() {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.GetSign(), nResult.GetSign())
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue() != nResult.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue(), nResult.IsFractionalValue())
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())
	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_CompareAbsoluteVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_02(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_04(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_05(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_06(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_07(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_08(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "1567.21"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_02(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_04(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_05(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_06(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_07(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_08(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_FormatForMathOps_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_FormatForMathOps_02(t *testing.T) {
	nStr1 := "-9211.40"
	nStr2 := "-12567.218956"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_FormatForMathOps_03(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_FormatForMathOps_04(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_FormatForMathOps_05(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-67.521"
	nStr4 := "06.000"
	expectedCompare := 1
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_FormatForMathOps_06(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "-67.521"
	nStr4 := "67.521"
	expectedCompare := 0
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.GetNumStr() != n1OutDto.GetNumStr() {
		t.Errorf("Expected n1OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut1.GetNumStr(), n1OutDto.GetNumStr())
	}

	if nOut2.GetNumStr() != n2OutDto.GetNumStr() {
		t.Errorf("Expected n2OutDto.GetNumStr()= '%v'. Instead got '%v'", nOut2.GetNumStr(), n2OutDto.GetNumStr())
	}

	if nOut1.GetSign() != n1OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut1.GetSign(), n1OutDto.GetSign())
	}

	if nOut2.GetSign() != n2OutDto.GetSign() {
		t.Errorf("Expected n1OutDto.GetSign()= '%v'. Instead got '%v'", nOut2.GetSign(), n2OutDto.GetSign())
	}

	if nOut1.GetPrecision() != n1OutDto.GetPrecision() {
		t.Errorf("Expected n1OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut1.GetPrecision(), n1OutDto.GetPrecision())
	}

	if nOut2.GetPrecision() != n2OutDto.GetPrecision() {
		t.Errorf("Expected n2OutDto.GetPrecision()= '%v'. Instead got '%v'", nOut2.GetPrecision(), n2OutDto.GetPrecision())
	}

}

func TestNumStrDto_CopyIn_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid() {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid())
	}

}

func TestNumStrDto_CopyIn_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid() {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid())
	}

}

func TestNumStrDto_CopyOut_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid() {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid())
	}

}

func TestNumStrDto_CopyOut_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.GetNumStr()

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsIntRunes())

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.GetAbsFracRunes())

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.GetSign() != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.GetSign())
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue() {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue())
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid() {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid())
	}

}

func TestNumStrDto_GetAbsoluteBigInt_01(t *testing.T) {

	nStr := "-123.456"
	absNumStr := "123456"
	expected, isOk := big.NewInt(0).SetString(absNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(absNumStr,10) Failed!. absNumStr= '%v'", absNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	absBigInt, err := n1.GetAbsoluteBigInt()

	if absBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected absBigInt= %v . Instead got, %v", expected.String(), absBigInt.String())
	}

}

func TestNumStrDto_GetSignedBigInt_01(t *testing.T) {

	nStr := "-123.456"
	signedNumStr := "-123456"
	expected, isOk := big.NewInt(0).SetString(signedNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(signedNumStr,10) Failed!. signedNumStr= '%v'", signedNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	signedBigInt, err := n1.GetSignedBigInt()

	if signedBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected signedBigInt= %v . Instead got, %v", expected.String(), signedBigInt.String())
	}

}
