package mathops

import (
	"testing"
	"math/big"
)

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

func TestNumStrDto_GetCurrencyStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123,456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyStr_06(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-$12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}


func TestNumStrDto_GetCurrencyStr_07(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-$12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123,456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}


func TestNumStrDto_GetCurrencyParen_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "$12,345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_05(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "($12,345.29)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetCurrencyParen_06(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "($12,345)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetCurrencyParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}


func TestNumStrDto_GetNumStr_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-12345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_09(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_10(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-0.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumStr_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "-1234567890.123"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumStr()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_01(t *testing.T) {

	nStr := "123456.97"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123456.97"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_02(t *testing.T) {

	nStr := "123.45"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "123.45"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_03(t *testing.T) {

	nStr := "12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.29"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_04(t *testing.T) {

	nStr := "12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_05(t *testing.T) {

	nStr := "12345.1234"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "12345.1234"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}


func TestNumStrDto_GetNumParen_06(t *testing.T) {

	nStr := "1234567890.25"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "1234567890.25"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_07(t *testing.T) {

	nStr := "-12345.29"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345.29)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_08(t *testing.T) {

	nStr := "-12345"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(12345)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_09(t *testing.T) {

	nStr := "-123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_10(t *testing.T) {

	nStr := "-0.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(0.123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}

func TestNumStrDto_GetNumParen_11(t *testing.T) {

	nStr := "-1234567890.123"
	currencySymbol := '$'
	decimalSeparator := '.'
	thousandsSeparator := ','
	expectedStr := "(1234567890.123)"

	nDto, err := NumStrDto{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) " +
			"nStr='%v' Error='%v'", nStr, err.Error())
	}

	nDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)


	if currencySymbol != nDto.GetCurrencySymbol() {
		t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
			currencySymbol, nDto.GetCurrencySymbol())
	}

	if decimalSeparator != nDto.GetDecimalSeparator() {
		t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
			decimalSeparator, nDto.GetDecimalSeparator() )

	}

	if thousandsSeparator != nDto.GetThousandsSeparator() {
		t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
			thousandsSeparator, nDto.GetThousandsSeparator() )

	}

	actualStr := nDto.GetNumParen()

	if expectedStr != actualStr {
		t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
			expectedStr, actualStr)
	}

}
