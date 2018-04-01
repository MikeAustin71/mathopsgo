package mathops

import (
	"testing"
	"math/big"
)

func TestNumStrDto_MultiplyNumStrs_01(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_02(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "-1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_03(t *testing.T) {
	nStr1 := "-35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_04(t *testing.T) {
	nStr1 := "57"
	nStr2 := "123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_05(t *testing.T) {
	nStr1 := "57"
	nStr2 := "-123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_06(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_07(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "-123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_08(t *testing.T) {
	nStr1 := "0"
	nStr2 := "123"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_09(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_10(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_11(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0.123"
	nStr3 := "7.011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_12(t *testing.T) {
	nStr1 := "62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_13(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "-194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_14(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "-3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

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

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.GetPrecision() != nResult.GetPrecision() {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.GetPrecision(), nResult.GetPrecision())

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_NewNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_NewNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_NewNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_NewNumStr_04(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_NewNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
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

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_04(t *testing.T) {
	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

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

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
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

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseSignedBigInt_01(t *testing.T) {

	signedAbsNumStr := "-123456789"
	absAllNumStr := "123456789"
	nStr := "-123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := -1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("BigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

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

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseSignedBigInt_02(t *testing.T) {

	signedAbsNumStr := "123456789"
	absAllNumStr := "123456789"
	nStr := "123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := 1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("BigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.GetAbsAllNumRunes())

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

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

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.GetPrecision() != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.GetPrecision())

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

