package mathops

import (
  "fmt"
  "testing"
)

func TestBigIntMathMultiply_MultiplyIntAry_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_01"

  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875.94572"

  expectedSignValue := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier='%v'\n"+
      "iaMultiplicand='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      "multiplierIntAry, multiplicandIntAry)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      multiplicandIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
      "expectedIntAry='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedIntAryNumStr, err.Error())
    return
  }

  if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
      "Expected resultBigIntNum = '%v'\n"+
      "  Actual resultBigIntNum = '%v'\n\n",
      ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != result.sign\n"+
      "Expected result.sign = '%v'\n"+
      "  Actual result.sign = '%v'\n\n",
      ePrefix, expectedSignValue, result.sign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAry_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_02"

  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedNumStr := "30987680500513189125.14259702468435"

  expectedSignValue := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier='%v'\n"+
      "iaMultiplicand='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      "multiplierIntAry, multiplicandIntAry)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      multiplicandIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedIntAryNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
      "expectedIntAry='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedIntAryNumStr, err.Error())
    return
  }

  if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
      "Expected resultBigIntNum = '%v'\n"+
      "  Actual resultBigIntNum = '%v'\n\n",
      ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != result.sign\n"+
      "Expected result.sign = '%v'\n"+
      "  Actual result.sign = '%v'\n\n",
      ePrefix, expectedSignValue, result.sign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAry_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_03"

  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedNumStr := "-30987680500513189125.14259702468435"

  expectedSignValue := -1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier='%v'\n"+
      "iaMultiplicand='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      "multiplierIntAry, multiplicandIntAry)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      multiplicandIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
      "expectedIntAry='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedIntAryNumStr, err.Error())
    return
  }

  if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
      "Expected resultBigIntNum = '%v'\n"+
      "  Actual resultBigIntNum = '%v'\n\n",
      ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != result.sign\n"+
      "Expected result.sign = '%v'\n"+
      "  Actual result.sign = '%v'\n\n",
      ePrefix, expectedSignValue, result.sign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAry_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_04"

  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedNumStr := "22197234145.3632"

  expectedSignValue := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier='%v'\n"+
      "iaMultiplicand='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      "multiplierIntAry, multiplicandIntAry)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      multiplicandIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedIntAryNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
      "expectedIntAry='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedIntAryNumStr, err.Error())
    return
  }

  if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
      "Expected resultBigIntNum = '%v'\n"+
      "  Actual resultBigIntNum = '%v'\n\n",
      ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != result.sign\n"+
      "Expected result.sign = '%v'\n"+
      "  Actual result.sign = '%v'\n\n",
      ePrefix, expectedSignValue, result.sign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  iaResultActualBigInt, err := iaResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultActualBigInt, err := iaResult.GetBigInt()\n"+
      "iaResult= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, iaResultNumStr, err.Error())
    return
  }

  if resultActualBigInt.Cmp(iaResultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because resultActualBigInt.Cmp(iaResultActualBigInt) != 0\n"+
      "Expected iaResultActualBigInt = '%v'\n"+
      "  Actual iaResultActualBigInt = '%v'\n\n",
      ePrefix, resultActualBigInt.Text(10), iaResultActualBigInt.Text(10))

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAry_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_05"

  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0
  expectedNumStr := "0"

  expectedSignValue := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier='%v'\n"+
      "iaMultiplicand='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "esult, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      "multiplierIntAry, multiplicandIntAry)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      multiplicandIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  expectedIntAryBigInt, err := expectedIntAry.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryBigInt, err := expectedIntAry.GetBigInt()\n"+
      "expectedIntAry='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedIntAryNumStr, err.Error())
    return
  }

  if expectedIntAryBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because  expectedIntAryBigInt.Cmp(result.bigInt) != 0\n"+
      "Expected resultBigIntNum = '%v'\n"+
      "  Actual resultBigIntNum = '%v'\n\n",
      ePrefix, expectedIntAryBigInt.Text(10), result.bigInt.Text(10))

    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != result.sign\n"+
      "Expected result.sign = '%v'\n"+
      "  Actual result.sign = '%v'\n\n",
      ePrefix, expectedSignValue, result.sign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  iaResultActualBigInt, err := iaResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultActualBigInt, err := iaResult.GetBigInt()\n"+
      "iaResult= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, iaResultNumStr, err.Error())
    return
  }

  if resultActualBigInt.Cmp(iaResultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because resultActualBigInt.Cmp(iaResultActualBigInt) != 0\n"+
      "Expected iaResultActualBigInt = '%v'\n"+
      "  Actual iaResultActualBigInt = '%v'\n\n",
      ePrefix, resultActualBigInt.Text(10), iaResultActualBigInt.Text(10))

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAry_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAry_06"

  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedNumStr := "22197234145,3632"

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  multiplierIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = multiplierIntAry.IsValid("Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAry, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = multiplicandIntAry.IsValid("Validating multiplicandIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplicandIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandIntAryNumStr, err := multiplicandIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(\n"+
      "  &iaMultiplier, &iaMultiplicand\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, err.Error())
    return
  }

  err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAry, err := new(IntAry).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAry, err := new(IntAry).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\n"+
      "expectedNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedIntAry.IsValid("Validating expectedIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedIntAryNumStr, err := expectedIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedIntAryNumStr, err := expectedIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAry(multiplierIntAry, multiplicandIntAry, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyIntAry(\n"+
      " multiplierIntAry, multiplicandIntAry, expectedNumSeps)\n"+
      "multiplierIntAry= '%v'\n"+
      "multiplicandIntAry= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierIntAryNumStr,
      multiplicandIntAryNumStr, err.Error())
    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedIntAryNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because!!!!\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, expectedIntAryNumStr, resultNumStr)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !expectedNumSeps.Equal(actualNumSeps)\n"+
      "Expected actualNumSeps = '%v'\n"+
      "  Actual actualNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), actualNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_01"

  var err error

  // multiplier = 2
  multiplierStr := "2"
  // multiplicandStrs
  multiplicandStrs := []string{
    "2",
    "2",
    "2",
    "2",
    "2",
    "2",
  }

  // product = 128
  expectedBigINumStr := "128"

  expectedBigINumSign := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_02"

  var err error

  // multiplier = 37.9876
  multiplierStr := "37.9876"

  // multiplicandStrs
  multiplicandStrs := []string{
    "-27.9",
    "48.123456",
    "59.48721",
    "-3",
    "19.1",
    "69",
  }

  // product = 11995826664.26376575446779648
  expectedBigINumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_03"

  var err error

  // multiplier = 10.1
  multiplierStr := "10.1"

  // multiplicandStrs
  multiplicandStrs := []string{
    "2",
    "5.8",
    "68.7",
    "3.1234567",
    "8.0",
    "11",
  }

  // product = 2212352.1767579232
  expectedBigINumStr := "2212352.1767579232"

  expectedBigINumSign := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "Validating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "Validating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = multiplicandIntArray[i].IsValid(ePrefix + fmt.Sprintf(" Validating multiplicandIntArray[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = multiplicandIntArray[%d].IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "iaResult= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], iaResultNumStr, err.Error())
      return
    }

    err = iaResult.IsValid(ePrefix + fmt.Sprintf(" Validating iaResult[%d]", i))

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult[%d].IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err := iaResult[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

  } // End of Loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  err = result.IsValid(ePrefix + "Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_04"

  var err error

  // multiplier = -5.123456
  multiplierStr := "-5.123456"

  // multiplicandStrs
  multiplicandStrs := []string{
    "1.879",
    "3.824",
    "21.756",
    "2.1234567",
    "6",
    "2",
  }

  // product = -20408.5138429311978576052224
  expectedBigINumStr := "-20408.5138429311978576052224"

  expectedBigINumSign := -1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryArray_05"

  var err error

  // multiplier = 10.1
  multiplierStr := "10.1"
  // multiplicandStrs
  multiplicandStrs := []string{
    "2",
    "5.8",
    "68.7",
    "3.1234567",
    "8.0",
    "11",
  }

  // product = 2212352.1767579232
  expectedNumStr := "2212352,1767579232"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  multiplierIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenArray)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps = '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  iaResultNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResultNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
      "expectedNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyIntAryArray(multiplierIntAry, multiplicandIntArray, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyIntAryArray(\n"+
      "  multiplierIntAry, multiplicandIntArray, expectedNumSeps))\n"+
      "multiplierIntAry= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierIntAryNumStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !expectedNumSeps.Equal(resultNumSeps)\n"+
      "Expected 'result' Numeric Separators = '%v'\n"+
      "  Actual 'result' Numeric Separators = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_01"

  var err error

  // multiplier = 2
  multiplierStr := "2"

  // multiplicandStrs
  multiplicandStrs := []string{
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "2",
    "4",
    "6",
    "8",
    "10",
    "12",
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of Loop

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(
    multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultIaArray, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStrs[j])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStr(expectedNumStrs[%d])\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_02"

  var err error

  // multiplier = 8
  multiplierStr := "8"

  // multiplicandStrs
  multiplicandStrs := []string{
    "100.1",
    "-26",
    "3.924",
    "8",
    "5297.123",
    "-4.896",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "800.8",
    "-208",
    "31.392",
    "64",
    "42376.984",
    "-39.168",
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStrs[j])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStr(expectedNumStrs[%d])\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_03"

  var err error

  // multiplier = -31.2
  multiplierStr := "-31.2"

  // multiplicandStrs
  multiplicandStrs := []string{
    "100.1",
    "-26",
    "3.924",
    "8",
    "5297.123",
    "-4.896",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "-3123.12",
    "811.2",
    "-122.4288",
    "-249.6",
    "-165270.2376",
    "152.7552",
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStrs[j])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStr(expectedNumStrs[%d])\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_04"

  var err error

  // multiplier = 283
  multiplierStr := "283"

  // multiplicandStrs
  multiplicandStrs := []string{
    "0",
    "-26",
    "0",
    "8",
    "5297.123",
    "0",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "0",
    "-7358",
    "0",
    "2264",
    "1499085.809",
    "0",
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStrs[j])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStr(expectedNumStrs[%d])\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_05"

  var err error

  // multiplier = 0
  multiplierStr := "0"

  // multiplicandStrs
  multiplicandStrs := []string{
    "5",
    "-26",
    "9",
    "8",
    "5297.123",
    "37",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray)\n"+
      "multiplierIntAry= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierIntAryNumStr,
      err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStr(expectedNumStrs[j])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStr(expectedNumStrs[%d])\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}

func TestBigIntMathMultiply_MultiplyIntAryOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntAryOutputToArray_06"

  var err error

  // multiplier = 8
  multiplierStr := "8"

  // multiplicandStrs
  multiplicandStrs := []string{
    "100.1",
    "-26",
    "3.924",
    "8",
    "5297.123",
    "-4.896",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "800,8",
    "-208",
    "31,392",
    "64",
    "42376,984",
    "-39,168",
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  multiplierIntAry, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr='%v'\n"+
      "usaNumSeps='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  lenExpectedNumStrsAry := len(expectedNumStrs)

  if lenExpectedNumStrsAry != lenMultiplicandStrsAry {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because lenExpectedNumStrsAry != lenMultiplicandStrsAry\n"+
      "Expected Length of ExpectedNumStrs Array = '%v'\n"+
      "  Actual Length of ExpectedNumStrs Array = '%v'\n\n",
      ePrefix, lenExpectedNumStrsAry, lenMultiplicandStrsAry)

    return
  }

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

  }

  resultIaArray, err := new(BigIntMathMultiply).MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyIntAryOutputToArray(multiplierIntAry, multiplicandIntArray, expectedNumSeps)\n"+
      "multiplierIntAry= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierIntAryNumStr, expectedNumSeps.String(), err.Error())

    return
  }

  var resultNumStr, expectedNumStr string

  var resultBigINum BigIntNum

  for j := 0; j < lenMultiplicandStrsAry; j++ {

    expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStrs[j], &expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err := new(BigIntNum).\n"+
        "  NewNumStrWithNumSeps(expectedNumStrs[%d], &expectedNumSeps)\n"+
        "expectedNumStrs[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n", ePrefix, j, j, expectedNumStrs[j], expectedNumSeps.String(), err.Error())
      return
    }

    expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    resultBigINum, err = resultIaArray[j].GetBigIntNum()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINum, err :=resultBigINum, err =\n"+
        "  resultIaArray[%d].GetBigIntNum())\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    resultBigINumNumberStr, err := resultBigINum.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultBigINumNumberStr, err := resultBigINum.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "expectedBigINumEqualsResult, err :=\n"+
        "  expectedBigINum.Equal(resultBigINum)\n"+
        "expectedBigINum= '%v'"+
        "resultBigINum= '%v'"+
        "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr,
        resultBigINumNumberStr, err.Error())
      return
    }

    if !expectedBigINumEqualsResult {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedBigINumEqualsResult = 'false'\n"+
        "Expected result = '%v'\n"+
        "  Actual result = '%v'\n\n",
        ePrefix, expectedBigIntNumberStr, resultBigINumNumberStr)

      return
    }

    expectedNumStr = expectedNumStrs[j]

    resultNumStr, err = resultIaArray[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = resultIaArray[j].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    if expectedNumStr != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Because expectedNumStr != resultNumStr\n"+
        "Expected resultNumStr = '%v'\n"+
        "  Actual resultNumStr = '%v'\n\n",
        ePrefix, expectedNumStr, resultNumStr)

      return
    }

  }

  return
}
