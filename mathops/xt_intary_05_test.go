package mathops

import (
	"math/big"
  "strconv"
  "testing"
)

func TestIntAry_MultiplyByTenToPower_01(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_01"

  originalNumberStr := "457.3"

  targetPowerOfTen := uint(2)

  expectedNumberStr := "45730"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{4, 5, 7, 3, 0}

  expectedArrayLen := len(expectedAry)

  intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

	err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expectedexpected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTenToPower_02(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_02"

  originalNumberStr := "457.3"

  targetPowerOfTen := uint(2)

  expectedNumberStr := "45730"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{4, 5, 7, 3, 0}

  expectedArrayLen := len(expectedAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTenToPower_03(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_03"

  originalNumberStr := "457.3"

  targetPowerOfTen := uint(10)

  expectedNumberStr := "4573000000000"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{4, 5, 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}

  expectedArrayLen := len(expectedAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTenToPower_04(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_04"

  originalNumberStr := "457.3"

  targetPowerOfTen := uint(0)

  expectedNumberStr := "457.3"

  expectedPrecisionUint := uint(1)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{4, 5, 7, 3}

  expectedArrayLen := len(expectedAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTenToPower_05(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_05"

  originalNumberStr := "-457.3"

  targetPowerOfTen := uint(1)

  expectedNumberStr := "-4573"

  expectedPrecisionUint := uint(0)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{4, 5, 7, 3}

  expectedArrayLen := len(expectedAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTenToPower_06(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTenToPower_06"

  originalNumberStr := "0"

  targetPowerOfTen := uint(2)

  expectedNumberStr := "000"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedAry := []uint8{0, 0, 0}

  expectedArrayLen := len(expectedAry)

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTenToPower(targetPowerOfTen)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTenToPower(targetPowerOfTen)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTen= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTen,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayElements, intArrayElementsLen, err := intAry.GetIntAryElements()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedArrayLen != intArrayElementsLen {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Array Lengths ARE NOT EQUAL!\n"+
      "Because expectedArrayLen != intArrayElementsLen\n"+
      "Expected intArrayElementsLen = '%v'\n"+
      "  Actual intArrayElementsLen = '%v'\n\n",
      ePrefix, expectedArrayLen, intArrayElementsLen)

    return
  }

  for i := 0; i < intArrayElementsLen; i++ {

    if expectedAry[i] != intArrayElements[i] {
      t.Errorf("%v\n"+
        "Error: Expected and Actual Array Elements Don't Match!\n"+
        "Because expectedAry[%v] != intArrayElements[%v]\n"+
        "Expected intArrayElements[%v] = '%v'\n"+
        "  Actual intArrayElements[%v] = '%v'\n"+
        "Expected Number String = '%v'\n"+
        "IntAry Number String   = '%v'\n",
        ePrefix, i, i, i, expectedAry[i], i, expectedAry[i],
        expectedNumberStr, intAryNumberStr)

      return
    }

  }

  return
}

func TestIntAry_MultiplyByTwoToPower_01(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTwoToPower_01"

  originalNumberStr := "23"

  targetPowerOfTwo := uint(19)

  expectedNumberStr := "12058624"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry := new(IntAry).New()

	err := intAry.SetIntAryWithNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err := intAry.SetIntAryWithNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTwo= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTwo,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyByTwoToPower_02(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTwoToPower_02"

  originalNumberStr := "23"

  targetPowerOfTwo := uint(0)

  expectedNumberStr := "23"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTwo= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTwo,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyByTwoToPower_03(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyByTwoToPower_03"

  originalNumberStr := "23"

  targetPowerOfTwo := uint(1)

  expectedNumberStr := "46"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAry Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, originalNumberStr, intAryNumberStr)

    return
  }

  err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAry.MultiplyByTwoToPower(targetPowerOfTwo)\n"+
      "intAry= '%v'\n" +
      "targetPowerOfTwo= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryNumberStr,
      targetPowerOfTwo,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating final intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating final intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err = intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionUint, err := intAry.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryPrecisionUint, err :=\n"+
      "  intAry.GetPrecisionUint()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionUint != intAryPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryPrecisionUint\n"+
      "Expected intAryPrecisionUint = '%v'\n"+
      "  Actual intAryPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryPrecisionUint)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_01(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_01"

  originalNumberStrThis := "96.8524"

  originalNumberStr2 := "8574.21396845"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "830433.20095790678"

  expectedPrecisionUint := uint(11)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

	err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)\n"+
      "intAryThis= '%v'\n" +
      "intAry2= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_02(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_02"

  originalNumberStrThis := "-96.8524"

  originalNumberStr2 := "8574.21396845"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr:= "-830433.20095790678"

  expectedPrecisionUint := uint(11)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)\n"+
      "intAryThis= '%v'\n" +
      "intAry2= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_03(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_03"

  originalNumberStrThis := "96.8524"

  originalNumberStr2 := "-8574.21396845"

  //                                     1         2         3
  //                          0.1234567890123456789012345678901234567
  expectedNumberStr := "-830433.20095790678"

  expectedPrecisionUint := uint(11)

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_04(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_04"

  originalNumberStrThis := "-96.8524"

  originalNumberStr2 := "-8574.21396845"

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "830433.20095790678"

  expectedPrecisionUint := uint(11)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, -1, -1)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_05(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_05"

  originalNumberStrThis := "96.8524"

  originalNumberStr2 := "8574.21396845"

  maxPrecision := 6

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr := "830433.200958"

  expectedPrecisionUint := uint(6)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      maxPrecision,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_06(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_06"

  originalNumberStrThis := "0"

  originalNumberStr2 := "-8574.21396845"

  maxPrecision := -1

  minimumPrecision := 6

  //                                    1         2         3
  //                         0.1234567890123456789012345678901234567
  expectedNumberStr :=      "0.00000000"

  expectedPrecisionUint := uint(8)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, minimumPrecision, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_07(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_07"

  originalNumberStrThis := "96.8524"

  originalNumberStr2 := "0"

  minimumPrecision := 0

  maxPrecision := -1

  //                               1         2         3
  //                    0.1234567890123456789012345678901234567
  expectedNumberStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, minimumPrecision, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_08(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_08"

  originalNumberStrThis := "40"

  originalNumberStr2 := "2"

  minimumPrecision := 0

  maxPrecision := -1

  //                                1         2         3
  //                     0.1234567890123456789012345678901234567
  expectedNumberStr := "80"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, minimumPrecision, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

	thisStats := intAryThis.GetIntAryStats()

  if thisStats.IntAryLen != 2 {
    t.Errorf("%v\n" +
      "Error: Unexpected Result!\n" +
      "Because thisStats.IntAryLen != 2\n" +
      "Expected thisStats.IntAryLen = '2'\n" +
      "  Actual thisStats.IntAryLen = '%v'\n\n",
      ePrefix, thisStats.IntAryLen)

    return
  }

  if thisStats.IntegerLen != 2 {
    t.Errorf("%v\n" +
      "Error: Unexpected Result!\n" +
      "Because  thisStats.IntegerLen != 2\n" +
      "Expected thisStats.IntegerLen = '2'\n" +
      "  Actual thisStats.IntegerLen = '%v'\n\n",
      ePrefix, thisStats.IntegerLen)

    return
  }

  return
}

func TestIntAry_MultiplyThisBy_09(t *testing.T) {

  ePrefix := "TestIntAry_MultiplyThisBy_09"

  originalNumberStrThis := "999.99"

  originalNumberStr2 := "99.9"

  minimumPrecision := -1

  maxPrecision := -1

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedNumberStr := "99899.001"

  expectedPrecisionUint := uint(3)

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThis, err := new(IntAry).NewNumStr(originalNumberStrThis)\n"+
      "originalNumberStrThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStrThis, err.Error())
    return
  }

  err = intAryThis.IsValid("Validating intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err := intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStrThis != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Original and intAryThis Number String Values ARE NOT Equal\n"+
      "Because originalNumberStrThis != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, originalNumberStrThis, intAryThisNumberStr)

    return
  }

  intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2, err := new(IntAry).NewNumStr(originalNumberStr2)\n"+
      "originalNumberStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr2, err.Error())
    return
  }

  err = intAry2.IsValid("Validating intAry2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry2.IsValid('Validating intAry2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAry2NumberStr, err := intAry2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAry2NumberStr, err := intAry2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalNumberStr2 != intAry2NumberStr {
    t.Errorf("%v\n"+
      "Error: Original2 and intAry2 Number String Values ARE NOT Equal\n"+
      "Because originalNumberStr2 != intAry2NumberStr \n"+
      "Expected intAry2NumberStr = '%v'\n"+
      "  Actual intAry2NumberStr = '%v'\n\n",
      ePrefix, originalNumberStr2, intAry2NumberStr)

    return
  }

  err = intAryThis.MultiplyThisBy(&intAry2, minimumPrecision, maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.MultiplyThisBy(&intAry2, maxPrecision, maxPrecision)\n"+
      "intAryThis= '%v'\n"+
      "intAry2= '%v'\n"+
      "minimumPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      intAryThisNumberStr,
      intAry2NumberStr,
      minimumPrecision,
      maxPrecision,
      err.Error())

    return
  }

  err = intAryThis.IsValid("Validating final intAryThis")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAryThis.IsValid('Validating final intAryThis')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisNumberStr, err = intAryThis.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumberStr, err := intAryThis.GetNumStr()\n"+
      "intAryThis set to final value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryThisPrecisionUint, err := intAryThis.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisPrecisionUint, err :=\n"+
      "  intAryThis.GetPrecisionUint()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisNumSeps, err := intAryThis.GetNumericSeparatorsDto()\n"+
      "intAryThis= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  intAryThisSignValue, err := intAryThis.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryThisSignValue, err := intAryThis.GetSign()\n"+
      "intAryThis= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryThisNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryThisNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryThisNumberStr \n"+
      "Expected intAryThisNumberStr = '%v'\n"+
      "  Actual intAryThisNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryThisNumberStr)

    return
  }

  if expectedPrecisionUint != intAryThisPrecisionUint {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionUint != intAryThisPrecisionUint\n"+
      "Expected intAryThisPrecisionUint = '%v'\n"+
      "  Actual intAryThisPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, intAryThisPrecisionUint)

    return
  }

  if expectedSignValue != intAryThisSignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAryThis Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intAryThisSignValue\n"+
      "Expected intAryThisSignValue = '%v'\n"+
      "  Actual intAryThisSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intAryThisSignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryThisNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryThisNumSeps \n"+
      "Expected intAryThisNumSeps = '%v'\n"+
      "  Actual intAryThisNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryThisNumSeps.String())

    return
  }

  thisStats := intAryThis.GetIntAryStats()

  if thisStats.IntAryLen != 8 {
    t.Errorf("%v\n" +
      "Error: Unexpected Result!\n" +
      "Because thisStats.IntAryLen != 8\n" +
      "Expected thisStats.IntAryLen = '8'\n" +
      "  Actual thisStats.IntAryLen = '%v'\n\n",
      ePrefix, thisStats.IntAryLen)

    return
  }

  if thisStats.IntegerLen != 5 {
    t.Errorf("%v\n" +
      "Error: Unexpected Result!\n" +
      "Because  thisStats.IntegerLen != 5\n" +
      "Expected thisStats.IntegerLen = '5'\n" +
      "  Actual thisStats.IntegerLen = '%v'\n\n",
      ePrefix, thisStats.IntegerLen)

    return
  }

  return
}

func TestIntAry_NewNumStr_01(t *testing.T) {

  ePrefix := "TestIntAry_NewNumStr_01"


  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "579.123456000"

  expectedNumberStr := originalNumberStr

	expectedPrecisionInt := 9

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewNumStr_02(t *testing.T) {

  ePrefix := "TestIntAry_NewNumStr_02"


  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  originalNumberStr := "-579.123456000"

  expectedNumberStr := originalNumberStr

  expectedPrecisionInt := 9

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewNumStr(originalNumberStr)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStr(originalNumberStr)\n"+
      "originalNumberStr= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix, originalNumberStr, err.Error())
    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewNumStrWithNumSeps_01(t *testing.T) {

  ePrefix := "TestIntAry_NewNumStrWithNumSeps_01"


  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "579,123456000"

  expectedNumberStr := originalNumberStr

  expectedPrecisionInt := 9

  expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}
	frenchDecSeparator := ','
	frenchThousandsSeparator := ' '
	frenchCurrencySymbol := '€'

	expectedNumSeps.DecimalSeparator = frenchDecSeparator
	expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
	expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  intAry, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStrWithNumSeps(\n" +
      "  originalNumberStr, expectedNumSeps)\n"+
      "originalNumberStr= '%v'\n" +
      "expectedNumSeps= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewNumStrWithNumSeps_02(t *testing.T) {

  ePrefix := "TestIntAry_NewNumStrWithNumSeps_02"


  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "579.123456000"

  expectedNumberStr := originalNumberStr

  expectedPrecisionInt := 9

  expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

	expectedNumSeps.DecimalSeparator = '.'
	expectedNumSeps.ThousandsSeparator = ','
	expectedNumSeps.CurrencySymbol = '$'

  intAry, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStrWithNumSeps(\n" +
      "  originalNumberStr, expectedNumSeps)\n"+
      "originalNumberStr= '%v'\n" +
      "expectedNumSeps= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewNumStrWithNumSeps_03(t *testing.T) {

  ePrefix := "TestIntAry_NewNumStrWithNumSeps_03"


  //                                 1         2         3
  //                      0.1234567890123456789012345678901234567
  originalNumberStr := "579.123456000"

  expectedNumberStr := originalNumberStr

  expectedPrecisionInt := 9

  expectedSignValue := 1

	expectedNumSeps := NumericSeparatorDto{}

  intAry, err := new(IntAry).NewNumStrWithNumSeps(originalNumberStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewNumStrWithNumSeps(\n" +
      "  originalNumberStr, expectedNumSeps)\n"+
      "originalNumberStr= '%v'\n" +
      "expectedNumSeps= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewBigInt_01(t *testing.T) {

  ePrefix := "TestIntAry_NewBigInt_01"

	originalNumberBigInt := big.NewInt(123456)

	originalPrecisionInt := 3

  expectedNumberStr := "123.456"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewBigInt(originalNumberBigInt, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewBigInt(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigInt.Text(10),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewBigInt_02(t *testing.T) {

  ePrefix := "TestIntAry_NewBigInt_02"

  originalNumberBigInt := big.NewInt(-123456)

  originalPrecisionInt := 3

  expectedNumberStr := "-123.456"

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewBigInt(originalNumberBigInt, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewBigInt(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigInt.Text(10),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloatBig_01(t *testing.T) {

  ePrefix := "TestIntAry_NewFloatBig_01"

  originalNumberBigFloat := big.NewFloat(123.456)

  originalPrecisionInt := 3

  expectedNumberStr := "123.456"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloatBig(originalNumberBigFloat, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloatBig(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigFloat.Text('f', originalPrecisionInt),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloatBig_02(t *testing.T) {

  ePrefix := "TestIntAry_NewFloatBig_02"

  originalNumberBigFloat := big.NewFloat(-123.456)

  originalPrecisionInt := 3

  expectedNumberStr := "-123.456"

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloatBig(originalNumberBigFloat, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloatBig(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigFloat.Text('f', originalPrecisionInt),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloatBig_03(t *testing.T) {

  ePrefix := "TestIntAry_NewFloatBig_03"

  originalNumberBigFloat := big.NewFloat(123.456000)

  originalPrecisionInt := 3

  expectedNumberStr := "123.456"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloatBig(originalNumberBigFloat, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloatBig(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigFloat.Text('f', originalPrecisionInt),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloatBig_04(t *testing.T) {

  ePrefix := "TestIntAry_NewFloatBig_04"

  originalNumberBigFloat := big.NewFloat(123.456700)

  originalPrecisionInt := 3

  expectedNumberStr := "123.457"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloatBig(originalNumberBigFloat, originalPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloatBig(\n" +
      "  originalNumberBigInt, originalPrecisionInt)\n"+
      "originalNumberBigInt= '%v'\n" +
      "originalPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      originalNumberBigFloat.Text('f', originalPrecisionInt),
      originalPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloat64_01(t *testing.T) {

  ePrefix := "TestIntAry_NewFloat64_01"

  // Defaults to float64 Type
  originalNumberFloat64 := 123.456111

  originalNumberPrecisionRaw := 6

  originalInputPrecisionInt := 3

  expectedNumberStr := "123.456"

  expectedPrecisionInt := 3

  expectedSignValue := 1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloat64(originalNumberFloat64, originalInputPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloat64(\n" +
      "  originalNumberFloat64, originalInputPrecisionInt)\n"+
      "originalNumberFloat64= '%v'\n" +
      "originalInputPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalNumberFloat64, 'f', originalNumberPrecisionRaw, 64),
      originalInputPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloat64_02(t *testing.T) {

  ePrefix := "TestIntAry_NewFloat64_01"

  // Defaults to float64 Type
  originalNumberFloat64 := -123.456111

  originalNumberPrecisionRaw := 6

  originalInputPrecisionInt := 3

  expectedNumberStr := "-123.456"

  expectedPrecisionInt := 3

  expectedSignValue := -1

  expectedNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  intAry, err := new(IntAry).NewFloat64(originalNumberFloat64, originalInputPrecisionInt)

  if err != nil {
    t.Errorf("%v\n" +
      "Error returned by:\n"+
      "intAry, err := new(IntAry).NewFloat64(\n" +
      "  originalNumberFloat64, originalInputPrecisionInt)\n"+
      "originalNumberFloat64= '%v'\n" +
      "originalInputPrecisionInt= '%v'\n" +
      "Error= '%v'\n\n",
      ePrefix,
      strconv.FormatFloat(originalNumberFloat64, 'f', originalNumberPrecisionRaw, 64),
      originalInputPrecisionInt,
      err.Error())

    return
  }

  err = intAry.IsValid("Validating initial intAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intAry.IsValid('Validating initial intAry')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryNumberStr, err := intAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumberStr, err := intAry.GetNumStr()\n"+
      "intAry set to initial value\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intAryPrecisionInt := intAry.GetPrecision()

  intAryNumSeps, err := intAry.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intAryNumSeps, err := intAry.GetNumericSeparatorsDto()\n"+
      "intAry= '%v\n"+
      "Error= '%v'\n\n", ePrefix, intAryNumberStr, err.Error())
    return
  }

  intArySignValue, err := intAry.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArySignValue, err := intAry.GetSign()\n"+
      "intAry= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intAryNumberStr, err.Error())
    return
  }

  if expectedNumberStr != intAryNumberStr {
    t.Errorf("%v\n"+
      "Error: Expected and IntAry Number String Values ARE NOT Equal\n"+
      "Because expectedNumberStr != intAryNumberStr \n"+
      "Expected intAryNumberStr = '%v'\n"+
      "  Actual intAryNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intAryNumberStr)

    return
  }

  if expectedPrecisionInt != intAryPrecisionInt {
    t.Errorf("%v\n"+
      "Error: expected & intAry Precision Values ARE NOT EQUAL!\n"+
      "Because expectedPrecisionInt != intAryPrecisionInt\n"+
      "Expected intAryPrecisionInt = '%v'\n"+
      "  Actual intAryPrecisionInt = '%v'\n\n",
      ePrefix, expectedPrecisionInt, intAryPrecisionInt)

    return
  }

  if expectedSignValue != intArySignValue {
    t.Errorf("%v\n"+
      "Error: expected & intAry Sign Values ARE NOT EQUAL!\n"+
      "Because expectedSignValue != intArySignValue\n"+
      "Expected intArySignValue = '%v'\n"+
      "  Actual intArySignValue = '%v'\n\n",
      ePrefix, expectedSignValue, intArySignValue)

    return
  }

  if !expectedNumSeps.Equal(intAryNumSeps) {
    t.Errorf("%v\n"+
      "Error: Numeric Separator Values ARE NOT Equal!\n"+
      "Because expectedNumSeps != intAryNumSeps \n"+
      "Expected intAryNumSeps = '%v'\n"+
      "  Actual intAryNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), intAryNumSeps.String())

    return
  }

  return
}

func TestIntAry_NewFloat32_01(t *testing.T) {
	num := float64(123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat32_02(t *testing.T) {
	num := float64(-123.456111)
	precision := 3
	ia, err := new(IntAry).NewFloat64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewInt_01(t *testing.T) {
	num := int(123456)
	precision := uint(3)
	ia := new(IntAry).NewInt(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt_02(t *testing.T) {
	num := int(-123456)
	precision := uint(3)
	ia := new(IntAry).NewInt(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewIntExponent_01(t *testing.T) {

	intNum := int(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_02(t *testing.T) {

	intNum := int(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_03(t *testing.T) {

	intNum := int(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_04(t *testing.T) {

	intNum := int(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_05(t *testing.T) {

	intNum := int(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewIntExponent_06(t *testing.T) {

	intNum := int(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewIntExponent(intNum, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32_01(t *testing.T) {
	num := int32(123456)
	precision := uint(3)
	eStr := "123.456"
	eSignVal := 1

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_02(t *testing.T) {

	num := int32(-123456)
	precision := uint(3)
	eStr := "-123.456"
	eSignVal := -1

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_03(t *testing.T) {

	num := int32(0)
	precision := uint(0)
	eStr := "0"

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewInt32_04(t *testing.T) {

	num := int32(0)
	precision := uint(3)
	eStr := "0.000"

	ia := new(IntAry).NewInt32(num, precision)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewInt32Exponent_01(t *testing.T) {

	int32Num := int32(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_02(t *testing.T) {

	int32Num := int32(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_03(t *testing.T) {

	int32Num := int32(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_04(t *testing.T) {

	int32Num := int32(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_05(t *testing.T) {

	int32Num := int32(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt32Exponent_06(t *testing.T) {

	int32Num := int32(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewInt32Exponent(int32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64_01(t *testing.T) {
	int64Num := int64(123456)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "123.456"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_02(t *testing.T) {
	int64Num := int64(-123456)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_03(t *testing.T) {
	int64Num := int64(0)
	precision := uint(3)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "0.000"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_04(t *testing.T) {
	int64Num := int64(0)
	precision := uint(0)
	ia := new(IntAry).NewInt64(int64Num, precision)
	eStr := "0"
	eSignVal := 1

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64Exponent_01(t *testing.T) {

	int64Num := int64(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_02(t *testing.T) {

	int64Num := int64(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_03(t *testing.T) {

	int64Num := int64(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_04(t *testing.T) {

	int64Num := int64(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_05(t *testing.T) {

	int64Num := int64(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewInt64Exponent(int64Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewInt64Exponent_06(t *testing.T) {

	num := int64(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewInt64Exponent(num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewOne_01(t *testing.T) {

	expected := "1"
	ePrecision := 0

	ia := new(IntAry).NewOne(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewOne_02(t *testing.T) {

	expected := "1.00"
	ePrecision := 2

	ia := new(IntAry).NewOne(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewFive_01(t *testing.T) {

	expected := "5"
	ePrecision := 0

	ia := new(IntAry).NewFive(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewFive_02(t *testing.T) {

	expected := "5.00"
	ePrecision := 2

	ia := new(IntAry).NewFive(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTen_01(t *testing.T) {

	expected := "10"
	ePrecision := 0

	ia := new(IntAry).NewTen(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTen_02(t *testing.T) {

	expected := "10.00"
	ePrecision := 2

	ia := new(IntAry).NewTen(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewThree_01(t *testing.T) {

	expected := "3"
	ePrecision := 0

	ia := new(IntAry).NewThree(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewThree_02(t *testing.T) {

	expected := "3.00"
	ePrecision := 2

	ia := new(IntAry).NewThree(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTwo_01(t *testing.T) {

	expected := "2"
	ePrecision := 0

	ia := new(IntAry).NewTwo(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewTwo_02(t *testing.T) {

	expected := "2.00"
	ePrecision := 2

	ia := new(IntAry).NewTwo(ePrecision)

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, ia.GetNumStr())
	}

}

func TestIntAry_NewUint_01(t *testing.T) {

	ePrefix := "TestIntAry_NewUint_01"

	uintNum := uint(123456)
	precision := uint(3)
	signVal := 1

	ia, err := new(IntAry).NewUint(uintNum, signVal, precision)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint(uintNum, signVal, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	eStr := "123.456"

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n"+
			"Expected ia.GetNumStr()== %v\n"+
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, iaNumStr, eStr)

		return
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n"+
			"Expected ia.GetPrecisionInt() == '%v'\n"+
			"Instead ia.GetPrecisionInt() == %v\n\n",
			ePrefix, precision, ia.GetPrecision())
	}

}

func TestIntAry_NewUint_02(t *testing.T) {
	uintNum := uint(123456)
	precision := uint(5)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "1.23456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint_03(t *testing.T) {
	uintNum := uint(0)
	precision := uint(3)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "0.000"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint_04(t *testing.T) {
	uintNum := uint(0)
	precision := uint(0)
	ia := new(IntAry).NewUint(uintNum, precision)
	eStr := "0"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUintExponent_01(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_01"
	uintNum := uint(123456)
	exponent := 3
	eStr := "123456.000"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}
}

func TestIntAry_NewUintExponent_02(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_02"
	uintNum := uint(123456)
	exponent := -3
	eStr := "123.456"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}

	return
}

func TestIntAry_NewUintExponent_03(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_03"
	uintNum := uint(123456)
	exponent := 0
	eStr := "123456"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}
}

func TestIntAry_NewUintExponent_04(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_04"
	uintNum := uint(0)
	exponent := 0
	eStr := "0"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}
}

func TestIntAry_NewUintExponent_05(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_05"
	uintNum := uint(0)
	exponent := 3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(uintNum, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}

	return
}

func TestIntAry_NewUintExponent_06(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_06"
	num := uint(0)
	exponent := -3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUintExponent(num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}


	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}
}

func TestIntAry_NewUintExponent_07(t *testing.T) {

	ePrefix := "TestIntAry_NewUintExponent_07"
	num := uint(123456)
	exponent := -3
	eStr := "-123.456"
	signValue := -1

	ia, err := new(IntAry).NewUintExponent(num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUintExponent(num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}

	return
}

func TestIntAry_NewUint32_01(t *testing.T) {
	ePrefix := "TestIntAry_NewUint32_01"
	uint32Num := uint32(123456)
	precision := uint(3)
	signValue := 1
	eStr := "123.456"

	ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

  iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n" +
			"Expected ia.GetPrecision()== '%v'\n" +
			"Instead ia.GetPrecision() == '%v'\n\n",
			ePrefix, precision, ia.GetPrecision())
		}

	return
}

func TestIntAry_NewUint32_02(t *testing.T) {

	ePrefix := "TestIntAry_NewUint32_02"

	uint32Num := uint32(123456)

	precision := uint(5)

	eStr := "1.23456"

	signValue := 1

	ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n" +
			"Expected ia.GetPrecisionInt()== '%v'\n" +
			"Instead ia.GetPrecisionInt() == '%v'\n\n",
			ePrefix, precision, ia.GetPrecision())
	}

	return
}

func TestIntAry_NewUint32_03(t *testing.T) {

	ePrefix := "TestIntAry_NewUint32_03"

	uint32Num := uint32(0)

	precision := uint(3)

	eStr := "0.000"

	signValue := 1

	ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n" +
			"Expected ia.GetPrecisionInt()== '%v'\n" +
			"Instead ia.GetPrecisionInt() == '%v'\n\n",
			ePrefix, precision, ia.GetPrecision())

	}

	return
}

func TestIntAry_NewUint32_04(t *testing.T) {
	ePrefix := "TestIntAry_NewUint32_04"
	uint32Num := uint32(0)
	precision := uint(0)
	eStr := "0"
	signValue := 1

	ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint32(uint32Num, signValue, precision)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
		return
	}


	if int(precision) != ia.GetPrecision() {
		t.Errorf("%v\n" +
			"Expected ia.GetPrecision()== '%v'\n" +
			"Instead ia.GetPrecision() == '%v'\n\n",
			ePrefix, precision, ia.GetPrecision())

	}

	return
}

func TestIntAry_NewUint32Exponent_01(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := 3
	eStr := "123456.000"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_02(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := -3
	eStr := "123.456"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_03(t *testing.T) {

	uint32Num := uint32(123456)
	exponent := 0
	eStr := "123456"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_04(t *testing.T) {

	uint32Num := uint32(0)
	exponent := 0
	eStr := "0"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_05(t *testing.T) {

	uint32Num := uint32(0)
	exponent := 3
	eStr := "0.000"

	ia := new(IntAry).NewUint32Exponent(uint32Num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint32Exponent_06(t *testing.T) {

	num := uint32(0)
	exponent := -3
	eStr := "0.000"

	ia := new(IntAry).NewUint32Exponent(num, exponent)

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}
}

func TestIntAry_NewUint64_01(t *testing.T) {
	uint64Num := uint64(123456)
	precision := uint(3)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "123.456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}

}

func TestIntAry_NewUint64_02(t *testing.T) {
	uint64Num := uint64(123456)
	precision := uint(5)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "1.23456"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64_03(t *testing.T) {
	uint64Num := uint64(0)
	precision := uint(3)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "0.000"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64_04(t *testing.T) {
	uint64Num := uint64(0)
	precision := uint(0)
	ia := new(IntAry).NewUint64(uint64Num, precision)
	eStr := "0"

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt() == %v  .   Instead ia.GetPrecisionInt() == %v", precision, ia.GetPrecision())
	}
}

func TestIntAry_NewUint64Exponent_01(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_01"

	uint64Num := uint64(123456)
	signValue := 1
	exponent := 3
	eStr := "123456.000"

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected ia.GetNumStr()== %v.\n"+
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, eStr, iaNumStr)

	}

	return
}

func TestIntAry_NewUint64Exponent_02(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_02"

	uint64Num := uint64(123456)
	exponent := -3
	signValue := 1
	eStr := "123.456"

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n"+
			"Instead ia.GetNumStr() == '%v'",
			ePrefix, eStr, iaNumStr)
	}
}

func TestIntAry_NewUint64Exponent_03(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_03"

	uint64Num := uint64(123456)
	exponent := 0
	eStr := "123456"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}


	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}
}

func TestIntAry_NewUint64Exponent_04(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_04"
	uint64Num := uint64(0)
	exponent := 0
	eStr := "0"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Error:\n"+
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == %v\n\n",
			ePrefix, eStr, iaNumStr))
	}
	return
}

func TestIntAry_NewUint64Exponent_05(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_05"
	uint64Num := uint64(0)
	exponent := 3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(uint64Num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	if eStr != iaNumStr {
		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)
	}

	return
}

func TestIntAry_NewUint64Exponent_06(t *testing.T) {

	ePrefix := "TestIntAry_NewUint64Exponent_06"
	num := uint64(0)
	exponent := -3
	eStr := "0.000"
	signValue := 1

	ia, err := new(IntAry).NewUint64Exponent(num, signValue, exponent)

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"ia, err := new(IntAry).NewUint64Exponent(num, signValue, exponent)\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}

	iaNumStr, err := ia.GetNumStr()

	if err != nil {
		t.Errorf("%v\n" +
			"Error returned by:\n"+
			"iaNumStr, err := ia.GetNumStr()\n"+
			"Error='%v'\n\n", ePrefix, err.Error())
		return
	}


	if eStr != iaNumStr {

		t.Errorf("%v\n" +
			"Expected ia.GetNumStr()== '%v'\n" +
			"Instead ia.GetNumStr() == '%v'\n\n",
			ePrefix, eStr, iaNumStr)

	}

	return
}

func TestIntAry_OptimizeIntArrayLen_01(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456000"
	ePrecision := 9
	eLen := 12

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(false)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_02(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456"
	ePrecision := 6
	eLen := 9

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_03(t *testing.T) {
	nStr1 := "00579.000000000"
	expected := "579"
	ePrecision := 0
	eLen := 3

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_04(t *testing.T) {
	nStr1 := "00000.123450000"
	expected := "0.12345"
	ePrecision := 5
	eLen := 6

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_05(t *testing.T) {
	nStr1 := "00000.000123450000"
	expected := "0.00012345"
	ePrecision := 8
	eLen := 9

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v' .", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_06(t *testing.T) {
	nStr1 := "0.0"
	expected := "0.0"
	ePrecision := 1
	eLen := 2

	ia := new(IntAry).New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecisionInt()= '%v'. Instead, ia.GetPrecisionInt()= '%v' .", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}
