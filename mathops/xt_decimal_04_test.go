package mathops

import (
  "math/big"
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

  nStr0 := "-999999.99999"

  nStr1 := "1.35"
  ePrecision := uint(2)
  eSignVal := 1

  nDto, err := NumStrDto{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). "+
      "nStr1='%v' Error='%v' ",
      nStr1, err.Error())
  }

  d1, err := Decimal{}.NewNumStr(nStr0)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). "+
      "nStr0='%v' Error='%v' ",
      nStr0, err.Error())
  }

  err = d1.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("Error returned by d1.SetNumStrDto(nDto). "+
      "nDto.GetNumStr()='%v' Error='%v' ",
      nDto.GetNumStr(), err.Error())
  }

  if nStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetNumStrDto_02(t *testing.T) {

  nStr0 := "-999999.99999"

  nStr1 := "-1555666.35"
  ePrecision := uint(2)
  eSignVal := -1

  nDto, err := NumStrDto{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). "+
      "nStr1='%v' Error='%v' ",
      nStr1, err.Error())
  }

  d1, err := Decimal{}.NewNumStr(nStr0)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). "+
      "nStr0='%v' Error='%v' ",
      nStr0, err.Error())
  }

  err = d1.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("Error returned by d1.SetNumStrDto(nDto). "+
      "nDto.GetNumStr()='%v' Error='%v' ",
      nDto.GetNumStr(), err.Error())
  }

  if nStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetNumStrDto_03(t *testing.T) {

  nStr0 := "-999999.99999"

  nStr1 := "1555777.123456"
  ePrecision := uint(6)
  eSignVal := 1

  nDto, err := NumStrDto{}.NewNumStr(nStr1)

  if err != nil {
    t.Errorf("Error returned by NumStrDto{}.NewNumStr(nStr1). "+
      "nStr1='%v' Error='%v' ",
      nStr1, err.Error())
  }

  d1, err := Decimal{}.NewNumStr(nStr0)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(nStr0). "+
      "nStr0='%v' Error='%v' ",
      nStr0, err.Error())
  }

  err = d1.SetNumStrDto(nDto)

  if err != nil {
    t.Errorf("Error returned by d1.SetNumStrDto(nDto). "+
      "nDto.GetNumStr()='%v' Error='%v' ",
      nDto.GetNumStr(), err.Error())
  }

  if nStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", nStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetFloat_01(t *testing.T) {

  fVal := float32(92.25)
  eNumStr1 := "92.25"
  ePrecision := uint(2)
  eSignVal := 1
  d1 := Decimal{}.New()

  err := d1.SetFloat32(fVal)

  if err != nil {
    t.Errorf("Received error from d1.SetFloat32(fVal). fVal= '%v' Error= %v ", fVal, err)
  }

  err = d1.SetPrecisionRound(ePrecision)

  if err != nil {
    t.Errorf("Error returned by d1.SetPrecisionRound(ePrecision).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetFloat_02(t *testing.T) {

  fVal := float32(-92.25)
  eNumStr1 := "-92.25"
  ePrecision := uint(2)
  eSignVal := -1
  d1 := Decimal{}.New()

  err := d1.SetFloat32(fVal)

  if err != nil {
    t.Errorf("Received error from d1.SetFloat32(fVal). fVal= '%v' Error= %v ", fVal, err)
  }

  err = d1.SetPrecisionRound(ePrecision)

  if err != nil {
    t.Errorf("Error returned by d1.SetPrecisionRound(ePrecision).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetFloat64_01(t *testing.T) {

  var fVal = 92.256
  eNumStr1 := "92.256"
  ePrecision := uint(3)
  eSignVal := 1
  d1 := Decimal{}.New()

  err := d1.SetFloat64(fVal)

  if err != nil {
    t.Errorf("Received error from d1.SetFloat64(fVal). fVal= '%v' Error= %v ", fVal, err)
  }

  err = d1.SetPrecisionRound(ePrecision)

  if err != nil {
    t.Errorf("Error returned by d1.SetPrecisionRound(ePrecision).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())

  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetFloat64_02(t *testing.T) {

  fVal := -92.25
  eNumStr1 := "-92.25"
  ePrecision := uint(2)
  eSignVal := -1
  d1 := Decimal{}.New()

  err := d1.SetFloat64(fVal)

  if err != nil {
    t.Errorf("Received error from d1.SetFloat64(fVal). fVal= '%v' Error= %v ", fVal, err)
  }

  err = d1.SetPrecisionRound(ePrecision)

  if err != nil {
    t.Errorf("Error returned by d1.SetPrecisionRound(ePrecision).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}
func TestDecimal_SetFloatBig_01(t *testing.T) {

  eNumStr1 := "92.256"
  fVal, isOk := big.NewFloat(0.0).SetString(eNumStr1)

  if !isOk {
    t.Errorf("bigFloat.SetString failed to convert eNumStr1. eNumStr1= '%v'", eNumStr1)
  }

  ePrecision := uint(3)
  eSignVal := 1
  d1 := Decimal{}.New()

  err := d1.SetFloatBig(fVal)

  if err != nil {
    t.Errorf("d1.SetFloatBig(fVal) returned an error. fVal= '%v' Error= %v", eNumStr1, err)
  }

  err = d1.SetPrecisionRound(ePrecision)

  if err != nil {
    t.Errorf("d1.SetPrecisionRound(ePrecision).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetFloatBig_02(t *testing.T) {

  eNumStr1 := "-92.25"
  ePrecision := uint(2)
  eSignVal := -1
  fVal, isOk := big.NewFloat(0.0).SetString(eNumStr1)

  if !isOk {
    t.Errorf("bigFloat.SetString failed to convert eNumStr1. eNumStr1= '%v'", eNumStr1)
  }

  d1 := Decimal{}.New()

  err := d1.SetFloatBig(fVal)

  if err != nil {
    t.Errorf("d1.SetFloatBig(fVal) returned an error. fVal= '%v' Error= %v", eNumStr1, err)
  }

  err = d1.SetPrecisionRound(2)

  if err != nil {
    t.Errorf("d1.SetPrecisionRound(2).\n"+
      "ePrecision='%v' Error='%v' ", ePrecision, err.Error())
  }

  if eNumStr1 != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr1, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetIntFracStrings_01(t *testing.T) {

  intStr := "123"
  fracStr := "456"
  eNumStr := "123.456"
  eSignVal := 1
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

  if err != nil {
    t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
  }

  if eNumStr != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetIntFracStrings_02(t *testing.T) {

  intStr := "123"
  fracStr := "0456"
  eNumStr := "-123.0456"
  eSignVal := -1
  ePrecision := uint(4)

  d1 := Decimal{}.New()

  err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

  if err != nil {
    t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
  }

  if eNumStr != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetIntFracStrings_03(t *testing.T) {

  intStr := "-0"
  fracStr := "04#5 6"
  eNumStr := "0.0456"
  eSignVal := 1
  ePrecision := uint(4)

  d1 := Decimal{}.New()

  err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

  if err != nil {
    t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
  }

  if eNumStr != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetIntFracStrings_04(t *testing.T) {

  intStr := "-0"
  fracStr := ".04#5 6"
  eNumStr := "0.0456"
  eSignVal := 1
  ePrecision := uint(4)

  d1 := Decimal{}.New()

  err := d1.SetIntFracStrings(intStr, fracStr, eSignVal)

  if err != nil {
    t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
  }

  if eNumStr != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

}

func TestDecimal_SetIntFracStrings_05(t *testing.T) {

  intStr := "-0"
  fracStr := ".04#5 6"
  eNumStr := "0!0456"
  eSignVal := 1
  ePrecision := uint(4)

  d1 := Decimal{}.New()

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '!'
  expectedNumSeps.ThousandsSeparator = 'X'
  expectedNumSeps.CurrencySymbol = '^'

  err := d1.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by d1.SetNumericSeparatorsDto(expectedNumSeps) "+
      "Error='%v' ", err.Error())
  }

  err = d1.SetIntFracStrings(intStr, fracStr, eSignVal)

  if err != nil {
    t.Errorf("d1.SetIntFracStrings(eSignVal, intStr, fracStr) returned an error. eSignVal= '%v' intStr= '%v' fracStr= '%v' Error= %v", eSignVal, intStr, fracStr, err)
  }

  if eNumStr != d1.GetNumStr() {
    t.Errorf("Expected NumStr = '%v'. Instead got NumStr= '%v'", eNumStr, d1.GetNumStr())
  }

  if int(ePrecision) != d1.GetPrecision() {
    t.Errorf("Expected precision= '%v'. Intead, got precision= '%v' ", ePrecision, d1.GetPrecision())
  }

  if eSignVal != d1.GetSign() {
    t.Errorf("Expected sign Value= '%v'. Intead, got sign Value = '%v' ", eSignVal, d1.GetSign())
  }

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumSeps := d1.GetNumericSeparatorsDto()

  if !actualNumSeps.Equal(expectedNumSeps) {
    t.Errorf("Error: Expected numeric separators are NOT equal to actual numeric separators. "+
      "Expected numSeps= %v   Actual numSeps= %v", expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestDecimal_SetUint_01(t *testing.T) {

  fVal := uint(9225)
  eNumStr := "92.25"
  ePrecision := uint(2)

  d1 := Decimal{}.New()

  d1.SetUint(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint_02(t *testing.T) {

  fVal := uint(9225)
  eNumStr := "9.225"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint_03(t *testing.T) {

  origNumStr := "123.456"

  d1, err := Decimal{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(origNumStr). ")
  }

  actualNumStr := d1.GetNumStr()

  if origNumStr != actualNumStr {
    t.Errorf("Error: Expected origNumStr='%v'. Instead, origNumStr='%v'. ",
      origNumStr, actualNumStr)
  }

  fVal := uint(9225)
  eNumStr := "9225"
  ePrecision := uint(0)

  d1.SetUint(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr = d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint_04(t *testing.T) {

  fVal := uint(0)
  eNumStr := "0"
  ePrecision := uint(0)

  d1 := Decimal{}.New()

  d1.SetUint(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint_05(t *testing.T) {

  fVal := uint(0)
  eNumStr := "0.000"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint32_01(t *testing.T) {

  fVal := uint32(9225)
  eNumStr := "92.25"
  ePrecision := uint(2)

  d1 := Decimal{}.New()

  d1.SetUint32(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint32_02(t *testing.T) {

  fVal := uint32(9225)
  eNumStr := "9.225"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint32(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint32_03(t *testing.T) {

  origNumStr := "123.456"

  d1, err := Decimal{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(origNumStr). ")
  }

  actualNumStr := d1.GetNumStr()

  if origNumStr != actualNumStr {
    t.Errorf("Error: Expected origNumStr='%v'. Instead, origNumStr='%v'. ",
      origNumStr, actualNumStr)
  }

  fVal := uint32(9225)
  eNumStr := "9225"
  ePrecision := uint(0)

  d1.SetUint32(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr = d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint32_04(t *testing.T) {

  fVal := uint32(0)
  eNumStr := "0"
  ePrecision := uint(0)

  d1 := Decimal{}.New()

  d1.SetUint32(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint32_05(t *testing.T) {

  fVal := uint32(0)
  eNumStr := "0.000"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint32(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint64_01(t *testing.T) {

  fVal := uint64(9225)
  eNumStr := "92.25"
  ePrecision := uint(2)

  d1 := Decimal{}.New()

  d1.SetUint64(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint64_02(t *testing.T) {

  fVal := uint64(9225)
  eNumStr := "9.225"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint64(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint64_03(t *testing.T) {

  origNumStr := "123.456"

  d1, err := Decimal{}.NewNumStr(origNumStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(origNumStr). ")
  }

  actualNumStr := d1.GetNumStr()

  if origNumStr != actualNumStr {
    t.Errorf("Error: Expected origNumStr='%v'. Instead, origNumStr='%v'. ",
      origNumStr, actualNumStr)
  }

  fVal := uint64(9225)
  eNumStr := "9225"
  ePrecision := uint(0)

  d1.SetUint64(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr = d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint64_04(t *testing.T) {

  fVal := uint64(0)
  eNumStr := "0"
  ePrecision := uint(0)

  d1 := Decimal{}.New()

  d1.SetUint64(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}

func TestDecimal_SetUint64_05(t *testing.T) {

  fVal := uint64(0)
  eNumStr := "0.000"
  ePrecision := uint(3)

  d1 := Decimal{}.New()

  d1.SetUint64(fVal, ePrecision)

  if !d1.GetIsValid() {
    t.Errorf("Expected IsValid == 'true'. Instead got IsValid= '%v'", d1.GetIsValid())
  }

  actualNumStr := d1.GetNumStr()

  if eNumStr != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      eNumStr, actualNumStr)
  }

}
