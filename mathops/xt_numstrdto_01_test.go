package mathops

import (
  "testing"
)

func TestNumStrDto_AddNumStrs_01(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_01"

  originalNumberStr1 := "-9589.21"

  originalNumberStr2 := "9211.40"

  expectedResultNumStr := "-377.81"

  expectedResultAbsIntStr := "377"

  expectedResultAbsFracStr := "81"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_02(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_02"

  originalNumberStr1 := "9589.21"

  originalNumberStr2 := "-9211.40"

  expectedResultNumStr := "377.81"

  expectedResultAbsIntStr := "377"

  expectedResultAbsFracStr := "81"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_03(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_03"

  originalNumberStr1 := "9589.21"

  originalNumberStr2 := "9211.40"

  expectedResultNumStr := "18800.61"

  expectedResultAbsIntStr := "18800"

  expectedResultAbsFracStr := "61"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_04(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_04"

  originalNumberStr1 := "-9589.21"

  originalNumberStr2 := "-9211.40"

  expectedResultNumStr := "-18800.61"

  expectedResultAbsIntStr := "18800"

  expectedResultAbsFracStr := "61"

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_05(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_05"

  originalNumberStr1 := "2"

  originalNumberStr2 := "3"

  expectedResultNumStr := "5"

  expectedResultAbsIntStr := "5"

  expectedResultAbsFracStr := ""

  expectedPrecisionInt := 2

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_06(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_06"

  originalNumberStr1 := "2"

  originalNumberStr2 := "0.0"

  expectedResultNumStr := "2.0"

  expectedResultAbsIntStr := "2"

  expectedResultAbsFracStr := "0"

  expectedPrecisionInt := 1

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_07(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_07"

  originalNumberStr1 := "0"

  originalNumberStr2 := "0.0"

  expectedResultNumStr := "0.0"

  expectedResultAbsIntStr := "0"

  expectedResultAbsFracStr := "0"

  expectedPrecisionInt := 1

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_08(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_08"

  originalNumberStr1 := "-6"

  originalNumberStr2 := "67.521"

  expectedResultNumStr := "61.521"

  expectedResultAbsIntStr := "61"

  expectedResultAbsFracStr := "521"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_09(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_09"

  originalNumberStr1 := "67.521"

  originalNumberStr2 := "-6"

  expectedResultNumStr := "61.521"

  expectedResultAbsIntStr := "61"

  expectedResultAbsFracStr := "521"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_10(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_10"

  originalNumberStr1 := "-67.521"

  originalNumberStr2 := "67.521"

  expectedResultNumStr := "0.000"

  expectedResultAbsIntStr := "0"

  expectedResultAbsFracStr := "000"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}

func TestNumStrDto_AddNumStrs_11(t *testing.T) {

  ePrefix := "TestNumStrDto_AddNumStrs_11"

  originalNumberStr1 := "67.521"

  originalNumberStr2 := "-67.521"

  expectedResultNumStr := "0.000"

  expectedResultAbsIntStr := "0"

  expectedResultAbsFracStr := "000"

  expectedPrecisionInt := 3

  expectedPrecisionUint := uint(expectedPrecisionInt)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  nDtoBase := new(NumStrDto).New()

  expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDto, err := nDtoBase.ParseNumStr(expectedResultNumStr)\n"+
      "expectedResultNumStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedResultNumStr, err.Error())
    return
  }

  err = expectedNumStrDto.IsValid("Validating expectedNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumStrDto.IsValid('Validating expectedNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedNumStrDtoNumberStr, err := expectedNumStrDto.GetNumStr()\n"+
      "expectedNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != expectedNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: expectedNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != expectedNumStrDtoNumberStr\n"+
      "Expected expectedNumStrDtoNumberStr = '%v'\n"+
      "  Actual expectedNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, expectedNumStrDtoNumberStr)

    return
  }

  originalNumStrDto1, err := nDtoBase.ParseNumStr(originalNumberStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr1)\n"+
      "originalNumberStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr1, err.Error())
    return
  }

  err = originalNumStrDto1.IsValid("Validating originalNumStrDto1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto1.IsValid('Validating originalNumStrDto1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto1NumberStr, err := originalNumStrDto1.GetNumStr()\n"+
      "originalNumStrDto1 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr1 != originalNumStrDto1NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto1 Number String is INVALID!\n"+
      "Because originalNumberStr1 != originalNumStrDto1NumberStr\n"+
      "Expected originalNumStrDto1NumberStr = '%v'\n"+
      "  Actual originalNumStrDto1NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr1, originalNumStrDto1NumberStr)

    return
  }

  originalNumStrDto2, err := nDtoBase.ParseNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2, err := nDtoBase.ParseNumStr(\n"+
      "  originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = originalNumStrDto2.IsValid("Validating originalNumStrDto2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = originalNumStrDto2.IsValid('Validating originalNumStrDto2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "originalNumStrDto2NumberStr, err := originalNumStrDto2.GetNumStr()\n"+
      "originalNumStrDto2 set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != originalNumStrDto2NumberStr {
    t.Errorf("%v\n"+
      "Error: originalNumStrDto2 Number String is INVALID!\n"+
      "Because originalNumberStr2 != originalNumStrDto2NumberStr\n"+
      "Expected originalNumStrDto2NumberStr = '%v'\n"+
      "  Actual originalNumStrDto2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, originalNumStrDto2NumberStr)

    return
  }

  resultNumStrDto, err := nDtoBase.AddNumStrs(originalNumStrDto1, originalNumStrDto2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDto, err := nDtoBase.AddNumStrs(\n"+
      "  originalNumStrDto1, originalNumStrDto2)\n"+
      "originalNumStrDto1= '%v'\n"+
      "originalNumStrDto2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      originalNumStrDto1NumberStr,
      originalNumStrDto2NumberStr,
      err.Error())

    return
  }

  err = resultNumStrDto.IsValid("Validating resultNumStrDto")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultNumStrDto.IsValid('Validating resultNumStrDto')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumberStr, err := resultNumStrDto.GetNumStr()\n"+
      "resultNumStrDto set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoPrecisionInt := resultNumStrDto.GetPrecision()

  resultNumStrDtoPrecisionUint, err := resultNumStrDto.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoPrecisionUint, err :=\n"+
      "  resultNumStrDto.GetPrecisionUint()\n"+
      "resultNumStrDtoResult= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoSignValue, err := resultNumStrDto.GetSign()\n"+
      "resultNumStrDto= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoNumSeps, err := resultNumStrDto.GetNumericSeparatorsDto()\n"+
      "resultNumStrDto= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStrDtoNumberStr, err.Error())
    return
  }

  resultNumStrDtoAbsIntRunes, err := resultNumStrDto.GetAbsIntRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsIntRunes, err := \n"+
      "   resultNumStrDto.GetAbsIntRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsIntStr := string(resultNumStrDtoAbsIntRunes)

  resultNumStrDtoAbsFracRunes, err := resultNumStrDto.GetAbsFracRunes()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStrDtoAbsFracRunes, err := \n"+
      "   resultNumStrDto.GetAbsFracRunes()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStrDtoAbsFracStr := string(resultNumStrDtoAbsFracRunes)

  if expectedResultNumStr != resultNumStrDtoNumberStr {
    t.Errorf("%v\n"+
      "Error: resultNumStrDto Number String is INVALID!\n"+
      "Because expectedResultNumStr != resultNumStrDtoNumberStr\n"+
      "Expected resultNumStrDtoNumberStr = '%v'\n"+
      "  Actual resultNumStrDtoNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStrDtoNumberStr)

    return
  }

  if expectedPrecisionInt != resultNumStrDtoPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != resultNumStrDtoPrecisionInt\n"+
      "Expected resultNumStrDtoPrecisionInt = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, resultNumStrDtoPrecisionInt)

    return
  }

  if expectedPrecisionUint != resultNumStrDtoPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Expected & resultNumStrDto Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != resultNumStrDtoPrecisionUint\n"+
      "Expected resultNumStrDtoPrecisionUint = '%v'\n"+
      "  Actual resultNumStrDtoPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultNumStrDtoPrecisionUint)

    return
  }

  if expectedSignValue != resultNumStrDtoSignValue {
    t.Errorf("%v\n"+
      "Error: expected & resultNumStrDto Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != resultNumStrDtoSignValue\n"+
      "Expected resultNumStrDtoSignValue = '%v'\n"+
      "  Actual resultNumStrDtoSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultNumStrDtoSignValue)

    return
  }

  if !expectedNumSeps.Equal(resultNumStrDtoNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != resultNumStrDtoNumSeps \n"+
      "Expected resultNumStrDtoNumSeps = '%v'\n"+
      "  Actual resultNumStrDtoNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumStrDtoNumSeps.String())

    return
  }

  if expectedResultAbsIntStr != resultNumStrDtoAbsIntStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Int Strings ARE NOT EQUAL!\n"+
      "Because expectedResultAbsIntStr != resultNumStrDtoAbsIntStr\n"+
      "Expected resultNumStrDtoAbsIntStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsIntStr = '%v'\n\n",
      ePrefix, expectedResultAbsIntStr, resultNumStrDtoAbsIntStr)

    return
  }

  if expectedResultAbsFracStr != resultNumStrDtoAbsFracStr {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Absolute Frac Strings ARE NOT EQUAL!\n"+
      "Because!!!!\n"+
      "Expected resultNumStrDtoAbsFracStr = '%v'\n"+
      "  Actual resultNumStrDtoAbsFracStr = '%v'\n\n",
      ePrefix, expectedResultAbsFracStr, resultNumStrDtoAbsFracStr)

    return
  }

  return
}
