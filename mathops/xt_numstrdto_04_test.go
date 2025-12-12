package mathops

import (
  "math/big"
  "testing"
)

func TestNumStrDto_GetSignedBigInt_01(t *testing.T) {

  ePrefix := "TestNumStrDto_GetSignedBigInt_01"

  inputNumberStr := "-123.456"

  originalBigIntNumStr := "-123456"

  expectedBigIntNum, isOk := big.NewInt(0).SetString(originalBigIntNumStr, 10)

  if !isOk {
    t.Errorf("%v\n"+
      "isOk='false' returned by:\n"+
      "expectedBigIntNum, isOk :=\n"+
      "  big.NewInt(0).SetString(originalBigIntNumStr, 10)\n"+
      "originalBigIntNumStr= '%v'\n\n",
      ePrefix, originalBigIntNumStr)
    return
  }

  numStrDtoResult, err := new(NumStrDto).NewPtr().ParseNumStr(inputNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResult, err :=\n"+
      "  new(NumStrDto).NewPtr().ParseNumStr(inputNumberStr)\n"+
      "inputNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, inputNumberStr, err.Error())
    return
  }

  err = numStrDtoResult.IsValid("Validating numStrDtoResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = numStrDtoResult.IsValid('Validating numStrDtoResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultNumberStr, err := numStrDtoResult.GetNumStr()\n"+
      "numStrDtoResult set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  numStrDtoResultBigIntNum, err := numStrDtoResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "numStrDtoResultBigIntNum, err := numStrDtoResult.GetBigInt()\n"+
      "numStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, numStrDtoResultNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(numStrDtoResultBigIntNum) != 0 {
    t.Errorf("%v\n"+
      "Error: expectedBigIntNum And numStrDtoResultBigIntNum ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(numStrDtoResultBigIntNum) != 0\n"+
      "Expected numStrDtoResultBigIntNum = '%v'\n"+
      "  Actual numStrDtoResultBigIntNum = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), numStrDtoResultBigIntNum.Text(10))

    return
  }

  return
}

func TestNumStrDto_GetThouStr_01(t *testing.T) {

  nStr := "123456.97"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123,456.97"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_02(t *testing.T) {

  nStr := "123.45"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123.45"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_03(t *testing.T) {

  nStr := "12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_04(t *testing.T) {

  nStr := "12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_05(t *testing.T) {

  nStr := "12345.1234"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345.1234"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_06(t *testing.T) {

  nStr := "-12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-12,345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_07(t *testing.T) {

  nStr := "-12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-12,345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_08(t *testing.T) {

  nStr := "-123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_09(t *testing.T) {

  nStr := "-0.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-0.123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouStr_10(t *testing.T) {

  nStr := "-1234567890.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "-1,234,567,890.123"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouStr()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_01(t *testing.T) {

  nStr := "123456.97"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123,456.97"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_02(t *testing.T) {

  nStr := "123.45"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "123.45"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_03(t *testing.T) {

  nStr := "12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345.29"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_04(t *testing.T) {

  nStr := "12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_05(t *testing.T) {

  nStr := "12345.1234"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "12,345.1234"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_06(t *testing.T) {

  nStr := "1234567890.25"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "1,234,567,890.25"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_07(t *testing.T) {

  nStr := "-12345.29"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(12,345.29)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_08(t *testing.T) {

  nStr := "-12345"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(12,345)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_09(t *testing.T) {

  nStr := "-123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_10(t *testing.T) {

  nStr := "-0.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(0.123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}

func TestNumStrDto_GetThouParen_11(t *testing.T) {

  nStr := "-1234567890.123"
  currencySymbol := '$'
  decimalSeparator := '.'
  thousandsSeparator := ','
  expectedStr := "(1,234,567,890.123)"

  nDto, err := NumStrDto{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr) "+
      "nStr='%v' Error='%v'", nStr, err.Error())
  }

  nDto.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

  if currencySymbol != nDto.GetCurrencySymbol() {
    t.Errorf("Expected Currency Symbol='%v'.  Instead, Currency Symbol='%v' .",
      currencySymbol, nDto.GetCurrencySymbol())
  }

  if decimalSeparator != nDto.GetDecimalSeparator() {
    t.Errorf("Expected Decimal Separator='%v'.  Instead, Decimal Separator='%v' .",
      decimalSeparator, nDto.GetDecimalSeparator())

  }

  if thousandsSeparator != nDto.GetThousandsSeparator() {
    t.Errorf("Expected Thousands Separator='%v'.  Instead, Thousands Separator='%v' .",
      thousandsSeparator, nDto.GetThousandsSeparator())

  }

  actualStr := nDto.GetThouParen()

  if expectedStr != actualStr {
    t.Errorf("Expected Currency Str='%v'. Instead, Currency Str='%v'",
      expectedStr, actualStr)
  }

}
