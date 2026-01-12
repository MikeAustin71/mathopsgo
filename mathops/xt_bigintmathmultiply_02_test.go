package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_01"

  // multiplicand = 23.321
  multiplicandBInt := big.NewInt(23321)
  multiplicandPrecision := uint(3)
  exponent := uint(5)
  expectedResult := "746.272"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\nError returned by:\n"+
      "new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum()\n"+""+
      "Error= %v\n\n", ePrefix, err.Error())
    return
  }

  numStrResult, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result.GetNumStr()\n"+
      "Error= %v\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != numStrResult {
    t.Errorf("%v\n"+
      "Error: Expected result.GetNumStr() result='%v'.\n"+
      "Instead, result='%v'.\n\n",
      ePrefix,
      expectedResult, numStrResult)
    return
  }

}

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_02"

  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := uint(0)
  exponent := uint(10)
  expectedResult := "8192"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(\n"+
      "  multiplicandBInt, multiplicandPrecision, exponent)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'.\n\n",
      ePrefix, expectedResult, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_03"

  // multiplicand = 9.871234
  multiplicandBInt := big.NewInt(9871234)
  multiplicandPrecision := uint(6)
  exponent := uint(1)
  expectedResult := "19.742468"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(\n"+
      "  multiplicandBInt, multiplicandPrecision, exponent)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplicandBInt.Text(10),
      multiplicandPrecision,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'.\n\n",
      ePrefix, expectedResult, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_04"

  // multiplicand = -9.871234
  multiplicandBInt := big.NewInt(-9871234)
  multiplicandPrecision := uint(6)
  exponent := uint(3)
  expectedResult := "-78.969872"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(\n"+
      "  multiplicandBInt, multiplicandPrecision, exponent)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplicandBInt.Text(10),
      multiplicandPrecision,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'.\n\n",
      ePrefix, expectedResult, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_05"

  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := uint(0)
  exponent := uint(0)
  expectedResult := "8"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(\n"+
      "     multiplicandBInt, multiplicandPrecision, exponent)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplicandBInt.Text(10),
      multiplicandPrecision,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'.\n\n",
      ePrefix, expectedResult, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntByTwoToPower_06"

  // multiplicand = 0
  multiplicandBInt := big.NewInt(0)
  multiplicandPrecision := uint(0)
  exponent := uint(4)
  expectedResult := "0"

  result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(
    multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigInt2ToPowerBigIntNum(\n"+
      "     multiplicandBInt, multiplicandPrecision, exponent)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplicandBInt.Text(10),
      multiplicandPrecision,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\nInstead, result='%v'.\n\n",
      ePrefix, expectedResult, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_01"

  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBigNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBigNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandBigNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBigNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBigNumStr,
      multiplicandBigNumStr,
      err.Error())

    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_02"

  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedBigINumStr := "30987680500513189125.14259702468435"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
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
      "err = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBiNumStr,
      multiplicandBiNumStr,
      err.Error())

    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix, expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix, expectedBigINumSign, result.sign)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("%v\n"+
      "Error: Expected actualNumStr='%v'\n"+
      "Instead, actualNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, actualNumStr)

  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_03"

  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedBigINumStr := "-30987680500513189125.14259702468435"

  expectedBigINumSign := -1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
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
      "err = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNums(\n"+
      "  multiplierBiNum, multiplicandBiNum)\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBiNumStr,
      multiplicandBiNumStr,
      err.Error())

    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\nInstead, BigIntNum= '%s'.\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\nComparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix, expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\nError: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix, expectedBigINumSign, result.sign)
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_04"

  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedBigINumStr := "22197234145.3632"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
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
      "err = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNums(\n"+
      "  multiplierBiNum, multiplicandBiNum)\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBiNumStr,
      multiplicandBiNumStr,
      err.Error())

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'. Instead, number sign='%v'",
      ePrefix, expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_05"

  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0
  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
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
      "err = iaMultiplier.Multiply()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = iaResult.OptimizeIntArrayLen(true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.OptimizeIntArrayLen(true)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNums(\n"+
      "  multiplierBiNum, multiplicandBiNum)\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBiNumStr,
      multiplicandBiNumStr,
      err.Error())

    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultNumBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultNumBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultNumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix, expectedBigINum.bigInt.Text(10),
      result.bigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix,
      expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }
  return
}

func TestBigIntMathMultiply_MultiplyBigIntNums_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNums_06"

  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875,94572"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps='%v'\nError='%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandBiNumStr, err := multiplicandBiNum.GetNumStr()\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, multiplicandStr, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNums(multiplierBiNum, multiplicandBiNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNums(\n"+
      "multiplierBiNum, multiplicandBiNum)\n"+
      "multiplierBiNum= '%v'\n"+
      "multiplicandBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierBiNumStr,
      multiplicandBiNumStr,
      err.Error())

    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("%v\n"+
      "Error: Expected NumStr='%v'.\nInstead, NumStr='%v'.\n\n",
      ePrefix, expectedNumStr, actualNumStr)
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
      "Error: Expected numSeps='%v'.\nInstead, numSeps='%v'.\n\n",
      ePrefix, expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwo_01"

  num1Int := 242
  precision := uint(2)
  expectedNum1 := "2.42"
  expectedResult := "4.84"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)\n"+
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\nInstead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwo_02"

  num1Int := 921123
  precision := uint(3)
  expectedNum1 := "921.123"
  expectedResult := "1842.246"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumByTwo(num1)\n"+
      "num1='%v'\nError='%v'\n\n",
      ePrefix, num1NumStr, err.Error())
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwo_03"

  num1Int := -4751
  precision := uint(1)
  expectedNum1 := "-475.1"
  expectedResult := "-950.2"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)\n"+
      "num1='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, num1NumStr, err.Error())
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwo_04"

  num1Int := 475
  precision := uint(0)
  expectedNum1 := "475"
  expectedResult := "950"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)\n"+
      "num1='%v'\nError='%v'\n\n",
      ePrefix, num1NumStr, err.Error())
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\nError: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwo_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwo_05"

  num1Int := 0
  precision := uint(0)
  expectedNum1 := "0"
  expectedResult := "0"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(num1)\n"+
      "num1='%v'\nError='%v'\n\n",
      ePrefix, num1NumStr, err.Error())
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\nInstead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_01"

  num1Int := 242
  precision := uint(2)
  exponent := uint(1)
  expectedNum1 := "2.42"
  expectedResult := "4.84"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_02"

  num1Int := 242
  precision := uint(2)
  exponent := uint(2)
  expectedNum1 := "2.42"
  expectedResult := "9.68"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_03"

  num1Int := 97123
  precision := uint(3)
  exponent := uint(4)
  expectedNum1 := "97.123"
  expectedResult := "1553.968"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_04"
  num1Int := -97123
  precision := uint(3)
  exponent := uint(4)
  expectedNum1 := "-97.123"
  expectedResult := "-1553.968"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_05"

  num1Int := 7
  precision := uint(0)
  exponent := uint(10)
  expectedNum1 := "7"
  expectedResult := "7168"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumByTwoToPower_06"

  num1Int := 0
  precision := uint(0)
  exponent := uint(10)
  expectedNum1 := "0"
  expectedResult := "0"

  num1, err := new(BigIntNum).NewInt(num1Int, precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1, err := new(BigIntNum).NewInt(num1Int, precision)\n"+
      "num1Int= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Int,
      precision,
      err.Error())

    return
  }

  num1NumStr, err := num1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "num1NumStr, err := num1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyBigIntNumByTwoToPower(num1, exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "MultiplyBigIntNumByTwoToPower(num1, exponent)\n"+
      "num1= '%v'\n"+
      "exponent= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1NumStr,
      exponent,
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

  if expectedResult != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedResult, resultNumStr)
    return
  }

  if expectedNum1 != num1NumStr {
    t.Errorf("%v\n"+
      "Error: Expected result='%v'.\n"+
      "Instead, result='%v'\n\n",
      ePrefix, expectedNum1, num1NumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumArray_01"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var bINumArrayNumStr, iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    bINumArrayNumStr, err = bINumArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArrayNumStr, err = bINumArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "bINumArray[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, bINumArrayNumStr, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v';\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, iaNumStr, i, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).
    NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, expectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr,err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumArray_02"
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
  originalExpectedBigINumStr := "11995826664.26376575446779648"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var bINumArrayNumStr, iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    bINumArrayNumStr, err = bINumArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArrayNumStr, err = bINumArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "bINumArray[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, bINumArrayNumStr, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v';\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, iaNumStr, i, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExpectedBigINumSign != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      ePrefix, originalExpectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("%v\n"+
      "Error: Expected actualNumStr='%v'\n"+
      "Instead, actualNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumArray_03"

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
  originalExpectedBigINumStr := "2212352.1767579232"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var bINumArrayNumStr, iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    bINumArrayNumStr, err = bINumArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArrayNumStr, err = bINumArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "bINumArray[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, bINumArrayNumStr, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v';\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, iaNumStr, i, err.Error())
      return
    }

  } // End of for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr2, err = expectedBigINum.GetNumStr()\n"+
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

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: BigIntNum Values are unequal!\n"+
      "Expected BigIntNum         = '%s'.\n"+
      "Instead, Result BigIntNum  = '%s'.\n"+
      "Original ExpectedBigINumStr= '%s'",
      ePrefix,
      expectedBigINumBigInt.Text(10), result.bigInt.Text(10),
      originalExpectedBigINumStr)
    return
  }

  if originalExpectedBigINumStr != resultNumStr ||
    originalExpectedBigINumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: BigInt Number strings are unequal!\n"+
      "Expected originalExpectedBigINumStr='%s'.\n"+
      "Instead, ResultNumStr              ='%s'.\n"+
      "expectedBigINumStr                 ='%s'.\n",
      ePrefix,
      originalExpectedBigINumStr,
      resultNumStr,
      expectedBigINumStr)
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error:  BigIntNum *big.Int values are unequal!\n"+
      "Expected BigIntNumBigInt='%s'.\n"+
      "Instead, ResultNumBigInt='%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "Instead,    resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumArray_04"

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
  originalExpectedBigINumStr := "-20408.5138429311978576052224"

  originalExpectedBigINumSign := -1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var bINumArrayNumStr, iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    bINumArrayNumStr, err = bINumArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArrayNumStr, err = bINumArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "bINumArray[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, bINumArrayNumStr, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v';\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, iaNumStr, i, err.Error())
      return
    }

  } // End of for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr2, err = expectedBigINum.GetNumStr()\n"+
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

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResult, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("%v\n"+
      "Error: BigIntNum Values are unequal!\n"+
      "Expected BigIntNum         = '%s'.\n"+
      "Instead, Result BigIntNum  = '%s'.\n"+
      "Original ExpectedBigINumStr= '%s'",
      ePrefix,
      expectedBigINumBigInt.Text(10), result.bigInt.Text(10),
      originalExpectedBigINumStr)
    return
  }

  if expectedBigINumStr != resultNumStr ||
    expectedBigINumStr != originalExpectedBigINumStr ||
    resultNumStr != originalExpectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error:  BigIntNum *big.Int values are unequal!\n"+
      "Expected BigIntNumBigInt='%s'.\n"+
      "Instead, ResultNumBigInt='%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
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

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumArray_05"

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

  // product = 11995826664,26376575446779648
  originalExpectedBigINumStr := "11995826664,26376575446779648"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  frenchNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  frenchNumSeps.DecimalSeparator = frenchDecSeparator
  frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
  frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierBiNum.SetNumericSeparatorsDto(frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierBiNum.SetNumericSeparatorsDto(frenchNumSeps)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var bINumArrayNumStr, iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    bINumArrayNumStr, err = bINumArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArrayNumStr, err = bINumArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "bINumArray[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, bINumArrayNumStr, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v';\n"+
        "Cycle Index='%d'\n"+
        "Error='%v'\n\n",
        ePrefix, iaNumStr, i, err.Error())
      return
    }

  } // End of for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(originalExpectedBigINumStr, &frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(originalExpectedBigINumStr, &frenchNumSeps)\n"+
      "originalExpectedBigINumStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, frenchNumSeps.String(), err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr2, err = expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExpectedBigINumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExpectedBigINumStr != expectedBigINumStr\n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, originalExpectedBigINumStr, expectedBigINumStr)

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

  if expectedBigINumStr != resultNumStr ||
    expectedBigINumStr != originalExpectedBigINumStr ||
    resultNumStr != originalExpectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
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

  if !frenchNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected numSeps='%v'.\n"+
      "Instead, numSeps='%v'.\n\n",
      ePrefix,
      frenchNumSeps.String(),
      actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_01"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

  } // End of 2nd ('j') For Loop

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_02"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

  } // End of 2nd ('j') For Loop

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_03"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

  } // End of 2nd ('j') For Loop

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_04"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

  } // End of 2nd ('j') For Loop

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_05"

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

  } // End of 2nd ('j') For Loop

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumOutputToArray_06"

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
    "-3123,12",
    "811,2",
    "-122,4288",
    "-249,6",
    "-165270,2376",
    "152,7552",
  }

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps='%v'\nError='%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
  }

  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of first for loop

  multiplierBiNumStr, err := multiplierBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNumStr, err := multiplierBiNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyBigIntNumOutputToArray(multiplierBiNum, bINumArray)\n"+
      "multiplierBiNum='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierBiNumStr, err.Error())
    return
  }

  var resultNumStr string
  var actualNumSeps NumericSeparatorDto

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected NumStr[%v]='%v'. Instead result[%v] NumStr='%v'. ",
        ePrefix, j, expectedNumStrs[j], j, resultNumStr)
      return
    }

    actualNumSeps, err = result[j].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualNumSeps, err = result[%d].GetNumericSeparatorsDto()\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("%v\n"+
        "Error: Expected NumSeps='%v'.\n"+
        "Instead, NumSeps='%v'.\n"+
        "Index='%v'\n\n",
        ePrefix, expectedNumSeps.String(), actualNumSeps.String(), j)
    }

  } // End of 2nd ('j') For Loop

  return
}
