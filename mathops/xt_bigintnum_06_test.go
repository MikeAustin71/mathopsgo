package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntNum_NumStrDto_01(t *testing.T) {

  ePrefix := "TestBigIntNum_NumStrDto_01"

  originalNumberStr := "123.456"

  expectedNumberStr := originalNumberStr

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedPrecisionUint := uint(3)

  expectedBigIntNumStr := "123456"

  expectedAbsBigIntNumStr := "123456"

  expectedScaleFactor := big.NewInt(1000)

  expectedSignVal := 1

  expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNum, isOk := big.NewInt(0).\n"+
      "   SetString(expectedBigIntNumStr, 10)\n"+
      "expectedBigIntNumStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedBigIntNumStr)
    return
  }

  expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
      "   SetString(expectedAbsBigIntNumStr, 10)\n"+
      "expectedAbsBigIntNumStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedAbsBigIntNumStr)
    return
  }

  nDto, err := new(NumStrDto).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum, err := new(BigIntNum).NewNumStrDto(nDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrDto(nDto)\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, bINumNumberStr)

    return
  }

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINumPrecisionUint)

    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

    return
  }

  bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  bINumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != bINumSeps \n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_NumStrDto_02(t *testing.T) {

  ePrefix := "TestBigIntNum_NumStrDto_02"

  originalNumberStr := "-123.456"

  expectedNumberStr := originalNumberStr

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedPrecisionUint := uint(3)

  expectedBigIntNumStr := "-123456"

  expectedAbsBigIntNumStr := "123456"

  expectedScaleFactor := big.NewInt(1000)

  expectedSignVal := -1

  expectedBigIntNum, isOk := big.NewInt(0).SetString(expectedBigIntNumStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNum, isOk := big.NewInt(0).\n"+
      "   SetString(expectedBigIntNumStr, 10)\n"+
      "expectedBigIntNumStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedBigIntNumStr)
    return
  }

  expectedAbsBigIntNum, isOk := big.NewInt(0).SetString(expectedAbsBigIntNumStr, 10)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedAbsBigIntNum, isOk := big.NewInt(0).\n"+
      "   SetString(expectedAbsBigIntNumStr, 10)\n"+
      "expectedAbsBigIntNumStr= '%v'\n"+
      "Error= 'isOk == false'\n\n",
      ePrefix, expectedAbsBigIntNumStr)
    return
  }

  nDto, err := new(NumStrDto).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).\n"+
      "  NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = nDto.IsValid("Validating nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = nDto.IsValid('Validating nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nDtoNumberStr, err := nDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDtoNumberStr, err := nDto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINum, err := new(BigIntNum).NewNumStrDto(nDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(BigIntNum).\n"+
      "  NewNumStrDto(nDto)\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, nDtoNumberStr, err.Error())
    return
  }

  err = bINum.IsValid("Validating bINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.IsValid('Validating bINum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bINumNumberStr, err := bINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumNumberStr, err := bINum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != bINumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != bINumNumberStr \n"+
      "Expected bINumNumberStr = '%v'\n"+
      "  Actual bINumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, bINumNumberStr)

    return
  }

  bINumPrecisionUint, err := bINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumPrecisionUint, err := bINum.GetPrecisionUint()\n"+
      "bINum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedPrecisionUint != bINumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/bINum Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != bINumPrecisionUint\n"+
      "Expected bINumPrecisionUint = '%v'\n"+
      "  Actual bINumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, bINumPrecisionUint)

    return
  }

  bINumBigInt, err := bINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumBigInt, err := bINum.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedBigIntNum.Cmp(bINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected/bINum Big Int Values ARE NOT EQUAL!\n"+
      "Because expectedBigIntNum.Cmp(bINumBigInt) != 0\n"+
      "Expected bINumBigInt = '%v'\n"+
      "  Actual bINumBigInt = '%v'\n\n",
      ePrefix, expectedBigIntNum.Text(10), bINumBigInt.Text(10))

    return
  }

  bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsValue, err := bINum.GetAbsoluteBigIntNumValue()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  bINumAbsBigInt, err := bINumAbsValue.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumAbsBigInt, err := bINumAbsValue.GetBigInt()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum Absolute Big Ints ARE NOT EQUAL!\n"+
      "Because expectedAbsBigIntNum.Cmp(bINumAbsBigInt) != 0\n"+
      "Expected bINumAbsBigInt = '%v'\n"+
      "  Actual bINumAbsBigInt = '%v'\n\n",
      ePrefix, expectedAbsBigIntNum.Text(10), bINumAbsBigInt.Text(10))

    return
  }

  bINumScaleFactor, err := bINum.GetScaleFactor()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumScaleFactor, err := bINum.GetScaleFactor()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedScaleFactor.Cmp(bINumScaleFactor) != 0 {
    t.Errorf("%v\n"+
      "Error: expected & bINum Scale Factors ARE NOT EQUAL!\n"+
      "Because expectedScaleFactor.Cmp(bINumScaleFactor) != 0\n"+
      "Expected bINumScaleFactor = '%v'\n"+
      "  Actual bINumScaleFactor = '%v'\n\n",
      ePrefix, expectedScaleFactor.Text(10), bINumScaleFactor.Text(10))

    return
  }

  bINumSignValue, err := bINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSignValue, err := bINum.GetSign()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if expectedSignVal != bINumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & bINum Sign Values ARE NOT EQUAL!\n"+
      "Because  expectedSignVal != bINumSignValue\n"+
      "Expected bINumSignValue = '%v'\n"+
      "  Actual bINumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, bINumSignValue)

    return
  }

  bINumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bINumSeps, err := bINum.GetNumericSeparatorsDto()\n"+
      "bINum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, bINumNumberStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(bINumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values NOT Equal\n"+
      "Because expectedNumSeps != bINumSeps \n"+
      "Expected bINumSeps = '%v'\n"+
      "  Actual bINumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), bINumSeps.String())

    return
  }

  return
}

func TestBigIntNum_RoundToDecPlace_01(t *testing.T) {

  nStr := "-123.567"
  expectedNumStr := "-123.57"
  roundToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_02(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.57"
  roundToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_03(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.567"
  roundToDec := uint(3)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_04(t *testing.T) {

  nStr := "123.567"
  expectedNumStr := "123.5670"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_05(t *testing.T) {

  nStr := "-123.567"
  expectedNumStr := "-123.5670"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_06(t *testing.T) {

  nStr := "0.000"
  expectedNumStr := "0.00"
  roundToDec := uint(2)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_07(t *testing.T) {

  nStr := "654.123456"
  expectedNumStr := "654.123"
  roundToDec := uint(3)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_08(t *testing.T) {

  nStr := "654.123456"
  expectedNumStr := "654.1235"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_09(t *testing.T) {

  nStr := "-654.123456"
  expectedNumStr := "-654.123"
  roundToDec := uint(3)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_10(t *testing.T) {

  nStr := "-654.123456"
  expectedNumStr := "-654.1235"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_11(t *testing.T) {

  nStr := "654"
  expectedNumStr := "654.0000"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_12(t *testing.T) {

  nStr := "654.123"
  expectedNumStr := "654.123000000"
  roundToDec := uint(9)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_13(t *testing.T) {

  nStr := "-654.123"
  expectedNumStr := "-654.123000000"
  roundToDec := uint(9)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_14(t *testing.T) {

  nStr := "-654"
  expectedNumStr := "-654.0000"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_RoundToDecPlace_15(t *testing.T) {

  nStr := "0"
  expectedNumStr := "0.0000"
  roundToDec := uint(4)

  bINum1, err := BigIntNum{}.NewNumStr(nStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(nStr). "+
      " nStr='%v'  Error='%v'",
      nStr, err.Error())
  }

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntNum_SetNumStr_01(t *testing.T) {

  origNumStr := "57.64"
  numStr := "89765.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }

}

func TestBigIntNum_SetNumStr_02(t *testing.T) {

  origNumStr := "257.1"
  numStr := "-89765.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }

}

func TestBigIntNum_SetNumStr_03(t *testing.T) {

  origNumStr := "0.000005"
  numStr := "0.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_04(t *testing.T) {

  origNumStr := "97.8"
  numStr := ".123456789012"
  expectedNumStr := "0.123456789012"
  expectedPrecision := uint(12)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_05(t *testing.T) {

  origNumStr := "87"
  numStr := "-.123456789012"
  expectedNumStr := "-0.123456789012"
  expectedPrecision := uint(12)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_06(t *testing.T) {
  origNumStr := "97.9821"
  numStr := "10"
  expectedNumStr := "10"
  expectedPrecision := uint(0)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_07(t *testing.T) {

  origNumStr := "9845.61"
  numStr := "-52"
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_08(t *testing.T) {

  origNumStr := "97"
  numStr := "-00052.1234"
  expectedNumStr := "-52.1234"
  expectedPrecision := uint(4)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }

}

func TestBigIntNum_SetNumStr_09(t *testing.T) {

  origNumStr := "87.123456"
  numStr := "(00052.1234)"
  expectedNumStr := "-52.1234"
  expectedPrecision := uint(4)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_10(t *testing.T) {

  origNumStr := "22.414141414"
  numStr := "(00052.1234"
  expectedNumStr := "52.1234"
  expectedPrecision := uint(4)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_11(t *testing.T) {
  origNumStr := "98.123456"
  numStr := "+00052.1234"
  expectedNumStr := "52.1234"
  expectedPrecision := uint(4)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_12(t *testing.T) {

  origNumStr := "97782345646"
  numStr := "-00052.1234567890123456"
  expectedNumStr := "-52.1234567890123456"
  expectedPrecision := uint(16)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_13(t *testing.T) {
  origNumStr := "52"
  numStr := "52 . 123 4567 8901 23456"
  expectedNumStr := "52.1234567890123456"
  expectedPrecision := uint(16)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_14(t *testing.T) {

  origNumStr := "987.123456"
  numStr := "5 2"
  expectedNumStr := "52"
  expectedPrecision := uint(0)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_15(t *testing.T) {

  origNumStr := "-8794521.12345"
  numStr := "    (52)     "
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_16(t *testing.T) {

  origNumStr := "98"
  numStr := "    (52)    1234567 "
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  if expectedNumStr != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      expectedNumStr, bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())

  }
}

func TestBigIntNum_SetNumStr_17(t *testing.T) {

  origNumStr := "98123456789"
  expectedPrecision := uint(1024)

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(EulersNum50kStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  bigINum.RoundToDecPlace(1024)

  fdEulers1k := GetEulersNum1k()

  if fdEulers1k.GetNumStr() != bigINum.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
      fdEulers1k.GetNumStr(), bigINum.GetNumStr())
  }

  if expectedPrecision != bigINum.GetPrecisionUint() {
    t.Errorf("Error: Expected precision='%v'. Instead, precision='%v'",
      expectedPrecision, bigINum.GetPrecisionInt())
  }
}

func TestBigIntNum_SetNumStr_18(t *testing.T) {

  origNumStr := "98"
  numStr := "abcdefghijklmnop"

  bigINum, err := BigIntNum{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by %v", err.Error())
  }

  err = bigINum.SetNumStr(numStr)

  if err == nil {
    t.Error("Expected an Error from INVALID Number String. NO ERROR RETURNED!")
  }
}

func TestBigIntNum_ShiftPrecisionLeft_01(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "123.456789"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_02(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "1234.56789"
  shiftPlacesLeft := uint(2)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_03(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "0.123456789"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_04(t *testing.T) {
  basicNumStr := "123456789"
  expectedResult := "123.456789"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_05(t *testing.T) {
  basicNumStr := "123"
  expectedResult := "0.00123"
  shiftPlacesLeft := uint(5)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_06(t *testing.T) {
  basicNumStr := "0"
  expectedResult := "0"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_07(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "123456.789"
  shiftPlacesLeft := uint(0)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_08(t *testing.T) {
  basicNumStr := "-123456.789"
  expectedResult := "-123456.789"
  shiftPlacesLeft := uint(0)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_09(t *testing.T) {
  basicNumStr := "-123456.789"
  expectedResult := "-123.456789"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionLeft_10(t *testing.T) {
  basicNumStr := "-123456789"
  expectedResult := "-123.456789"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_01(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "123456789"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_02(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "12345678.9"
  shiftPlacesLeft := uint(2)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_03(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "123456789000"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_04(t *testing.T) {
  basicNumStr := "123456789"
  expectedResult := "123456789000000"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_05(t *testing.T) {
  basicNumStr := "123"
  expectedResult := "12300000"
  shiftPlacesLeft := uint(5)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_06(t *testing.T) {
  basicNumStr := "0"
  expectedResult := "0"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_07(t *testing.T) {
  basicNumStr := "123456.789"
  expectedResult := "123456.789"
  shiftPlacesLeft := uint(0)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_08(t *testing.T) {
  basicNumStr := "-123456.789"
  expectedResult := "-123456.789"
  shiftPlacesLeft := uint(0)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_09(t *testing.T) {
  basicNumStr := "-123456.789"
  expectedResult := "-123456789"
  shiftPlacesLeft := uint(3)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}

func TestBigIntNum_ShiftPrecisionRight_10(t *testing.T) {
  basicNumStr := "-123456789"
  expectedResult := "-123456789000000"
  shiftPlacesLeft := uint(6)

  bIntNum, err := BigIntNum{}.NewNumStr(basicNumStr)

  if err != nil {
    t.Errorf("Error returned by BigIntNum{}.NewNumStr(basicNumStr). "+
      "basicNumStr='%v' Error='%v' ",
      basicNumStr, err.Error())
  }

  bIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  actualResult := bIntNum.GetNumStr()

  if expectedResult != actualResult {
    t.Errorf("Error: Expected result='%v'. Actual result='%v' ",
      expectedResult, actualResult)
  }

}
