package mathops

import (
  "math/big"
  "strconv"
  "testing"
)

func TestDecimal_Pow_01(t *testing.T) {

  ePrefix := "TestDecimal_Pow_01"

  originalNumberStr := "2.325"

  exponentNumberStr := "3.8"

  //                                1         2         3
  //                     0.123456789012345678901234567890
  expectedNumberStr := "24.683528241199594112131040124684"

  originalMaxPrecision := uint(30)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_Pow_02(t *testing.T) {

  ePrefix := "TestDecimal_Pow_02"

  originalNumberStr := "-2.19"

  exponentNumberStr := "8.256"

  //                                 1         2        2
  //                      0.12345678901234567890123456789
  expectedNumberStr := "646.70558582734148834284281249154"

  originalMaxPrecision := uint(29)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_Pow_03(t *testing.T) {

  ePrefix := "TestDecimal_Pow_03"

  originalNumberStr := "2.19"

  exponentNumberStr := "-8.256"

  //                               1         2         3
  //                    0.12345678901234567890123456789012
  expectedNumberStr := "0.00154629868972089198924071380209"

  originalMaxPrecision := uint(32)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_Pow_04(t *testing.T) {

  ePrefix := "TestDecimal_Pow_04"

  originalNumberStr := "2"

  exponentNumberStr := "36"

  //                                         1         2         3
  //                              0.12345678901234567890123456789012
  expectedNumberStr := "68719476736"

  originalMaxPrecision := uint(1)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_Pow_05(t *testing.T) {

  ePrefix := "TestDecimal_Pow_05"

  originalNumberStr := "87"

  exponentNumberStr := "0"

  //                            1         2         3
  //                 0.12345678901234567890123456789012
  expectedNumberStr := "1"

  originalMaxPrecision := uint(1)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_Pow_06(t *testing.T) {

  ePrefix := "TestDecimal_Pow_05"

  originalNumberStr := "0"

  exponentNumberStr := "87"

  //                               1         2         3
  //                    0.12345678901234567890123456789012
  expectedNumberStr := "0"

  originalMaxPrecision := uint(1)

  expectedPrecisionUint := originalMaxPrecision

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponent, err := new(Decimal).NewNumStr(exponentNumberStr)\n"+
      "exponentNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, exponentNumberStr, err.Error())
    return
  }

  err = decNumExponent.IsValid("Validating decNumExponent")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumExponent.IsValid('Validating decNumExponent')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumExponentNumberStr, err := decNumExponent.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumExponentNumberStr, err := decNumExponent.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResult, err := decNumOriginal.Pow(decNumExponent, originalMaxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.Pow(\n"+
      "  decNumExponent, originalMaxPrecision)\n"+
      "decNumOriginal= '%v'\n"+
      "decNumExponent= '%v'\n"+
      "originalMaxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      decNumExponentNumberStr,
      originalMaxPrecision,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_PowInt_01(t *testing.T) {

  ePrefix := "TestDecimal_PowInt_01"

  originalNumberStr := "2.125"

  exponentNumberInt := 5

  //                                1         2         3
  //                     0.12345678901234567890123456789012
  expectedNumberStr := "43.330596923828125"

  originalMaxPrecisionUint := uint(15)

  expectedPrecisionUint := originalMaxPrecisionUint

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumOriginalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumOriginalNumberStr \n"+
      "Expected decNumOriginalNumberStr = '%v'\n"+
      "  Actual decNumOriginalNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumOriginalNumberStr)

    return
  }

  decNumFinalResult, err := decNumOriginal.PowInt(exponentNumberInt, originalMaxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.PowInt(\n"+
      "  exponentNumberInt, originalMaxPrecisionUint)\n"+
      "decNumOriginal= '%v'\n"+
      "exponentNumberInt= '%v'\n"+
      "originalMaxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      exponentNumberInt,
      originalMaxPrecisionUint,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_PowInt_02(t *testing.T) {

  ePrefix := "TestDecimal_PowInt_02"

  originalNumberStr := "2.125"

  exponentNumberInt := -5

  //                               1         2         3
  //                    0.12345678901234567890123456789012
  expectedNumberStr := "0.02307838042845159759046157465153"

  originalMaxPrecisionUint := uint(32)

  expectedPrecisionUint := originalMaxPrecisionUint

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumOriginalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumOriginalNumberStr \n"+
      "Expected decNumOriginalNumberStr = '%v'\n"+
      "  Actual decNumOriginalNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumOriginalNumberStr)

    return
  }

  decNumFinalResult, err := decNumOriginal.PowInt(exponentNumberInt, originalMaxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.PowInt(\n"+
      "  exponentNumberInt, originalMaxPrecisionUint)\n"+
      "decNumOriginal= '%v'\n"+
      "exponentNumberInt= '%v'\n"+
      "originalMaxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      exponentNumberInt,
      originalMaxPrecisionUint,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_PowInt_03(t *testing.T) {

  ePrefix := "TestDecimal_PowInt_03"

  originalNumberStr := "2.125"

  exponentNumberInt := -5

  expectedNumberStr := "0.0230783804284515975904615746515318091892352539727592285702010836302529057503678187310412245740240038257373805953698154109885713843013768287933221444131345621425256205378428954465132756326869536861810731644102187755527493261645362878092652992519669234"

  originalMaxPrecisionUint := uint(250)

  expectedPrecisionUint := originalMaxPrecisionUint

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumOriginal.IsValid("Validating decNumOriginal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumOriginal.IsValid('Validating decNumOriginal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumOriginalNumberStr, err := decNumOriginal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumOriginalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumOriginalNumberStr \n"+
      "Expected decNumOriginalNumberStr = '%v'\n"+
      "  Actual decNumOriginalNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumOriginalNumberStr)

    return
  }

  decNumFinalResult, err := decNumOriginal.PowInt(exponentNumberInt, originalMaxPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResult, err := decNumOriginal.PowInt(\n"+
      "  exponentNumberInt, originalMaxPrecisionUint)\n"+
      "decNumOriginal= '%v'\n"+
      "exponentNumberInt= '%v'\n"+
      "originalMaxPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumOriginalNumberStr,
      exponentNumberInt,
      originalMaxPrecisionUint,
      err.Error())

    return
  }

  err = decNumFinalResult.IsValid("Validating decNumFinalResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinalResult.IsValid('Validating decNumFinalResult')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumberStr, err := decNumFinalResult.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalResultPrecisionUint, err := decNumFinalResult.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultPrecisionUint, err :=\n"+
      "  decNumFinalResult.GetPrecisionUint()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultSignValue, err := decNumFinalResult.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultSignValue, err := decNumFinalResult.GetSign()\n"+
      "decNumFinalResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalResultNumSeps, err := decNumFinalResult.GetNumericSeparatorsDto()\n"+
      "decNumFinalResult= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalResultNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalResultNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalResultNumberStr \n"+
      "Expected decNumFinalResultNumberStr = '%v'\n"+
      "  Actual decNumFinalResultNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalResultNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalResultPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalResultPrecisionUint\n"+
      "Expected decNumFinalResultPrecisionUint = '%v'\n"+
      "  Actual decNumFinalResultPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalResultPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalResultSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalResultSignValue\n"+
      "Expected decNumFinalResultSignValue = '%v'\n"+
      "  Actual decNumFinalResultSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalResultSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalResultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalResultNumSeps \n"+
      "Expected decNumFinalResultNumSeps = '%v'\n"+
      "  Actual decNumFinalResultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalResultNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetPrecisionRound_01(t *testing.T) {

  ePrefix := "TestDecimal_SetPrecisionRound_01"

  originalNumberStr := "2.0105500"

  expectedPrecisionUint := uint(4)

  //                               1         2         3
  //                    0.12345678901234567890123456789012
  expectedNumberStr := "2.0106"

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #1')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number #1 String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr #1 = '%v'\n"+
      "  Actual decNumFinalNumberStr #1 = '%v'\n\n",
      ePrefix, originalNumberStr, decNumFinalNumberStr)

    return
  }

  err = decNumFinal.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.SetPrecisionRound(expectedPrecisionUint)\n"+
      "decNumFinal= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, expectedPrecisionUint, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #2')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err = decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #2 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetPrecisionRound_02(t *testing.T) {

  ePrefix := "TestDecimal_SetPrecisionRound_02"

  originalNumberStr := "-2.0105500"

  expectedNumberStr := "-2.0106"

  expectedPrecisionUint := uint(4)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #1')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number #1 String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr #1 = '%v'\n"+
      "  Actual decNumFinalNumberStr #1 = '%v'\n\n",
      ePrefix, originalNumberStr, decNumFinalNumberStr)

    return
  }

  err = decNumFinal.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.SetPrecisionRound(expectedPrecisionUint)\n"+
      "decNumFinal= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, expectedPrecisionUint, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #2')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err = decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #2 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetPrecisionTrunc_01(t *testing.T) {

  ePrefix := "TestDecimal_SetPrecisionTrunc_01"

  originalNumberStr := "2.0105500"

  expectedNumberStr := "2.0105"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #1')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number #1 String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr #1 = '%v'\n"+
      "  Actual decNumFinalNumberStr #1 = '%v'\n\n",
      ePrefix, originalNumberStr, decNumFinalNumberStr)

    return
  }

  err = decNumFinal.SetPrecisionTrunc(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.SetPrecisionTrunc(expectedPrecisionUint)\n"+
      "decNumFinal= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, expectedPrecisionUint, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #2')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err = decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #2 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetPrecisionTrunc_02(t *testing.T) {

  ePrefix := "TestDecimal_SetPrecisionTrunc_02"

  originalNumberStr := "-2.0105500"

  expectedNumberStr := "-2.0105"

  expectedPrecisionUint := uint(4)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinal, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #1')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number #1 String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr #1 = '%v'\n"+
      "  Actual decNumFinalNumberStr #1 = '%v'\n\n",
      ePrefix, originalNumberStr, decNumFinalNumberStr)

    return
  }

  err = decNumFinal.SetPrecisionTrunc(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.SetPrecisionTrunc(expectedPrecisionUint)\n"+
      "decNumFinal= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, expectedPrecisionUint, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal #2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal #2')\n"+
      "decNumFinal #1 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err = decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "decNumFinal #2 \n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_01(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_01"

  expectedNumberStr := "1.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_02(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_02"

  expectedNumberStr := "-1.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_03(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_03"

  expectedNumberStr := "0.00"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_04(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_04"

  originalNumberStr := "-0.00"

  expectedNumberStr := "0.00"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(originalNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_05(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_05"

  originalNumberStr := "92"

  expectedNumberStr := "92"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(originalNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStr_06(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStr_06"

  originalNumberStr := "124443,912456"

  expectedNumberStr := "124443,912456"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  decNumFinal := new(Decimal).New()

  err := decNumFinal.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())

    return
  }

  err = decNumFinal.SetNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumFinal.SetNumStr(originalNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNumFinal.IsValid("Validating decNumFinal")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumFinal.IsValid('Validating decNumFinal')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalNumberStr, err := decNumFinal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumberStr, err := decNumFinal.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumFinalPrecisionUint, err := decNumFinal.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalPrecisionUint, err :=\n"+
      "  decNumFinal.GetPrecisionUint()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalSignValue, err := decNumFinal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalSignValue, err := decNumFinal.GetSign()\n"+
      "decNumFinal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumFinalNumSeps, err := decNumFinal.GetNumericSeparatorsDto()\n"+
      "decNumFinal= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumFinalNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumFinalNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumFinalNumberStr \n"+
      "Expected decNumFinalNumberStr = '%v'\n"+
      "  Actual decNumFinalNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumFinalNumberStr)

    return
  }

  if expectedPrecisionUint != decNumFinalPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumFinalPrecisionUint\n"+
      "Expected decNumFinalPrecisionUint = '%v'\n"+
      "  Actual decNumFinalPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumFinalPrecisionUint)

    return
  }

  if expectedSignVal != decNumFinalSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumFinalSignValue\n"+
      "Expected decNumFinalSignValue = '%v'\n"+
      "  Actual decNumFinalSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumFinalSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumFinalNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumFinalNumSeps \n"+
      "Expected decNumFinalNumSeps = '%v'\n"+
      "  Actual decNumFinalNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumFinalNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStrDto_01(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStrDto_01"

  originalNumberStr := "-999999.99999"

  expectedNumberStr := "1.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetNumStrDto(nDto)\n"+
      "decNum= '%v'\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumNumberStr,
      nDtoNumberStr,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum-nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-nDto')\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStrDto_02(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStrDto_02"

  originalNumberStr := "-999999.99999"

  expectedNumberStr := "-1555666.35"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetNumStrDto(nDto)\n"+
      "decNum= '%v'\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumNumberStr,
      nDtoNumberStr,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum-nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-nDto')\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetNumStrDto_03(t *testing.T) {

  ePrefix := "TestDecimal_SetNumStrDto_03"

  originalNumberStr := "-999999.99999"

  expectedNumberStr := "1555777.123456"

  expectedPrecisionUint := uint(6)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nDto, err := new(NumStrDto).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
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

  if expectedNumberStr != nDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != nDtoNumberStr \n"+
      "Expected nDtoNumberStr = '%v'\n"+
      "  Actual nDtoNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, nDtoNumberStr)

    return
  }

  decNum, err := new(Decimal).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-original")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-original')\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to originalNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, decNumNumberStr)

    return
  }

  err = decNum.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetNumStrDto(nDto)\n"+
      "decNum= '%v'\n"+
      "nDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      decNumNumberStr,
      nDtoNumberStr,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum-nDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-nDto')\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "decNum set to nDto\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetFloat_01(t *testing.T) {

  ePrefix := "TestDecimal_SetFloat_01"

  originalFloat32Value := float32(92.25)

  originalFloat64Value := float64(originalFloat32Value)

  expectedNumberStr := "92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetFloat32(originalFloat32Value)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloat32(originalFloat32Value)\n"+
      "originalFloat32Value= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Value, 'f', int(expectedPrecisionUint), 64),
      err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetFloat_02(t *testing.T) {

  ePrefix := "TestDecimal_SetFloat_02"

  originalFloat32Value := float32(-92.25)

  originalFloat64Value := float64(originalFloat32Value)

  expectedNumberStr := "-92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetFloat32(originalFloat32Value)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloat32(originalFloat32Value)\n"+
      "originalFloat32Value= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalFloat64Value, 'f', int(expectedPrecisionUint), 64),
      err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-nDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetFloat64_01(t *testing.T) {

  ePrefix := "TestDecimal_SetFloat64_01"

  var originalFloat64Value = 92.256 // Float64

  expectedNumberStr := "92.256"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetFloat64(originalFloat64Value)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloat64(originalFloat64Value)\n"+
      "originalFloat64Value= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalFloat64Value, err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetFloat64_02(t *testing.T) {

  ePrefix := "TestDecimal_SetFloat64_02"

  originalFloat64Value := -92.25 // Float64

  expectedNumberStr := "-92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetFloat64(originalFloat64Value)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloat64(originalFloat64Value)\n"+
      "originalFloat64Value= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalFloat64Value, err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}
func TestDecimal_SetFloatBig_01(t *testing.T) {

  ePrefix := "TestDecimal_SetFloatBig_01"

  expectedNumberStr := "92.256"

  expectedPrecisionUint := uint(3)

  expectedPrecisionInt := int(expectedPrecisionUint)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  bigFloatValue, isOk := big.NewFloat(0.0).SetString(expectedNumberStr)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigFloatValue, isOk := big.NewFloat(0.0).SetString(expectedNumberStr)\n"+
      "isOk == 'false'\n"+
      "expectedNumberStr= '%v'\n\n",
      ePrefix, expectedNumberStr)
    return

  }

  decNum := new(Decimal).New()

  err := decNum.SetFloatBig(bigFloatValue)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloatBig(bigFloatValue)\n"+
      "bigFloatValue= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigFloatValue.Text('f', expectedPrecisionInt), err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetFloatBig_02(t *testing.T) {

  ePrefix := "TestDecimal_SetFloatBig_02"

  expectedNumberStr := "-92.25"

  expectedPrecisionUint := uint(2)

  expectedPrecisionInt := int(expectedPrecisionUint)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  bigFloatValue, isOk := big.NewFloat(0.0).SetString(expectedNumberStr)

  if !isOk {

    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bigFloatValue, isOk := big.NewFloat(0.0).SetString(expectedNumberStr)\n"+
      "isOk == 'false'\n"+
      "expectedNumberStr= '%v'\n\n",
      ePrefix, expectedNumberStr)
    return

  }

  decNum := new(Decimal).New()

  err := decNum.SetFloatBig(bigFloatValue)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetFloatBig(bigFloatValue)\n"+
      "bigFloatValue= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, bigFloatValue.Text('f', expectedPrecisionInt), err.Error())
    return
  }

  err = decNum.SetPrecisionRound(expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetPrecisionRound(expectedPrecisionUint)\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedPrecisionUint, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetIntFracStrings_01(t *testing.T) {

  ePrefix := "TestDecimal_SetIntFracStrings_01"

  originalIntNumberStr := "123"

  originalFractionNumberStr := "456"

  expectedNumberStr := "123.456"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetIntFracStrings(originalIntNumberStr, originalFractionNumberStr, expectedSignVal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetIntFracStrings(\n"+
      "  originalIntNumberStr, originalFractionNumberStr, expectedSignVal)\n"+
      "originalIntNumberStr= '%v'\n"+
      "originalFractionNumberStr= '%v'\n"+
      "expectedSignVal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalIntNumberStr,
      originalFractionNumberStr,
      expectedSignVal,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetIntFracStrings_02(t *testing.T) {

  ePrefix := "TestDecimal_SetIntFracStrings_02"

  originalIntNumberStr := "123"

  originalFractionNumberStr := "0456"

  expectedNumberStr := "-123.0456"

  expectedPrecisionUint := uint(4)

  expectedSignVal := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetIntFracStrings(originalIntNumberStr, originalFractionNumberStr, expectedSignVal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetIntFracStrings(\n"+
      "  originalIntNumberStr, originalFractionNumberStr, expectedSignVal)\n"+
      "originalIntNumberStr= '%v'\n"+
      "originalFractionNumberStr= '%v'\n"+
      "expectedSignVal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalIntNumberStr,
      originalFractionNumberStr,
      expectedSignVal,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetIntFracStrings_03(t *testing.T) {

  ePrefix := "TestDecimal_SetIntFracStrings_03"

  originalIntNumberStr := "-0"

  originalFractionNumberStr := "04#5 6"

  expectedNumberStr := "0.0456"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetIntFracStrings(originalIntNumberStr, originalFractionNumberStr, expectedSignVal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetIntFracStrings(\n"+
      "  originalIntNumberStr, originalFractionNumberStr, expectedSignVal)\n"+
      "originalIntNumberStr= '%v'\n"+
      "originalFractionNumberStr= '%v'\n"+
      "expectedSignVal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalIntNumberStr,
      originalFractionNumberStr,
      expectedSignVal,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetIntFracStrings_04(t *testing.T) {

  ePrefix := "TestDecimal_SetIntFracStrings_04"

  originalIntNumberStr := "-0"

  originalFractionNumberStr := ".04#5 6"

  expectedNumberStr := "0.0456"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetIntFracStrings(originalIntNumberStr, originalFractionNumberStr, expectedSignVal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetIntFracStrings(\n"+
      "  originalIntNumberStr, originalFractionNumberStr, expectedSignVal)\n"+
      "originalIntNumberStr= '%v'\n"+
      "originalFractionNumberStr= '%v'\n"+
      "expectedSignVal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalIntNumberStr,
      originalFractionNumberStr,
      expectedSignVal,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetIntFracStrings_05(t *testing.T) {

  ePrefix := "TestDecimal_SetIntFracStrings_05"

  originalIntNumberStr := "0"

  originalFractionNumberStr := ",0456"

  expectedNumberStr := "0,0456"

  expectedPrecisionUint := uint(4)

  expectedSignVal := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  decNum := new(Decimal).New()

  err := decNum.SetIntFracStrings(originalIntNumberStr, originalFractionNumberStr, expectedSignVal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetIntFracStrings(\n"+
      "  originalIntNumberStr, originalFractionNumberStr, expectedSignVal)\n"+
      "originalIntNumberStr= '%v'\n"+
      "originalFractionNumberStr= '%v'\n"+
      "expectedSignVal= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalIntNumberStr,
      originalFractionNumberStr,
      expectedSignVal,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint_01(t *testing.T) {

  ePrefix := "TestDecimal_SetUint_01"

  originalUintValue := uint(9225)

  expectedNumberStr := "92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint(originalUintValue, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint(originalUintValue, expectedPrecisionUint)\n"+
      "originalUintValue= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUintValue,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint_02(t *testing.T) {

  ePrefix := "TestDecimal_SetUint_02"

  originalUintValue := uint(9225)

  expectedNumberStr := "9.225"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint(originalUintValue, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint(originalUintValue, expectedPrecisionUint)\n"+
      "originalUintValue= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUintValue,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint_03(t *testing.T) {

  ePrefix := "TestDecimal_SetUint_03"

  origNumStr := "123.456"

  originalUintValue := uint(9225)

  expectedNumberStr := "9225"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum, err := new(Decimal).NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum, err := new(Decimal).NewNumStr(origNumStr)\n"+
      "origNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, origNumStr, err.Error())
    return
  }

  err = decNum.IsValid("Validating decNum-origNumStr")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum-origNumStr')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "decNum set to origNumStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if origNumStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and Decimal Number String Values ARE NOT Equal\n"+
      "Because origNumStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, origNumStr, decNumNumberStr)

    return
  }

  err = decNum.SetUint(originalUintValue, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.SetUint(originalUintValue, expectedPrecisionUint)\n"+
      "originalUintValue= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUintValue,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err = decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err = decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint_04(t *testing.T) {

  ePrefix := "TestDecimal_SetUint_04"

  originalUintValue := uint(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint(originalUintValue, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint(originalUintValue, expectedPrecisionUint)\n"+
      "originalUintValue= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUintValue,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint_05(t *testing.T) {

  ePrefix := "TestDecimal_SetUint_04"

  originalUintValue := uint(0)

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint(originalUintValue, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint(originalUintValue, expectedPrecisionUint)\n"+
      "originalUintValue= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUintValue,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint32_01(t *testing.T) {

  ePrefix := "TestDecimal_SetUint32_01"

  originalUint32Value := uint32(9225)

  expectedNumberStr := "92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)\n"+
      "originalUint32Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint32Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint32_02(t *testing.T) {

  ePrefix := "TestDecimal_SetUint32_02"

  originalUint32Value := uint32(9225)

  expectedNumberStr := "9.225"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)\n"+
      "originalUint32Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint32Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint32_03(t *testing.T) {

  ePrefix := "TestDecimal_SetUint32_03"

  baseNumberStr := "123.456"

  originalUint32Value := uint32(9225)

  expectedNumberStr := "9225"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseDecNum, err := new(Decimal).NewNumStr(baseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNum, err := new(Decimal).NewNumStr(baseNumberStr)\n"+
      "baseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseNumberStr, err.Error())
    return
  }

  err = baseDecNum.IsValid("Validating baseDecNum-initial")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.IsValid('Validating baseDecNum-initial')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumNumberStr, err := baseDecNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumberStr, err := baseDecNum.GetNumStr()\n"+
      "baseDecNum=baseNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumberStr != baseDecNumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseNumberStr and baseDecNumNumberStr Values ARE NOT Equal\n"+
      "Because baseNumberStr != baseDecNumNumberStr \n"+
      "Expected baseDecNumNumberStr = '%v'\n"+
      "  Actual baseDecNumNumberStr = '%v'\n\n",
      ePrefix, baseNumberStr, baseDecNumNumberStr)

    return
  }

  err = baseDecNum.SetUint32(originalUint32Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.SetUint32(originalUint32Value, expectedPrecisionUint)\n"+
      "originalUint32Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint32Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = baseDecNum.IsValid("Validating baseDecNum-final")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.IsValid('Validating baseDecNum-final')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumNumberStr, err = baseDecNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumberStr, err := baseDecNum.GetNumStr()\n"+
      "baseDecNum=originalUint32Value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumPrecisionUint, err := baseDecNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumPrecisionUint, err :=\n"+
      "  baseDecNum.GetPrecisionUint()\n"+
      "baseDecNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  baseDecNumSignValue, err := baseDecNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumSignValue, err := baseDecNum.GetSign()\n"+
      "baseDecNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  baseDecNumNumSeps, err := baseDecNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumSeps, err := baseDecNum.GetNumericSeparatorsDto()\n"+
      "baseDecNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != baseDecNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != baseDecNumNumberStr \n"+
      "Expected baseDecNumNumberStr = '%v'\n"+
      "  Actual baseDecNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, baseDecNumNumberStr)

    return
  }

  if expectedPrecisionUint != baseDecNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != baseDecNumPrecisionUint\n"+
      "Expected baseDecNumPrecisionUint = '%v'\n"+
      "  Actual baseDecNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, baseDecNumPrecisionUint)

    return
  }

  if expectedSignVal != baseDecNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != baseDecNumSignValue\n"+
      "Expected baseDecNumSignValue = '%v'\n"+
      "  Actual baseDecNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, baseDecNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(baseDecNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != baseDecNumNumSeps \n"+
      "Expected baseDecNumNumSeps = '%v'\n"+
      "  Actual baseDecNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), baseDecNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint32_04(t *testing.T) {

  ePrefix := "TestDecimal_SetUint32_04"

  originalUint32Value := uint32(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)\n"+
      "originalUint32Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint32Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint32_05(t *testing.T) {

  ePrefix := "TestDecimal_SetUint32_05"

  originalUint32Value := uint32(0)

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint32(originalUint32Value, expectedPrecisionUint)\n"+
      "originalUint32Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint32Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint64_01(t *testing.T) {

  ePrefix := "TestDecimal_SetUint64_01"

  originalUint64Value := uint64(9225)

  expectedNumberStr := "92.25"

  expectedPrecisionUint := uint(2)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)\n"+
      "originalUint64Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint64_02(t *testing.T) {

  ePrefix := "TestDecimal_SetUint64_02"

  originalUint64Value := uint64(9225)

  expectedNumberStr := "9.225"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)\n"+
      "originalUint64Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint64_03(t *testing.T) {

  ePrefix := "TestDecimal_SetUint64_03"

  baseNumberStr := "123.456"

  originalUint64Value := uint64(9225)

  expectedNumberStr := "9225"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  baseDecNum, err := new(Decimal).NewNumStr(baseNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNum, err := new(Decimal).NewNumStr(baseNumberStr)\n"+
      "baseNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseNumberStr, err.Error())
    return
  }

  err = baseDecNum.IsValid("Validating baseDecNum-initial")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.IsValid('Validating baseDecNum-initial')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumNumberStr, err := baseDecNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumberStr, err := baseDecNum.GetNumStr()\n"+
      "baseDecNum=baseNumberStr\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if baseNumberStr != baseDecNumNumberStr {
    t.Errorf("%v\n"+
      "Error: baseNumberStr and baseDecNumNumberStr Values ARE NOT Equal\n"+
      "Because baseNumberStr != baseDecNumNumberStr \n"+
      "Expected baseDecNumNumberStr = '%v'\n"+
      "  Actual baseDecNumNumberStr = '%v'\n\n",
      ePrefix, baseNumberStr, baseDecNumNumberStr)

    return
  }

  err = baseDecNum.SetUint64(originalUint64Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.SetUint64(originalUint64Value, expectedPrecisionUint)\n"+
      "originalUint64Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = baseDecNum.IsValid("Validating baseDecNum-final")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = baseDecNum.IsValid('Validating baseDecNum-final')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumNumberStr, err = baseDecNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumberStr, err := baseDecNum.GetNumStr()\n"+
      "baseDecNum=originalUint32Value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  baseDecNumPrecisionUint, err := baseDecNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumPrecisionUint, err :=\n"+
      "  baseDecNum.GetPrecisionUint()\n"+
      "baseDecNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  baseDecNumSignValue, err := baseDecNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumSignValue, err := baseDecNum.GetSign()\n"+
      "baseDecNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  baseDecNumNumSeps, err := baseDecNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "baseDecNumNumSeps, err := baseDecNum.GetNumericSeparatorsDto()\n"+
      "baseDecNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, baseDecNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != baseDecNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != baseDecNumNumberStr \n"+
      "Expected baseDecNumNumberStr = '%v'\n"+
      "  Actual baseDecNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, baseDecNumNumberStr)

    return
  }

  if expectedPrecisionUint != baseDecNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != baseDecNumPrecisionUint\n"+
      "Expected baseDecNumPrecisionUint = '%v'\n"+
      "  Actual baseDecNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, baseDecNumPrecisionUint)

    return
  }

  if expectedSignVal != baseDecNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != baseDecNumSignValue\n"+
      "Expected baseDecNumSignValue = '%v'\n"+
      "  Actual baseDecNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, baseDecNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(baseDecNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != baseDecNumNumSeps \n"+
      "Expected baseDecNumNumSeps = '%v'\n"+
      "  Actual baseDecNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), baseDecNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint64_04(t *testing.T) {

  ePrefix := "TestDecimal_SetUint64_02"

  originalUint64Value := uint64(0)

  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)\n"+
      "originalUint64Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}

func TestDecimal_SetUint64_05(t *testing.T) {

  ePrefix := "TestDecimal_SetUint64_05"

  originalUint64Value := uint64(0)

  expectedNumberStr := "0.000"

  expectedPrecisionUint := uint(3)

  expectedSignVal := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  decNum := new(Decimal).New()

  err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNum.SetUint64(originalUint64Value, expectedPrecisionUint)\n"+
      "originalUint64Value= '%v'\n"+
      "expectedPrecisionUint= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalUint64Value,
      expectedPrecisionUint,
      err.Error())

    return
  }

  err = decNum.IsValid("Validating decNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum.IsValid('Validating decNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumNumberStr, err := decNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumberStr, err := decNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumPrecisionUint, err := decNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumPrecisionUint, err :=\n"+
      "  decNum.GetPrecisionUint()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumSignValue, err := decNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumSignValue, err := decNum.GetSign()\n"+
      "decNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumNumberStr, err.Error())
    return
  }

  decNumNumSeps, err := decNum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumNumSeps, err := decNum.GetNumericSeparatorsDto()\n"+
      "decNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumNumberStr, err.Error())
    return
  }

  if expectedNumberStr != decNumNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumberStr and Decimal Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != decNumNumberStr \n"+
      "Expected decNumNumberStr = '%v'\n"+
      "  Actual decNumNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, decNumNumberStr)

    return
  }

  if expectedPrecisionUint != decNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected/dec Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != decNumPrecisionUint\n"+
      "Expected decNumPrecisionUint = '%v'\n"+
      "  Actual decNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNumPrecisionUint)

    return
  }

  if expectedSignVal != decNumSignValue {
    t.Errorf("%v\n"+
      "Error: expected & dec Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignVal != decNumSignValue\n"+
      "Expected decNumSignValue = '%v'\n"+
      "  Actual decNumSignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNumSignValue)

    return
  }

  if !expectedNumSeps.Equal(decNumNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != decNumNumSeps \n"+
      "Expected decNumNumSeps = '%v'\n"+
      "  Actual decNumNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), decNumNumSeps.String())

    return
  }

  return
}
