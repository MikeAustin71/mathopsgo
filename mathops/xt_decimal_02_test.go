package mathops

import (
  "testing"
)

func TestDecimal_MakeDecimalFromIntAry_01(t *testing.T) {

  ePrefix := "TestDecimal_MakeDecimalFromIntAry_01"

  expectedNumberStr := "982.123456"

  expectedPrecisionUint := uint(6)

  expectedSignVal := 1

  intArray, err := new(IntAry).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArray, err := new(IntAry).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = intArray.IsValid("Validating intArray")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intArray.IsValid('Validating intArray')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArrayNumberStr, err := intArray.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayNumberStr, err := intArray.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != intArrayNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != intArrayNumberStr\n"+
      "Expected intArrayNumberStr = '%v'\n"+
      "  Actual intArrayNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intArrayNumberStr)

    return
  }

  decNum1 := new(Decimal).New()

  decNum2, err := decNum1.MakeDecimalFromIntAry(&intArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := decNum1.MakeDecimalFromIntAry(&intArray)\n"+
      "intArray= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intArrayNumberStr, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2Str, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2Str, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2PrecisionUint, err := decNum2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2PrecisionUint, err := decNum2.GetPrecisionUint()\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNum2Str, err.Error())
    return
  }

  decNum2SignValue, err := decNum2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2SignValue, err := decNum2.GetSign()\n"+
      "decNum2= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNum2Str, err.Error())
    return
  }

  if expectedNumberStr != decNum2Str {

    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != decNum2Str\n"+
      "Expected decNum2Str = '%v'\n"+
      "  Actual decNum2Str = '%v'\n\n",
      ePrefix, expectedNumberStr, decNum2Str)

    return
  }

  if expectedPrecisionUint != decNum2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != decNum2PrecisionUint\n"+
      "Expected decNum2PrecisionUint = '%v'\n"+
      "  Actual decNum2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNum2PrecisionUint)

    return
  }

  if expectedSignVal != decNum2SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != decNum2SignValue\n"+
      "Expected decNum2SignValue = '%v'\n"+
      "  Actual decNum2SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNum2SignValue)

    return
  }

  return
}

func TestDecimal_MakeDecimalFromIntAry_02(t *testing.T) {

  ePrefix := "TestDecimal_MakeDecimalFromIntAry_02"

  expectedNumberStr := "-982.123456"

  expectedPrecisionUint := uint(6)

  expectedSignVal := -1

  intArray, err := new(IntAry).NewNumStr(expectedNumberStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArray, err := new(IntAry).NewNumStr(expectedNumberStr)\n"+
      "expectedNumberStr= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumberStr, err.Error())
    return
  }

  err = intArray.IsValid("Validating intArray")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = intArray.IsValid('Validating intArray')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  intArrayNumberStr, err := intArray.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "intArrayNumberStr, err := intArray.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumberStr != intArrayNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != intArrayNumberStr\n"+
      "Expected intArrayNumberStr = '%v'\n"+
      "  Actual intArrayNumberStr = '%v'\n\n",
      ePrefix, expectedNumberStr, intArrayNumberStr)

    return
  }

  decNum1 := new(Decimal).New()

  decNum2, err := decNum1.MakeDecimalFromIntAry(&intArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2, err := decNum1.MakeDecimalFromIntAry(&intArray)\n"+
      "intArray= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, intArrayNumberStr, err.Error())
    return
  }

  err = decNum2.IsValid("Validating decNum2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNum2.IsValid('Validating decNum2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2Str, err := decNum2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2Str, err := decNum2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNum2PrecisionUint, err := decNum2.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2PrecisionUint, err := decNum2.GetPrecisionUint()\n"+
      "decNum2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNum2Str, err.Error())
    return
  }

  decNum2SignValue, err := decNum2.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNum2SignValue, err := decNum2.GetSign()\n"+
      "decNum2= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNum2Str, err.Error())
    return
  }

  if expectedNumberStr != decNum2Str {

    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedNumberStr != decNum2Str\n"+
      "Expected decNum2Str = '%v'\n"+
      "  Actual decNum2Str = '%v'\n\n",
      ePrefix, expectedNumberStr, decNum2Str)

    return
  }

  if expectedPrecisionUint != decNum2PrecisionUint {
    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedPrecisionUint != decNum2PrecisionUint\n"+
      "Expected decNum2PrecisionUint = '%v'\n"+
      "  Actual decNum2PrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, decNum2PrecisionUint)

    return
  }

  if expectedSignVal != decNum2SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedSignVal != decNum2SignValue\n"+
      "Expected decNum2SignValue = '%v'\n"+
      "  Actual decNum2SignValue = '%v'\n\n",
      ePrefix, expectedSignVal, decNum2SignValue)

    return
  }

  return
}

func TestDecimal_Multiply_01(t *testing.T) {

  ePrefix := "TestDecimal_Multiply_01"

  // product = multiplier x multiplicand

  multiplierNumStr1 := "575.63"

  multiplicandNumStr2 := "2014.123"

  expectedProductNumberStr := "1159389.62249"

  expectedProductPrecisionUint := uint(5)

  expectedProductSignVal := 1

  decNumMultiplier1 := new(Decimal).New()

  err := decNumMultiplier1.SetNumStr(multiplierNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumMultiplier1.SetNumStr(multiplierNumStr1)\n"+
      "multiplierNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierNumStr1, err.Error())
    return
  }

  err = decNumMultiplier1.IsValid("Validating decNumMultiplier1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplier1.IsValid('Validating decNumMultiplier1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierNumStr1 != decNumMultiplier1NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplierNumStr1 != decNumMultiplier1NumberStr\n"+
      "Expected decNumMultiplier1NumberStr = '%v'\n"+
      "  Actual decNumMultiplier1NumberStr = '%v'\n\n",
      ePrefix, multiplierNumStr1, decNumMultiplier1NumberStr)

    return
  }

  decNumMultiplicand2 := new(Decimal).New()

  err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)\n"+
      "multiplicandNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplicandNumStr2, err.Error())
    return
  }

  err = decNumMultiplicand2.IsValid("Validating decNumMultiplicand2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.IsValid('Validating decNumMultiplicand2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplicandNumStr2 != decNumMultiplicand2NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplicandNumStr2 != decNumMultiplicand2NumberStr\n"+
      "Expected decNumMultiplicand2NumberStr = '%v'\n"+
      "  Actual decNumMultiplicand2NumberStr = '%v'\n\n",
      ePrefix, multiplicandNumStr2, decNumMultiplicand2NumberStr)

    return
  }

  decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)\n"+
      "decNumMultiplier1= '%v'\n"+
      "decNumMultiplicand2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      multiplierNumStr1,
      multiplicandNumStr2,
      err.Error())

    return
  }

  err = decNumProduct3.IsValid("Validating decNumProduct3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct3.IsValid('Validating decNumProduct3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()\n"+
      "decNumProduct3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  decNumProduct3SignValue, err := decNumProduct3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3SignValue, err := decNumProduct3.GetSign()\n"+
      "decNumProduct3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  if expectedProductNumberStr != decNumProduct3NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because expectedProductNumberStr != decNumProduct3NumberStr\n"+
      "Expected decNumProduct3NumberStr = '%v'\n"+
      "  Actual decNumProduct3NumberStr = '%v'\n\n",
      ePrefix, expectedProductNumberStr, decNumProduct3NumberStr)

    return
  }

  if expectedProductPrecisionUint != decNumProduct3PrecisionUint {

    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedProductPrecisionUint != decNumProduct3PrecisionUint\n"+
      "Expected decNumProduct3PrecisionUint = '%v'\n"+
      "  Actual decNumProduct3PrecisionUint = '%v'\n\n",
      ePrefix, expectedProductPrecisionUint, decNumProduct3PrecisionUint)

    return
  }

  if expectedProductSignVal != decNumProduct3SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedProductSignVal != decNumProduct3SignValue\n"+
      "Expected decNumProduct3SignValue = '%v'\n"+
      "  Actual decNumProduct3SignValue = '%v'\n\n",
      ePrefix, expectedProductSignVal, decNumProduct3SignValue)

    return
  }

  return
}

func TestDecimal_Multiply_02(t *testing.T) {

  ePrefix := "TestDecimal_Multiply_02"

  // product = multiplier x multiplicand

  multiplierNumStr1 := "-575.63"

  multiplicandNumStr2 := "2014.123"

  expectedProductNumberStr := "-1159389.62249"

  expectedProductPrecisionUint := uint(5)

  expectedProductSignVal := -1

  decNumMultiplier1 := new(Decimal).New()

  err := decNumMultiplier1.SetNumStr(multiplierNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumMultiplier1.SetNumStr(multiplierNumStr1)\n"+
      "multiplierNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierNumStr1, err.Error())
    return
  }

  err = decNumMultiplier1.IsValid("Validating decNumMultiplier1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplier1.IsValid('Validating decNumMultiplier1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierNumStr1 != decNumMultiplier1NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplierNumStr1 != decNumMultiplier1NumberStr\n"+
      "Expected decNumMultiplier1NumberStr = '%v'\n"+
      "  Actual decNumMultiplier1NumberStr = '%v'\n\n",
      ePrefix, multiplierNumStr1, decNumMultiplier1NumberStr)

    return
  }

  decNumMultiplicand2 := new(Decimal).New()

  err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)\n"+
      "multiplicandNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplicandNumStr2, err.Error())
    return
  }

  err = decNumMultiplicand2.IsValid("Validating decNumMultiplicand2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.IsValid('Validating decNumMultiplicand2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplicandNumStr2 != decNumMultiplicand2NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplicandNumStr2 != decNumMultiplicand2NumberStr\n"+
      "Expected decNumMultiplicand2NumberStr = '%v'\n"+
      "  Actual decNumMultiplicand2NumberStr = '%v'\n\n",
      ePrefix, multiplicandNumStr2, decNumMultiplicand2NumberStr)

    return
  }

  decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)\n"+
      "decNumMultiplier1= '%v'\n"+
      "decNumMultiplicand2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      multiplierNumStr1,
      multiplicandNumStr2,
      err.Error())

    return
  }

  err = decNumProduct3.IsValid("Validating decNumProduct3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct3.IsValid('Validating decNumProduct3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()\n"+
      "decNumProduct3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  decNumProduct3SignValue, err := decNumProduct3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3SignValue, err := decNumProduct3.GetSign()\n"+
      "decNumProduct3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  if expectedProductNumberStr != decNumProduct3NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because expectedProductNumberStr != decNumProduct3NumberStr\n"+
      "Expected decNumProduct3NumberStr = '%v'\n"+
      "  Actual decNumProduct3NumberStr = '%v'\n\n",
      ePrefix, expectedProductNumberStr, decNumProduct3NumberStr)

    return
  }

  if expectedProductPrecisionUint != decNumProduct3PrecisionUint {

    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedProductPrecisionUint != decNumProduct3PrecisionUint\n"+
      "Expected decNumProduct3PrecisionUint = '%v'\n"+
      "  Actual decNumProduct3PrecisionUint = '%v'\n\n",
      ePrefix, expectedProductPrecisionUint, decNumProduct3PrecisionUint)

    return
  }

  if expectedProductSignVal != decNumProduct3SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedProductSignVal != decNumProduct3SignValue\n"+
      "Expected decNumProduct3SignValue = '%v'\n"+
      "  Actual decNumProduct3SignValue = '%v'\n\n",
      ePrefix, expectedProductSignVal, decNumProduct3SignValue)

    return
  }

  return
}

func TestDecimal_Multiply_03(t *testing.T) {

  ePrefix := "TestDecimal_Multiply_03"

  // product = multiplier x multiplicand

  multiplierNumStr1 := "-575.63"

  multiplicandNumStr2 := "-2014.123"

  expectedProductNumberStr := "1159389.62249"

  expectedProductPrecisionUint := uint(5)

  expectedProductSignVal := 1

  decNumMultiplier1 := new(Decimal).New()

  err := decNumMultiplier1.SetNumStr(multiplierNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumMultiplier1.SetNumStr(multiplierNumStr1)\n"+
      "multiplierNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierNumStr1, err.Error())
    return
  }

  err = decNumMultiplier1.IsValid("Validating decNumMultiplier1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplier1.IsValid('Validating decNumMultiplier1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierNumStr1 != decNumMultiplier1NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplierNumStr1 != decNumMultiplier1NumberStr\n"+
      "Expected decNumMultiplier1NumberStr = '%v'\n"+
      "  Actual decNumMultiplier1NumberStr = '%v'\n\n",
      ePrefix, multiplierNumStr1, decNumMultiplier1NumberStr)

    return
  }

  decNumMultiplicand2 := new(Decimal).New()

  err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)\n"+
      "multiplicandNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplicandNumStr2, err.Error())
    return
  }

  err = decNumMultiplicand2.IsValid("Validating decNumMultiplicand2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.IsValid('Validating decNumMultiplicand2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplicandNumStr2 != decNumMultiplicand2NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplicandNumStr2 != decNumMultiplicand2NumberStr\n"+
      "Expected decNumMultiplicand2NumberStr = '%v'\n"+
      "  Actual decNumMultiplicand2NumberStr = '%v'\n\n",
      ePrefix, multiplicandNumStr2, decNumMultiplicand2NumberStr)

    return
  }

  decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)\n"+
      "decNumMultiplier1= '%v'\n"+
      "decNumMultiplicand2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      multiplierNumStr1,
      multiplicandNumStr2,
      err.Error())

    return
  }

  err = decNumProduct3.IsValid("Validating decNumProduct3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct3.IsValid('Validating decNumProduct3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()\n"+
      "decNumProduct3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  decNumProduct3SignValue, err := decNumProduct3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3SignValue, err := decNumProduct3.GetSign()\n"+
      "decNumProduct3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  if expectedProductNumberStr != decNumProduct3NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because expectedProductNumberStr != decNumProduct3NumberStr\n"+
      "Expected decNumProduct3NumberStr = '%v'\n"+
      "  Actual decNumProduct3NumberStr = '%v'\n\n",
      ePrefix, expectedProductNumberStr, decNumProduct3NumberStr)

    return
  }

  if expectedProductPrecisionUint != decNumProduct3PrecisionUint {

    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedProductPrecisionUint != decNumProduct3PrecisionUint\n"+
      "Expected decNumProduct3PrecisionUint = '%v'\n"+
      "  Actual decNumProduct3PrecisionUint = '%v'\n\n",
      ePrefix, expectedProductPrecisionUint, decNumProduct3PrecisionUint)

    return
  }

  if expectedProductSignVal != decNumProduct3SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedProductSignVal != decNumProduct3SignValue\n"+
      "Expected decNumProduct3SignValue = '%v'\n"+
      "  Actual decNumProduct3SignValue = '%v'\n\n",
      ePrefix, expectedProductSignVal, decNumProduct3SignValue)

    return
  }

  return
}

func TestDecimal_Multiply_04(t *testing.T) {

  ePrefix := "TestDecimal_Multiply_04"

  // product = multiplier x multiplicand

  multiplierNumStr1 := "0"

  multiplicandNumStr2 := "-2014.123"

  expectedProductNumberStr := "0"

  expectedProductPrecisionUint := uint(0)

  expectedProductSignVal := 1

  decNumMultiplier1 := new(Decimal).New()

  err := decNumMultiplier1.SetNumStr(multiplierNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := decNumMultiplier1.SetNumStr(multiplierNumStr1)\n"+
      "multiplierNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierNumStr1, err.Error())
    return
  }

  err = decNumMultiplier1.IsValid("Validating decNumMultiplier1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplier1.IsValid('Validating decNumMultiplier1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierNumStr1 != decNumMultiplier1NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplierNumStr1 != decNumMultiplier1NumberStr\n"+
      "Expected decNumMultiplier1NumberStr = '%v'\n"+
      "  Actual decNumMultiplier1NumberStr = '%v'\n\n",
      ePrefix, multiplierNumStr1, decNumMultiplier1NumberStr)

    return
  }

  decNumMultiplicand2 := new(Decimal).New()

  err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.SetNumStr(multiplicandNumStr2)\n"+
      "multiplicandNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplicandNumStr2, err.Error())
    return
  }

  err = decNumMultiplicand2.IsValid("Validating decNumMultiplicand2")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplicand2.IsValid('Validating decNumMultiplicand2')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplicand2NumberStr, err := decNumMultiplicand2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplicandNumStr2 != decNumMultiplicand2NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplicandNumStr2 != decNumMultiplicand2NumberStr\n"+
      "Expected decNumMultiplicand2NumberStr = '%v'\n"+
      "  Actual decNumMultiplicand2NumberStr = '%v'\n\n",
      ePrefix, multiplicandNumStr2, decNumMultiplicand2NumberStr)

    return
  }

  decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3, err := decNumMultiplier1.Multiply(decNumMultiplicand2)\n"+
      "decNumMultiplier1= '%v'\n"+
      "decNumMultiplicand2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      multiplierNumStr1,
      multiplicandNumStr2,
      err.Error())

    return
  }

  err = decNumProduct3.IsValid("Validating decNumProduct3")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct3.IsValid('Validating decNumProduct3')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3NumberStr, err := decNumProduct3.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3PrecisionUint, err := decNumProduct3.GetPrecisionUint()\n"+
      "decNumProduct3= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  decNumProduct3SignValue, err := decNumProduct3.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct3SignValue, err := decNumProduct3.GetSign()\n"+
      "decNumProduct3= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumProduct3NumberStr, err.Error())
    return
  }

  if expectedProductNumberStr != decNumProduct3NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because expectedProductNumberStr != decNumProduct3NumberStr\n"+
      "Expected decNumProduct3NumberStr = '%v'\n"+
      "  Actual decNumProduct3NumberStr = '%v'\n\n",
      ePrefix, expectedProductNumberStr, decNumProduct3NumberStr)

    return
  }

  if expectedProductPrecisionUint != decNumProduct3PrecisionUint {

    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedProductPrecisionUint != decNumProduct3PrecisionUint\n"+
      "Expected decNumProduct3PrecisionUint = '%v'\n"+
      "  Actual decNumProduct3PrecisionUint = '%v'\n\n",
      ePrefix, expectedProductPrecisionUint, decNumProduct3PrecisionUint)

    return
  }

  if expectedProductSignVal != decNumProduct3SignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedProductSignVal != decNumProduct3SignValue\n"+
      "Expected decNumProduct3SignValue = '%v'\n"+
      "  Actual decNumProduct3SignValue = '%v'\n\n",
      ePrefix, expectedProductSignVal, decNumProduct3SignValue)

    return
  }

  return
}

func TestDecimal_Multiply_05(t *testing.T) {

  ePrefix := "TestDecimal_Multiply_05"

  // product = multiplier x multiplicand

  multiplierNumStr1 := "3"

  multiplicandNumStr2 := "1"

  expectedProductNumberStr := "81"

  expectedProductPrecisionUint := uint(0)

  expectedProductSignVal := 1

  decNumMultiplier1, err := new(Decimal).NewNumStr(multiplierNumStr1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1, err := new(Decimal).NewNumStr(multiplierNumStr1)\n"+
      "multiplierNumStr1= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierNumStr1, err.Error())
    return
  }

  err = decNumMultiplier1.IsValid("Validating decNumMultiplier1")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumMultiplier1.IsValid('Validating decNumMultiplier1')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumMultiplier1NumberStr, err := decNumMultiplier1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierNumStr1 != decNumMultiplier1NumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because multiplierNumStr1 != decNumMultiplier1NumberStr\n"+
      "Expected decNumMultiplier1NumberStr = '%v'\n"+
      "  Actual decNumMultiplier1NumberStr = '%v'\n\n",
      ePrefix, multiplierNumStr1, decNumMultiplier1NumberStr)

    return
  }

  decNumProduct, err := new(Decimal).NewNumStr(multiplicandNumStr2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProduct, err := new(Decimal).NewNumStr(multiplicandNumStr2)\n"+
      "multiplicandNumStr2= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplicandNumStr2, err.Error())
    return
  }

  err = decNumProduct.IsValid("Validating Initial decNumProduct")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct.IsValid('Validating Initial decNumProduct')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProductNumberStr, err := decNumProduct.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProductNumberStr, err := decNumProduct.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  for i := 0; i < 4; i++ {

    decNumProduct, err = decNumProduct.Multiply(decNumMultiplier1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumProduct, err =\n"+
        "  decNumProduct.Multiply(decNumMultiplier1)\n"+
        "decNumMultiplier1= '%v'\n"+
        "decNumProduct= '%v'\n"+
        "Cycle No = '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix,
        decNumMultiplier1NumberStr,
        decNumProductNumberStr,
        i,
        err.Error())

      return
    }

    decNumProductNumberStr, err = decNumProduct.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decNumProductNumberStr, err := decNumProduct.GetNumStr()\n"+
        "Multiply Loop Cycle No = '%v'\n"+
        "Error= '%v'\n\n", ePrefix, i, err.Error())
      return
    }

  }

  err = decNumProduct.IsValid("Validating After Multiplication decNumProduct")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = decNumProduct.IsValid('Validating After Multiplication decNumProduct')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  decNumProductPrecisionUint, err := decNumProduct.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProductPrecisionUint, err := decNumProduct.GetPrecisionUint()\n"+
      "decNumProduct= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, decNumProductNumberStr, err.Error())
    return
  }

  decNumProductSignValue, err := decNumProduct.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "decNumProductSignValue, err := decNumProduct.GetSign()\n"+
      "decNumProduct= '%v\n"+
      "Error= '%v'\n\n", ePrefix, decNumProductNumberStr, err.Error())
    return
  }

  if expectedProductNumberStr != decNumProductNumberStr {

    t.Errorf("%v\n"+
      "Error: UnexpectedProduct Result!\n"+
      "Because expectedProductNumberStr != decNumProductNumberStr\n"+
      "Expected decNumProductNumberStr = '%v'\n"+
      "  Actual decNumProductNumberStr = '%v'\n\n",
      ePrefix, expectedProductNumberStr, decNumProductNumberStr)

    return
  }

  if expectedProductPrecisionUint != decNumProductPrecisionUint {

    t.Errorf("%v\n"+
      "Error: Precision Values ARE NOT Equal!\n"+
      "Because expectedProductPrecisionUint != decNumProductPrecisionUint\n"+
      "Expected decNumProductPrecisionUint = '%v'\n"+
      "  Actual decNumProductPrecisionUint = '%v'\n\n",
      ePrefix, expectedProductPrecisionUint, decNumProductPrecisionUint)

    return
  }

  if expectedProductSignVal != decNumProductSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values ARE NOT Equal!\n"+
      "Because expectedProductSignVal != decNumProductSignValue\n"+
      "Expected decNumProductSignValue = '%v'\n"+
      "  Actual decNumProductSignValue = '%v'\n\n",
      ePrefix, expectedProductSignVal, decNumProductSignValue)

    return
  }

  return
}
